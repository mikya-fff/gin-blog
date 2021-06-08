package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Pssword string `gorm:"type:varchar(20);not null " json:"pssword"`
	Role int `gorm:"type:int" json:"role"`
}


