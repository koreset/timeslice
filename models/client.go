package models

import "github.com/jinzhu/gorm"

type Client struct {
	gorm.Model
	Name string `gorm:"not null"`
}
