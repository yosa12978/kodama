package middleware

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/yosa12978/kodama/internal/templates"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error(
					"panic recovery",
					"error", fmt.Sprintf("%v", err),
					"endpoint", r.URL.Path,
					"method", r.Method,
				)
				w.WriteHeader(http.StatusInternalServerError)
				templates.ErrorTemplate.Execute(
					w,
					templates.ErrorPayload{
						StatusCode: http.StatusInternalServerError,
						Message:    "Internal Server Error",
					},
				)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}
