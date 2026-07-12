package models

type RequestRideBody struct {
	PickupLat     float64 `json:"pickup_lat"`
	PickupLng     float64 `json:"pickup_lng"`
	PickupAddress string  `json:"pickup_address"`
	DropLat       float64 `json:"drop_lat"`
	DropLng       float64 `json:"drop_lng"`
	DropAddress   string  `json:"drop_address"`
}
