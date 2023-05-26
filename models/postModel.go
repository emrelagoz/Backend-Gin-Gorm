package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Phone     string
}
