package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title    string
	Caption  string
	PhotoUrl string
	UserId   uint
	User     []User `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
