// models/user.go
package models

import (
	"github.com/SuryaEko/go-auth-jwt-boilerplate/pkg"
	"gorm.io/gorm"
)

type UserGorm struct {
	DB *gorm.DB
}

type User struct {
	gorm.Model
	Username string `gorm:"unique;size:32" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Role     string `gorm:"default:'user'" json:"role"` // Default role is 'user'
}

func (ug *UserGorm) List(pagination pkg.Pagination) (*pkg.Pagination, error) {
	var users []*User

	ug.DB.Scopes(paginate(users, &pagination, ug.DB)).Find(&users)
	pagination.Rows = users

	return &pagination, nil
}
