package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/preetigupta1005/ridehail-go/database"
	"github.com/preetigupta1005/ridehail-go/models"
)

func CreateUser(tx *sqlx.Tx, user *models.User) error {
	query := `INSERT INTO users (name, email, phone, password_hash, role) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
	return tx.QueryRowx(query, user.Name, user.Email, user.Phone, user.PasswordHash, user.Role).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, name, email, phone, password_hash, role, created_at, updated_at 
	          FROM users WHERE email=$1`
	err := database.DB.Get(&user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
