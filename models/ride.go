package models

import "time"

type Ride struct {
	ID          string    `db:"id" json:"id"`
	PassengerID string    `db:"passenger_id" json:"passenger_id"`
	DriverID    *string   `db:"driver_id" json:"driver_id,omitempty"`
	PickupLat   float64   `db:"pickup_lat" json:"pickup_lat"`
	PickupLng   float64   `db:"pickup_lng" json:"pickup_lng"`
	DropLat     float64   `db:"drop_lat" json:"drop_lat"`
	DropLng     float64   `db:"drop_lng" json:"drop_lng"`
	Status      string    `db:"status" json:"status"`
	RequestedAt time.Time `db:"requested_at" json:"requested_at"`
}
