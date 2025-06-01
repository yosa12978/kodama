package middleware

import (
	"net"
	"net/http"
	"strings"
)

var (
	xForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")
	trueClientIP  = http.CanonicalHeaderKey("True-Client-IP")
	xRealIP       = http.CanonicalHeaderKey("X-Real-IP")
)

func RealIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if rip := realIP(r); rip != "" {
			r.RemoteAddr = rip
		}
		next.ServeHTTP(w, r)
	})
}

func realIP(r *http.Request) string {
	var ip string

	if tcip := r.Header.Get(trueClientIP); tcip != "" {
		ip = tcip
	} else if xrip := r.Header.Get(xRealIP); xrip != "" {
		ip = xrip
	} else if xff := r.Header.Get(xForwardedFor); xff != "" {
		ip, _, _ = strings.Cut(xff, ",")
	}
	if ip == "" || net.ParseIP(ip) == nil {
		return ""
	}
	return ip
}
