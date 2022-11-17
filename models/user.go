package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Mobile    string
	Email     string
}

func (u *User) SaveUser() (*User, error) {
	err := DB.Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}
