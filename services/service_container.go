package services

import (
	"gorm.io/gorm"
)

type ServiceContainer struct {
	ProfileService *ProfileService
	AuthService    *AuthService
	UserService    *UserService
	// Add other services here as needed
}

func InitServiceContainer(db *gorm.DB) *ServiceContainer {
	return &ServiceContainer{
		ProfileService: &ProfileService{DB: db},
		AuthService:    &AuthService{DB: db},
		UserService:    &UserService{DB: db},
		// Initialize other services here
	}
}
