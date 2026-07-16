package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/preetigupta1005/ridehail-go/database"
	"github.com/preetigupta1005/ridehail-go/models"
	"github.com/preetigupta1005/ridehail-go/repository"
)

func RequestRide(passengerID string, pickupLat, pickupLng float64, pickupAddr string, dropLat, dropLng float64, dropAddr string) (*models.Ride, error) {
	ride := &models.Ride{
		PassengerID:   passengerID,
		PickupLat:     pickupLat,
		PickupLng:     pickupLng,
		PickupAddress: &pickupAddr,
		DropLat:       dropLat,
		DropLng:       dropLng,
		DropAddress:   &dropAddr,
	}

	err := database.Tx(func(tx *sqlx.Tx) error {

		if err := repository.CreateRide(tx, ride); err != nil {
			return err
		}

		driverIDs, err := repository.GetNearbyDrivers(tx, pickupLat, pickupLng, 5000)
		if err != nil {
			return err
		}

		if len(driverIDs) > 0 {
			if err := repository.CreateRideRequests(tx, ride.ID, driverIDs); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return ride, nil
}

func GetMyRides(userID, role string) ([]models.Ride, error) {
	if role == "passenger" {
		return repository.GetRidesByPassenger(userID)
	}
	return repository.GetRidesByDriver(userID)
}
