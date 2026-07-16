package models

import "time"

type User struct {
	ID           string    `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Email        string    `db:"email" json:"email"`
	Phone        string    `db:"phone" json:"phone"`
	PasswordHash string    `db:"password_hash" json:"-"`
	Role         string    `db:"role" json:"role"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type SignupRequest struct {
	Name          string `json:"name" validate:"required,min=2,max=100"`
	Email         string `json:"email" validate:"required,email"`
	Phone         string `json:"phone" validate:"required,len=10,numeric"`
	Password      string `json:"password" validate:"required,min=6"`
	Role          string `json:"role" validate:"required,oneof=passenger driver"`
	VehicleNumber string `json:"vehicle_number,omitempty"`
	VehicleType   string `json:"vehicle_type,omitempty"`
	LicenseNumber string `json:"license_number,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
