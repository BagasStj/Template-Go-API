package models

// LoginRequest represents the login request payload
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterRequest represents the registration request payload
type RegisterRequest struct {
	Name            string `json:"name" binding:"required,min=2"`
	Email           string `json:"email" binding:"required,email"`
	Username        string `json:"username" binding:"required,min=3"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	PhoneNumber     string `json:"phone_number,omitempty"`
}

// UpdateUserRequest represents the update user request payload
type UpdateUserRequest struct {
	Name        string `json:"name,omitempty"`
	Username    string `json:"username,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

// PaginationRequest represents pagination parameters
type PaginationRequest struct {
	Page  int `form:"page,default=1" binding:"min=1"`
	Limit int `form:"limit,default=10" binding:"min=1,max=100"`
}
