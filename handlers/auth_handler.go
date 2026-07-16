package handlers

import (
	"net/http"

	"github.com/preetigupta1005/ridehail-go/models"
	"github.com/preetigupta1005/ridehail-go/services"
	"github.com/preetigupta1005/ridehail-go/utils"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var req models.SignupRequest
	if err := utils.ParseBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}
	if err := utils.ValidateStruct(req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "validation failed")
		return
	}

	if req.Role == "driver" && (req.VehicleNumber == "" || req.VehicleType == "" || req.LicenseNumber == "") {
		utils.RespondError(w, http.StatusBadRequest, nil, "vehicle details required for driver signup")
		return
	}

	user, err := services.Signup(req.Name, req.Email, req.Phone, req.Password, req.Role,
		req.VehicleNumber, req.VehicleType, req.LicenseNumber)
	if err != nil {
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

	if err := utils.ValidateStruct(req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "validation failed")
		return
	}

	token, err := services.Login(req.Email, req.Password)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "invalid email or password")
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]string{"token": token})
}
