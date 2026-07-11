package repository

import (
	"github.com/preetigupta1005/ridehail-go/database"
	"github.com/preetigupta1005/ridehail-go/models"
)

func CreateRide(ride *models.Ride) error {
	query := `INSERT INTO rides (passenger_id, pickup_lat, pickup_lng, drop_lat, drop_lng, status)
	          VALUES ($1, $2, $3, $4, $5, 'requested') RETURNING id,status, requested_at`
	return database.DB.QueryRowx(query, ride.PassengerID, ride.PickupLat, ride.PickupLng, ride.DropLat, ride.DropLng).
		Scan(&ride.ID, &ride.Status, &ride.RequestedAt)
}

func GetNearbyDrivers(lat, lng float64, radiusMeters int) ([]string, error) {
	query := `SELECT user_id FROM driver_details
	          WHERE is_available = true
	          AND ST_DWithin(current_location, ST_SetSRID(ST_MakePoint($1, $2), 4326), $3)`
	var driverIDs []string
	err := database.DB.Select(&driverIDs, query, lng, lat, radiusMeters)
	return driverIDs, err
}
