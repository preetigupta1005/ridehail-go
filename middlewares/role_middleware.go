package middlewares

import "net/http"

func RequireRole(allowedRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role, ok := r.Context().Value(RoleKey).(string)
			if !ok || role != allowedRole {
				http.Error(w, "forbidden: not allowed for this role", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
