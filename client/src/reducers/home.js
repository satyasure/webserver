import {USER_ACTIVE,USERS,USERS_LOADING} from '../constants/actions.js'

//users
export function users(state=[],action){
    switch(action.type){
        case USERS:
            return action.payload;
        default:
            return state;
    }
}
export function usersLoading(state=false,action){
    switch(action.type){
        case USERS_LOADING:
        return action.payload;
        default:
        return state;
    }
}
export function userActive(state={},action){
    switch(action.type){
        case USER_ACTIVE:
        return action.payload;
        default:
        return state;
    }
}