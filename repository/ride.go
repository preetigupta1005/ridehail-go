package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/preetigupta1005/ridehail-go/database"
	"github.com/preetigupta1005/ridehail-go/models"
)

func CreateRide(ride *models.Ride) error {
	query := `INSERT INTO rides (passenger_id, pickup_lat, pickup_lng, pickup_address, drop_lat, drop_lng, drop_address, status)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, 'requested') 
	          RETURNING id, status, requested_at`
	return database.DB.QueryRowx(query, ride.PassengerID, ride.PickupLat, ride.PickupLng, ride.PickupAddress,
		ride.DropLat, ride.DropLng, ride.DropAddress).
		Scan(&ride.ID, &ride.Status, &ride.RequestedAt)
}

func GetNearbyDrivers(lat, lng float64, radiusMeters int) ([]string, error) {
	query := `SELECT user_id FROM driver_details
	          WHERE is_available = true  AND is_on_ride = false
	          AND ST_DWithin(current_location, ST_SetSRID(ST_MakePoint($1, $2), 4326), $3)`
	var driverIDs []string
	err := database.DB.Select(&driverIDs, query, lng, lat, radiusMeters)
	return driverIDs, err
}

func StartRide(rideID, driverID string) error {
	result, err := database.DB.Exec(
		`UPDATE rides SET status='ongoing', started_at=now() 
		 WHERE id=$1 AND driver_id=$2 AND status='accepted'`,
		rideID, driverID,
	)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("ride not found or not in accepted state")
	}
	return nil
}

func EndRide(rideID, driverID string) (float64, float64, error) {
	var pickupLat, pickupLng, dropLat, dropLng float64
	err := database.DB.QueryRowx(
		`SELECT pickup_lat, pickup_lng, drop_lat, drop_lng FROM rides WHERE id=$1`, rideID,
	).Scan(&pickupLat, &pickupLng, &dropLat, &dropLng)
	if err != nil {
		return 0, 0, err
	}

	var distanceMeters float64
	err = database.DB.QueryRowx(
		`SELECT ST_Distance(
			ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography,
			ST_SetSRID(ST_MakePoint($3, $4), 4326)::geography
		)`,
		pickupLng, pickupLat, dropLng, dropLat,
	).Scan(&distanceMeters)
	if err != nil {
		return 0, 0, err
	}

	distanceKm := distanceMeters / 1000
	fare := 30 + (distanceKm * 10)

	result, err := database.DB.Exec(
		`UPDATE rides SET status='completed', completed_at=now(), 
		 fare_amount=$1, distance_km=$2 
		 WHERE id=$3 AND driver_id=$4 AND status='ongoing'`,
		fare, distanceKm, rideID, driverID,
	)
	if err != nil {
		return 0, 0, err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return 0, 0, errors.New("ride not found or not ongoing")
	}

	_, err = database.DB.Exec(`UPDATE driver_details SET is_on_ride=false WHERE user_id=$1`, driverID)
	if err != nil {
		return 0, 0, err
	}

	return fare, distanceKm, nil
}

func CancelRide(rideID, userID, role, reason string) error {
	return database.Tx(func(tx *sqlx.Tx) error {
		var passengerID, driverID *string
		err := tx.QueryRowx(
			`SELECT passenger_id, driver_id FROM rides WHERE id=$1`, rideID,
		).Scan(&passengerID, &driverID)
		if err != nil {
			return err
		}

		if (passengerID == nil || *passengerID != userID) && (driverID == nil || *driverID != userID) {
			return errors.New("not authorized to cancel this ride")
		}

		result, err := tx.Exec(
			`UPDATE rides SET status='cancelled', cancelled_at=now(), cancelled_by=$1 ,cancellation_reason=$2
		 WHERE id=$3 AND status NOT IN ('completed', 'cancelled')`,
			role, reason, rideID,
		)
		if err != nil {
			return err
		}
		rows, _ := result.RowsAffected()
		if rows == 0 {
			return errors.New("ride cannot be cancelled")
		}

		_, err = tx.Exec(
			`UPDATE ride_requests SET status='expired' 
			 WHERE ride_id=$1 AND status='sent'`,
			rideID,
		)
		if err != nil {
			return err
		}

		if driverID != nil {
			_, err = tx.Exec(`UPDATE driver_details SET is_on_ride=false WHERE user_id=$1`, *driverID)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func GetRidesByPassenger(passengerID string) ([]models.Ride, error) {
	var rides []models.Ride
	query := `SELECT * FROM rides WHERE passenger_id=$1 ORDER BY requested_at DESC`
	err := database.DB.Select(&rides, query, passengerID)
	return rides, err
}

func GetRidesByDriver(driverID string) ([]models.Ride, error) {
	var rides []models.Ride
	query := `SELECT * FROM rides WHERE driver_id=$1 ORDER BY requested_at DESC`
	err := database.DB.Select(&rides, query, driverID)
	return rides, err
}
