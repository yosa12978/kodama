package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type WrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *WrappedWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()
		writer := &WrappedWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(writer, r)
		latencyMicro := time.Since(start).Microseconds()
		slog.Info(
			"incoming request",
			"latency_us", latencyMicro,
			"method", r.Method,
			"path", r.URL.Path,
			"ip", r.RemoteAddr,
			//"user_agent", r.UserAgent(),
			"status_code", writer.statusCode,
		)
	})
}
