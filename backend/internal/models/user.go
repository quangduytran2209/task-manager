package models

import (
	"gorm.io/gorm"
)

type User struct{
	gorm.Model 
	Username 	string `gorm:"unique;not null" json:"username"`
	Password 	string `gorm:"-"`
	Email 		string `gorm:"unique;not null" json:"email"`
}