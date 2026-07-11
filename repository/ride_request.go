package repository

import "github.com/preetigupta1005/ridehail-go/database"

func CreateRideRequests(rideID string, driverIDs []string) error {
	for _, driverID := range driverIDs {
		_, err := database.DB.Exec(
			`INSERT INTO ride_requests (ride_id, driver_id, status) VALUES ($1, $2, 'sent')`,
			rideID, driverID,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
