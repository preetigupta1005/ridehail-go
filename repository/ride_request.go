package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/preetigupta1005/ridehail-go/database"
)

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

func AcceptRideRequest(rideID, driverID string) error {
	return database.Tx(func(tx *sqlx.Tx) error {

		var currentStatus string
		err := tx.Get(&currentStatus, `SELECT status FROM rides WHERE id=$1 FOR UPDATE`, rideID)
		if err != nil {
			return err
		}
		if currentStatus != "requested" {
			return errors.New("ride already accepted or not available")
		}

		_, err = tx.Exec(
			`UPDATE ride_requests SET status='accepted', responded_at=now() 
			 WHERE ride_id=$1 AND driver_id=$2`,
			rideID, driverID,
		)
		if err != nil {
			return err
		}

		_, err = tx.Exec(
			`UPDATE ride_requests SET status='expired' 
			 WHERE ride_id=$1 AND driver_id!=$2 AND status='sent'`,
			rideID, driverID,
		)
		if err != nil {
			return err
		}

		_, err = tx.Exec(
			`UPDATE rides SET driver_id=$1, status='accepted', accepted_at=now() WHERE id=$2`,
			driverID, rideID,
		)
		return err
	})
}
