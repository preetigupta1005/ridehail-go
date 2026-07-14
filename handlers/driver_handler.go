package handlers

import (
	"net/http"

	"github.com/preetigupta1005/ridehail-go/middlewares"
	"github.com/preetigupta1005/ridehail-go/repository"
	"github.com/preetigupta1005/ridehail-go/utils"
)

type AvailabilityRequest struct {
	IsAvailable bool `json:"is_available"`
}

func UpdateAvailabilityHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.UserIDKey).(string)

	var req AvailabilityRequest
	if err := utils.ParseBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if err := repository.UpdateAvailability(userID, req.IsAvailable); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to update availability")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "availability updated"})
}

type LocationRequest struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func UpdateLocationHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.UserIDKey).(string)

	var req LocationRequest
	if err := utils.ParseBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if err := repository.UpdateLocation(userID, req.Lat, req.Lng); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to update location")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "location updated"})
}
