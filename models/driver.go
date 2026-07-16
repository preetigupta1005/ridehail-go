package models

type AvailabilityRequest struct {
	IsAvailable bool `json:"is_available"`
}
type LocationRequest struct {
	Lat float64 `json:"lat" validate:"required,latitude"`
	Lng float64 `json:"lng" validate:"required,longitude"`
}
