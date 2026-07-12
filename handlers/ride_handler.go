package handlers

import (
	"net/http"

	"github.com/preetigupta1005/ridehail-go/middlewares"
	"github.com/preetigupta1005/ridehail-go/models"
	"github.com/preetigupta1005/ridehail-go/repository"
	"github.com/preetigupta1005/ridehail-go/utils"
)

func RequestRideHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.UserIDKey).(string)

	var req models.RequestRideBody
	if err := utils.ParseBody(r, &req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}

	ride := &models.Ride{
		PassengerID: userID,
		PickupLat:   req.PickupLat,
		PickupLng:   req.PickupLng,
		DropLat:     req.DropLat,
		DropLng:     req.DropLng,
	}

	if err := repository.CreateRide(ride); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to create ride")
		return
	}

	driverIDs, err := repository.GetNearbyDrivers(req.PickupLat, req.PickupLng, 5000)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to find nearby drivers")
		return
	}

	if len(driverIDs) == 0 {
		utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "ride created, no drivers nearby yet"})
		return
	}

	if err := repository.CreateRideRequests(ride.ID, driverIDs); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to notify drivers")
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

	if err := repository.EndRide(rideID, driverID); err != nil {
		utils.RespondError(w, http.StatusConflict, err, "failed to end ride")
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "ride completed"})
}

func CancelRideHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middlewares.UserIDKey).(string)
	role := r.Context().Value(middlewares.RoleKey).(string)
	rideID := r.PathValue("id")

	if err := repository.CancelRide(rideID, userID, role); err != nil {
		utils.RespondError(w, http.StatusForbidden, err, "failed to cancel ride")
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]string{"message": "ride cancelled"})
}
