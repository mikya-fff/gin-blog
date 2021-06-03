package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type: " json:""`
	Pssword string `gorm:"type: " json:""`
	Role int `gorm:"type: " json:""`
}


