package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/preetigupta1005/ridehail-go/database"
)

func CreateDriverDetails(tx *sqlx.Tx, userID, vehicleNumber, vehicleType, licenseNumber string) error {
	query := `INSERT INTO driver_details (user_id, vehicle_number, vehicle_type, license_number) 
	          VALUES ($1, $2, $3, $4)`
	_, err := tx.Exec(query, userID, vehicleNumber, vehicleType, licenseNumber)
	return err
}

func UpdateAvailability(userID string, isAvailable bool) error {
	query := `UPDATE driver_details SET is_available=$1, updated_at=now() WHERE user_id=$2`
	_, err := database.DB.Exec(query, isAvailable, userID)
	return err
}

func UpdateLocation(userID string, lat, lng float64) error {
	query := `UPDATE driver_details 
	          SET current_location = ST_SetSRID(ST_MakePoint($1, $2), 4326),
	              last_location_update = now()
	          WHERE user_id=$3`
	_, err := database.DB.Exec(query, lng, lat, userID)
	return err
}
