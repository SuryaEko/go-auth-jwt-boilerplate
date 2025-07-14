package dto

type UpdateProfileInput struct {
	Username string `json:"username" binding:"required"`
}

type UpdatePasswordProfileInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
	RePassword  string `json:"re_password" binding:"required,eqfield=NewPassword"`
}
