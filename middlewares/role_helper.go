package middlewares

import "net/http"

func DriverOnly(h http.HandlerFunc) http.Handler {
	return AuthMiddleware(
		RequireRole("driver")(
			http.HandlerFunc(h),
		),
	)
}

func PassengerOnly(h http.HandlerFunc) http.Handler {
	return AuthMiddleware(
		RequireRole("passenger")(
			http.HandlerFunc(h),
		),
	)
}

func AuthOnly(h http.HandlerFunc) http.Handler {
	return AuthMiddleware(http.HandlerFunc(h))
}
