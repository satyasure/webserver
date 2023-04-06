package user

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Create creates a user in the database
func Create(db *gorm.DB, user *User) (string, error) {
	err := db.Create(user).Error
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

// FindById returns a user with a given id, or nil if not found
func FindById(db *gorm.DB, id string) (*User, error) {
	var user User
	err := db.Find(&user, &User{ID: id}).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByName returns a user with a given name, or nil if not found
func FindByName(db *gorm.DB, name string) (*User, error) {
	var user User
	err := db.Find(&user, &User{Name: name}).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// List returns all Users in database, with a given limit
func List(db *gorm.DB, limit uint) (*[]User, error) {
	var users []User
	err := db.Find(&users).Limit(limit).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

// Delete deletes a user in the database
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
