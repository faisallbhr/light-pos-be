package dto

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=255"`
}

type RegisterRequest struct {
	Name                 string `json:"name" binding:"required"`
	Email                string `json:"email" binding:"required,email"`
	Password             string `json:"password" binding:"required,min=6,max=255"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required,eqfield=Password"`
}

type RefreshRequest struct {
	Refresh string `json:"refresh" binding:"required"`
}

type TokenResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type LoginResponse struct {
	User  UserResponse  `json:"user"`
	Token TokenResponse `json:"token"`
}
