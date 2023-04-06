import {
    USERS,
    USERS_LOADING,
    USER_ACTIVE,
    API_URL,
} from '../constants/actions.js';
import { push } from 'connected-react-router';
import { isNullOrUndefined } from '../utility/utility.js';

//pipelines
export function refreshUsers({userId}) {
    return (dispatch) => {
        dispatch(usersLoading(true));
        dispatch(listUsers(userId));
    }
}

export function selectUser({ user }) {
    return (dispatch) => {
        if(isNullOrUndefined(user)){
            dispatch(activeUser({}));
            dispatch(push({ pathname: '/',search:``}))
            return
        }
        dispatch(activeUser(user));
        dispatch(push({ pathname: '/', search: `?user=${user.id}` }))
        dispatch(refreshUsers({userId:null}));
    }
}

export function createUser({ user,userId }) {
    return (dispatch) => {
        const data = user
        const endpoint = `${API_URL}/users`;
        dispatch(usersLoading(true));
        fetch(endpoint,{
                method: 'POST',
                mode: "cors",
                body: JSON.stringify(data),
                headers: {
                    "Content-Type": "application/json",
                },
        })
        .then(()=>{dispatch(refreshUsers({userId:null}))})
        .catch(
            e => {console.log(e);dispatch(usersLoading(false))}

        )
    }
}
export function deleteUser( {userId} ) {
    return (dispatch) => {
        const endpoint = `${API_URL}/users/${userId}`;
        dispatch(usersLoading(true));
        fetch(endpoint,{
                method: 'DELETE',
                mode: "cors",
        })
        .then(()=>{dispatch(refreshUsers({userId:null}))})
        .catch(
            e => {console.log(e);dispatch(usersLoading(false))}
        )
    }
}

function activeUser(user) {
    return {
        type: USER_ACTIVE,
        payload: user
    }
}

function listUsers(userId=null) {
    return (dispatch) => {
        const endpoint = `${API_URL}/users`;
        dispatch(usersLoading(true));
        fetch(endpoint, { method: 'GET', mode: "cors", })
            .then(res => {
                res.json().then((json) => {
                    dispatch(users(json.users));
                    if(!isNullOrUndefined(userId)){
                        let user = json.users.find((user)=>{return user.id===userId})
                        dispatch(selectUser({user}))
                    }
                }).catch(e => { console.log(e) })
            })
            .catch(e => { console.log(e)})
            .finally(()=>{dispatch(usersLoading(false))})
    }
}


export function usersLoading(bool) {
    return {
        type: USERS_LOADING,
        payload: bool
    }
}


export function users(usersList) {
    return {
        type: USERS,
        payload: usersList
    }
}