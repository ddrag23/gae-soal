package model

import (
	"gorm.io/gorm"
)

type Role struct{
	gorm.Model
	Name uint `gorm:"not null" json:"name"`
}