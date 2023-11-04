package model

import "gorm.io/gorm"

type Book struct {
	Id          int    `json:"id" gorm:"type:int;primary_key"`
	Book        string `json:"book" gorm:"type:varchar(50);not null"`
	Description string `json:"description" gorm:"type:varchar(255);not null"`
	Author      string `json:"author" gorm:"type:varchar(50)"`
	*gorm.Model
}
