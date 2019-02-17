package models

import (
	"github.com/jinzhu/gorm"
)

// Role is a domain model for role
type Role struct {
	gorm.Model
	Code string `json:"code" gorm:"not null;unique"`
	Name string `json:"name" gorm:"not null"`
}
