package users

import (
	"github.com/jinzhu/gorm"
	"github.com/satyasure/webserver/server/model/user"
)

// ListUserRequest request struct
type ListUsersRequest struct {
	Limit uint
}

// ListUserResponse response struct
type ListUsersResponse struct {
	Users *[]user.User `json:"users"`
}

// ListUsers returns a list of users from database
func ListUsers(db *gorm.DB, req *ListUsersRequest) (*ListUsersResponse, error) {
	users, err := user.List(db, req.Limit)
	res := &ListUsersResponse{Users: users}
	return res, err
}
