package server

import (
	"context"
	"net/http"
	"time"

	"github.com/preetigupta1005/ridehail-go/handlers"
	"github.com/preetigupta1005/ridehail-go/middlewares"
)

type Server struct {
	Mux    *http.ServeMux
	server *http.Server
}

const (
	readTimeout       = 5 * time.Minute
	readHeaderTimeout = 30 * time.Second
	writeTimeout      = 5 * time.Minute
)

func SetUpRoutes() *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/signup", handlers.SignupHandler)
	mux.HandleFunc("POST /v1/login", handlers.LoginHandler)

	mux.Handle("PATCH /v1/driver/availability", middlewares.DriverOnly(handlers.UpdateAvailabilityHandler))
	mux.Handle("PATCH /v1/driver/location", middlewares.DriverOnly(handlers.UpdateLocationHandler))

	mux.Handle("POST /v1/rides/request", middlewares.PassengerOnly(handlers.RequestRideHandler))

	mux.Handle("POST /v1/rides/{id}/accept", middlewares.DriverOnly(handlers.AcceptRideHandler))

	mux.Handle("POST /v1/rides/{id}/start", middlewares.DriverOnly(handlers.StartRideHandler))
	mux.Handle("POST /v1/rides/{id}/end", middlewares.DriverOnly(handlers.EndRideHandler))
	mux.Handle("POST /v1/rides/{id}/cancel", middlewares.AuthOnly(handlers.CancelRideHandler))

	mux.Handle("GET /rides/activity", middlewares.AuthOnly(handlers.GetMyRidesHandler))

	mux.Handle("GET /v1/rides/{id}/location", middlewares.AuthOnly(handlers.GetRideLocationHandler))

	return &Server{Mux: mux}
}
func (svc *Server) Run(port string) error {
	svc.server = &http.Server{
		Addr:              port,
		Handler:           svc.Mux,
		ReadTimeout:       readTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
	}
	return svc.server.ListenAndServe()
}

func (svc *Server) Shutdown(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return svc.server.Shutdown(ctx)
}
