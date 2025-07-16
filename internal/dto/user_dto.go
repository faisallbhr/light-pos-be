package dto

type UserResponse struct {
	ID    uint           `json:"id"`
	Name  string         `json:"name"`
	Email string         `json:"email"`
	Roles []RoleResponse `json:"roles"`
}

type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserUpdateRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type ChangePasswordRequest struct {
	CurrentPassword      string `json:"current_password" binding:"required"`
	NewPassword          string `json:"new_password" binding:"required,min=6,max=255"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required,eqfield=NewPassword"`
}
