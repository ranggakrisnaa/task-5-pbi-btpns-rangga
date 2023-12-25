package models

import (
	"html"
	"strings"

	"github.com/ranggakrisnaa/task-5-pbi-btpns-rangga/helpers"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	passwordHash, err := helpers.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(passwordHash)

	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return nil
}

func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Password") {
		passwordHash, err := helpers.HashPassword(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordHash)
	}

	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return nil
}

func (user *User) ComparePassword(password string) error {
	return helpers.CheckPasswordHash(password, user.Password)
}
