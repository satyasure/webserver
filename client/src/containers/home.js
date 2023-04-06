import {bindActionCreators} from 'redux';
import {connect} from 'react-redux';
import Home from '../components/views/home.jsx'
import {
    createUser,
    refreshUsers,
    selectUser,
    deleteUser,
} from '../actions/home.js';

const mapStateToProps=(state)=>{
   
return{
    users: state.users.users,
    usersLoading: state.users.usersLoading,
    userActive: state.users.userActive,
    router: state.router,
    };
};

const mapDispatchToProps = (dispatch) =>{
    return{
      refreshUsers: bindActionCreators(refreshUsers,dispatch),
      selectUser: bindActionCreators(selectUser,dispatch),
      createUser: bindActionCreators(createUser,dispatch),
      deleteUser: bindActionCreators(deleteUser,dispatch),
    }
};

export default connect(mapStateToProps,mapDispatchToProps)(Home);