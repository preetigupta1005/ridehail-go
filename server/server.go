package server

import (
	"net/http"
)

func SetUpRoutes() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /signup", signupHandler)
	mux.HandleFunc("POST /login", loginHandler)

	mux.Handle("POST /rides/request", passengerOnly(requestRideHandler))

	mux.Handle("POST /rides/{id}/accept", driverOnly(acceptRideHandler))
	mux.Handle("POST /rides/{id}/start", driverOnly(startRideHandler))

	mux.Handle("POST /rides/{id}/cancel", authOnly(cancelRideHandler))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	return srv
}
