package utils

import (
	"net/http"

	"github.com/preetigupta1005/ridehail-go/middlewares"
)

func driverOnly(h http.HandlerFunc) http.Handler {
	return middlewares.AuthMiddleware(
		middlewares.RequireRole("driver")(
			http.HandlerFunc(h),
		),
	)
}

func passengerOnly(h http.HandlerFunc) http.Handler {
	return middlewares.AuthMiddleware(
		middlewares.RequireRole("passenger")(
			http.HandlerFunc(h),
		),
	)
}

func authOnly(h http.HandlerFunc) http.Handler {
	return middlewares.AuthMiddleware(http.HandlerFunc(h))
}
