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

	mux.HandleFunc("POST /signup", handlers.SignupHandler)
	mux.HandleFunc("POST /login", handlers.LoginHandler)

	mux.Handle("PATCH /driver/availability", middlewares.DriverOnly(handlers.UpdateAvailabilityHandler))
	mux.Handle("PATCH /driver/location", middlewares.DriverOnly(handlers.UpdateLocationHandler))

	mux.Handle("POST /v1/rides/request", middlewares.PassengerOnly(handlers.RequestRideHandler))

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
