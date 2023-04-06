package users

import (
	"github.com/jinzhu/gorm"
	"github.com/satyasure/webserver/server/model/user"
)

//DeleteUserRequest request struct
type DeleteUserRequest struct {
	ID string
}

//DeleteUserResponse response struct
type DeleteUserResponse struct {
}

//DeleteUser deletes a user from database
func DeleteUser(db *gorm.DB, req *DeleteUserRequest) (*DeleteUserResponse, error) {
	err := user.Delete(db, req.ID)
	res := &DeleteUserResponse{}
	return res, err
}
