package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	//添加外键依赖
	Category Category `gorm:"foreignkey:Name"`
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	Cid int `gorm:"type:int;not null" json:"cid"`
	Desc string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img string `gorm:"type:varchar(100)" json:"img"`
}
