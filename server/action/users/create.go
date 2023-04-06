package users

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/satyasure/webserver/server/src/server/model/user"
)

//CreateUserRequest request struct
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Sex      string `json:"sex"`
	Species  string `json:"species"`
	Color    string `json:"color"`
	Breed    string `json:"breed"`
	ImageURL string `json:"imageURL"`
}

//CreateUserResponse response struct
type CreateUserResponse struct {
	ID string `json:"id"`
}

//CreateUser creates a user in database
func CreateUser(db *gorm.DB, req *CreateUserRequest) (*CreateUserResponse, error) {
	uuid, _ := uuid.NewRandom()
	newUser := &user.user{
		ID:       uuid.String(),
		Name:     req.Name,
		Sex:      req.Sex,
		Species:  req.Species,
		Color:    req.Color,
		Breed:    req.Breed,
		ImageURL: req.ImageURL,
	}
	id, err := user.Create(db, newUser)
	res := &CreateUserResponse{ID: id}
	return res, err
}
