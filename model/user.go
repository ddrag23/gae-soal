package model

import (
	"gorm.io/gorm"
)
type User struct{
	gorm.Model
	Username string `gorm:"unique_index;not null" json:"username"`
	Email string `gorm:"unique_index;not null" json:"email"`
	Name string `gorm:"unique_index;not null" json:"name"`
	Password string `gorm:"not null" json:"password"`
	RoleId string `gorm:"not null" json:"role_id"`
}