package models

import "github.com/jinzhu/gorm"

type Employee struct {
	gorm.Model
	FirstName    string `gorm:"not null;type:varchar(100)" json:"first_name"`
	LastName     string `gorm:"not null;type:varchar(100)" json:"last_name"`
	Email        string `gorm:"unique;not null;index:email" json:"email"`
	Password     string `gorm:"-" json:"password"`
	PasswordHash string `gorm:"not null"`
}
