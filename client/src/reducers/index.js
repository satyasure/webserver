import {combineReducers} from 'redux';
import {userActive,users,usersLoading} from './home.js';
const rootReducer=combineReducers({
   users:combineReducers({users,usersLoading,userActive}),
})
export default rootReducer;