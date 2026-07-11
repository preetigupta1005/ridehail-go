package handlers

import (
	"net/http"

	"github.com/preetigupta1005/ridehail-go/models"
	"github.com/preetigupta1005/ridehail-go/repository"
	"github.com/preetigupta1005/ridehail-go/utils"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req models.SignupRequest
	if err := utils.ParseBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if req.Role != "passenger" && req.Role != "driver" {
		utils.RespondError(w, http.StatusBadRequest, nil, "role must be passenger or driver")
		return
	}

	hashedPwd, err := utils.HashedPassword(req.Password)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to process password")
		return
	}

	user := &models.User{
		Name:         req.Name,
		Email:        req.Email,
		Phone:        req.Phone,
		PasswordHash: hashedPwd,
		Role:         req.Role,
	}

	if err := repository.CreateUser(user); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to create user")
		return
	}
	utils.RespondJSON(w, http.StatusCreated, user)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := utils.ParseBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "invalid email or password")
		return
	}

	if err := utils.CheckPassword(req.Password, user.PasswordHash); err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "invalid email or password")
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to generate token")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"token": token})
}
