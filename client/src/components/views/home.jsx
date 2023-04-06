import React, { Component } from 'react';
import PropTypes from 'prop-types';
import classNames from 'classnames';
import { withStyles } from '@material-ui/core/styles';
import {
    IconButton,
    Typography,
    Toolbar,
    CircularProgress
} from '@material-ui/core';
import {
    Refresh as RefreshIcon,
} from '@material-ui/icons';
import withRoot from '../../themes/withRoot';
import { styles } from './home.js';
import user from '../cards/user.jsx';
import { isNullOrUndefined } from '../../utility/utility.js'
import Create from '../dialog/create.jsx';

class Home extends Component {
    state = {};
    componentWillMount = () => {
        const kvs = this.props.router.location.search.split("&")
        let userId
        for (let kv of kvs) {
          const kv_pair = kv.split("=")
          const key = kv_pair[0]
          const value = kv_pair[1]
          if (key.includes("user")){
             userId = value;
          }
          if(isNullOrUndefined(userId)){
            userId=null;

          }
            this.props.refreshUsers({userId});
        }
    }

    renderUsers = (users,userActive) => {
        if (isNullOrUndefined(users)) {
            return null
        }
        return users.map((user, index) => {
            return (
                <user
                    key={index}
                    index={index}
                    user={user}
                    selectUser={this.props.selectUser}
                    userActive={userActive}
                    deleteUser={this.props.deleteUser}
                />
            )
        })
    }

    renderContent = (usersList,usersLoading) => {
        if(usersLoading){
            return   <CircularProgress/>
        }
        return usersList.length>0?usersList:<Typography variant="h5">Click the "+" button to add users</Typography>
        
    }

    render() {
        const { users, userActive, usersLoading,classes } = this.props;
        const usersList = this.renderUsers(users,userActive)
        const content = this.renderContent(usersList,usersLoading)
        return (
            <div className={classNames(classes.root)}>
                <Toolbar>
                     <Typography variant="h3" internalDeprecatedVariant>My Users!</Typography>
                    <div className={classes.flex}/>
                    <Create createUser={this.props.createUser} userActive/>
                    <IconButton onClick={() => { this.props.refreshUsers({userId:userActive.id}) }} >
                        <RefreshIcon />
                    </IconButton>
                </Toolbar>
               {content}
            </div>
        )
    }
}

Home.propTypes = {
    classes: PropTypes.object.isRequired,
    theme: PropTypes.object.isRequired,
    refreshUsers: PropTypes.func.isRequired,
    selectUser: PropTypes.func.isRequired,
    users: PropTypes.array.isRequired,
    userActive: PropTypes.object.isRequired,
    usersLoading:PropTypes.bool.isRequired,
    deleteUser: PropTypes.func.isRequired,
    createUser: PropTypes.func.isRequired,
};


export default withRoot(withStyles(styles, { withTheme: true })(Home));