package dto

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required,min=6"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}
