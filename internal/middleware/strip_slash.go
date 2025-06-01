package middleware

import (
	"net/http"
	"regexp"
)

var pathRegex = regexp.MustCompile(`^(?P<path>(?:/[^/\n\t\s]+)+?)/*$`)

func StripSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = pathRegex.ReplaceAllString(r.URL.Path, `$path`)
		next.ServeHTTP(w, r)
	})
}
