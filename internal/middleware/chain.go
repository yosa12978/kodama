package middleware

import "net/http"

func Chain(f http.Handler, m ...Middleware) http.Handler {
	for _, middleware := range m {
		f = middleware(f)
	}
	return f
}
