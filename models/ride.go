package models

import "time"

type Ride struct {
	ID                 string     `db:"id" json:"id"`
	PassengerID        string     `db:"passenger_id" json:"passenger_id"`
	DriverID           *string    `db:"driver_id" json:"driver_id,omitempty"`
	PickupLat          float64    `db:"pickup_lat" json:"pickup_lat"`
	PickupLng          float64    `db:"pickup_lng" json:"pickup_lng"`
	PickupAddress      *string    `db:"pickup_address" json:"pickup_address,omitempty"`
	DropLat            float64    `db:"drop_lat" json:"drop_lat"`
	DropLng            float64    `db:"drop_lng" json:"drop_lng"`
	DropAddress        *string    `db:"drop_address" json:"drop_address,omitempty"`
	Status             string     `db:"status" json:"status"`
	FareAmount         *float64   `db:"fare_amount" json:"fare_amount,omitempty"`
	DistanceKm         *float64   `db:"distance_km" json:"distance_km,omitempty"`
	RequestedAt        time.Time  `db:"requested_at" json:"requested_at"`
	AcceptedAt         *time.Time `db:"accepted_at" json:"accepted_at,omitempty"`
	StartedAt          *time.Time `db:"started_at" json:"started_at,omitempty"`
	CompletedAt        *time.Time `db:"completed_at" json:"completed_at,omitempty"`
	CancelledAt        *time.Time `db:"cancelled_at" json:"cancelled_at,omitempty"`
	CancelledBy        *string    `db:"cancelled_by" json:"cancelled_by,omitempty"`
	CancellationReason *string    `db:"cancellation_reason" json:"cancellation_reason,omitempty"`
}

type CancelRideBody struct {
	Reason string `json:"reason,omitempty"`
}

type ActiveRideLocation struct {
	RideID string  `db:"ride_id"`
	Lat    float64 `db:"lat"`
	Lng    float64 `db:"lng"`
}
