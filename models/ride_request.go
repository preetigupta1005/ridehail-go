package models

type RequestRideBody struct {
	PickupLat     float64 `json:"pickup_lat" validate:"required,latitude"`
	PickupLng     float64 `json:"pickup_lng" validate:"required,longitude"`
	PickupAddress string  `json:"pickup_address"`
	DropLat       float64 `json:"drop_lat" validate:"required,latitude"`
	DropLng       float64 `json:"drop_lng" validate:"required,longitude"`
	DropAddress   string  `json:"drop_address"`
}
