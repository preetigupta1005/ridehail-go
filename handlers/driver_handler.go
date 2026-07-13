package handlers

import (
	"fmt"
	"net/http"

	"github.com/preetigupta1005/ridehail-go/middlewares"
	"github.com/preetigupta1005/ridehail-go/repository"
	"github.com/preetigupta1005/ridehail-go/utils"
	"github.com/preetigupta1005/ridehail-go/websocket"
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

func UpdateLocationHandler(hub *websocket.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		passengerID, err := repository.GetActivePassengerID(userID)
		if err == nil { // agar koi active ride mili
			message := fmt.Sprintf(`{"driver_id":"%s","lat":%f,"lng":%f}`, userID, req.Lat, req.Lng)
			hub.SendToUser(passengerID, []byte(message))
		}

		utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "location updated"})
	}
}
