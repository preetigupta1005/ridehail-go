package handlers

import (
	"net/http"

	"github.com/preetigupta1005/ridehail-go/middlewares"
	"github.com/preetigupta1005/ridehail-go/models"
	"github.com/preetigupta1005/ridehail-go/repository"
	"github.com/preetigupta1005/ridehail-go/services"
	"github.com/preetigupta1005/ridehail-go/utils"
)

func RequestRideHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.UserIDKey).(string)

	var req models.RequestRideBody
	if err := utils.ParseBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	if err := utils.ValidateStruct(req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "validation failed")
		return
	}
	
	ride, err := services.RequestRide(
		userID,
		req.PickupLat, req.PickupLng, req.PickupAddress,
		req.DropLat, req.DropLng, req.DropAddress,
	)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to create ride")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, ride)
}

func AcceptRideHandler(w http.ResponseWriter, r *http.Request) {
	driverID := r.Context().Value(middlewares.UserIDKey).(string)
	rideID := r.PathValue("id")

	if err := repository.AcceptRideRequest(rideID, driverID); err != nil {
		utils.RespondError(w, http.StatusConflict, err, "failed to accept ride")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "ride accepted"})
}

func StartRideHandler(w http.ResponseWriter, r *http.Request) {
	driverID := r.Context().Value(middlewares.UserIDKey).(string)
	rideID := r.PathValue("id")

	if err := repository.StartRide(rideID, driverID); err != nil {
		utils.RespondError(w, http.StatusConflict, err, "failed to start ride")
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "ride started"})
}

type EndRideBody struct {
	FareAmount float64 `json:"fare_amount"`
}

func EndRideHandler(w http.ResponseWriter, r *http.Request) {
	driverID := r.Context().Value(middlewares.UserIDKey).(string)
	rideID := r.PathValue("id")

	fare, distance, err := repository.EndRide(rideID, driverID)
	if err != nil {
		utils.RespondError(w, http.StatusConflict, err, "failed to end ride")
		return
	}

	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{
		"message":     "ride completed",
		"fare_amount": fare,
		"distance_km": distance,
	})
}

func CancelRideHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.UserIDKey).(string)
	role := r.Context().Value(middlewares.RoleKey).(string)
	rideID := r.PathValue("id")

	var req models.CancelRideBody
	utils.ParseBody(r, &req) //optional-no need to handle error

	if err := repository.CancelRide(rideID, userID, role, req.Reason); err != nil {
		utils.RespondError(w, http.StatusForbidden, err, "failed to cancel ride")
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "ride cancelled"})
}

func GetMyRidesHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.UserIDKey).(string)
	role := r.Context().Value(middlewares.RoleKey).(string)

	var rides []models.Ride
	var err error

	if role == "passenger" {
		rides, err = repository.GetRidesByPassenger(userID)
	} else {
		rides, err = repository.GetRidesByDriver(userID)
	}

	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to fetch rides")
		return
	}
	utils.RespondJSON(w, http.StatusOK, rides)
}
