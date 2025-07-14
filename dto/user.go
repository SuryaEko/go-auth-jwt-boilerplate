package dto

type CreateUserInput struct {
	Username   string `json:"username" binding:"required"`
	Role       string `json:"role" binding:"required,oneof=admin user"`
	Password   string `json:"password" binding:"required,min=6"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type UpdateUserInput struct {
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=admin user"`
}

type UpdatePasswordUserInput struct {
	NewPassword string `json:"new_password" binding:"required,min=6"`
	RePassword  string `json:"re_password" binding:"required,eqfield=NewPassword"`
}
