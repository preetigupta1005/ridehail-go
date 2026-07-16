package services

import (
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

	if err := repository.CreateRide(ride); err != nil {
		return nil, err
	}

	driverIDs, err := repository.GetNearbyDrivers(pickupLat, pickupLng, 5000)
	if err != nil {
		return nil, err
	}

	if len(driverIDs) == 0 {
		return ride, nil
	}

	if err := repository.CreateRideRequests(ride.ID, driverIDs); err != nil {
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
