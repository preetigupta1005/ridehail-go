package models

type SignupRequest struct {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Phone    string `json:"phone" validate:"required"`
    Password string `json:"password" validate:"required,min=6"`
    Role     string `json:"role" validate:"required,oneof=passenger driver"`
}