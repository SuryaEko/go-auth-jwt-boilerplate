// models/user.go
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;size:32" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"default:'user'" json:"role"` // Default role is 'user'
}
