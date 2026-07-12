package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/preetigupta1005/ridehail-go/database"
	"github.com/preetigupta1005/ridehail-go/models"
	"github.com/preetigupta1005/ridehail-go/repository"
	"github.com/preetigupta1005/ridehail-go/utils"
)

func Signup(name, email, phone, password, role, vehicleNumber, vehicleType, licenseNumber string) (*models.User, error) {
	hashedPwd, err := utils.HashedPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:         name,
		Email:        email,
		Phone:        phone,
		PasswordHash: hashedPwd,
		Role:         role,
	}

	err = database.Tx(func(tx *sqlx.Tx) error {
		if err := repository.CreateUser(tx, user); err != nil {
			return err
		}
		if role == "driver" {
			return repository.CreateDriverDetails(tx, user.ID, vehicleNumber, vehicleType, licenseNumber)
		}
		return nil
	})

	return user, err
}

func Login(email, password string) (string, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	if err := utils.CheckPassword(password, user.PasswordHash); err != nil {
		return "", err
	}
	return utils.GenerateToken(user.ID, user.Role)
}
