package models

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	ClientID uint
	Name string
}