package users

import (
	"github.com/jinzhu/gorm"
	"github.com/satyasure/webserver/tree/master/server/model/user"
)

//GetUserRequest request struct
type GetUserRequest struct {
	ID string
}

//GetUserResponse response struct
type GetUserResponse struct {
	user *user.user `json:"user"`
}

//GetUser returns a user from database
func GetUser(db *gorm.DB, req *GetUserRequest) (*GetUserResponse, error) {
	user, err := user.FindById(db, req.ID)
	res := &GetUserResponse{user: user}
	return res, err
}
