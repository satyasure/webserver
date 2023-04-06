package user

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//Create creates a user in the database
func Create(db *gorm.DB, user *user) (string, error) {
	err := db.Create(user).Error
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

//FindById returns a user with a given id, or nil if not found
func FindById(db *gorm.DB, id string) (*user, error) {
	var user user
	err := db.Find(&user, &user{ID: id}).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//FindByName returns a user with a given name, or nil if not found
func FindByName(db *gorm.DB, name string) (*user, error) {
	var user user
	err := db.Find(&user, &user{Name: name}).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//List returns all Users in database, with a given limit
func List(db *gorm.DB, limit uint) (*[]user, error) {
	var users []user
	err := db.Find(&users).Limit(limit).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

//Delete deletes a user in the database
func Delete(db *gorm.DB, id string) error {
	user, err := FindById(db, id)
	if err != nil {
		fmt.Printf("1:%v", err)
		return err
	}
	err = db.Delete(user).Error
	fmt.Printf("2:%v", err)
	return err
}