package router

import (
	"net/http"

	"github.com/yosa12978/kodama/assets"
	"github.com/yosa12978/kodama/internal/middleware"
	"github.com/yosa12978/kodama/internal/templates"
)

type RouterOptions struct {
}

func defaultRouterOptions() *RouterOptions {
	return &RouterOptions{}
}

func New(opts *RouterOptions) http.Handler {
	// options := defaultRouterOptions()
	// if opts != nil {
	// 	options = opts
	// }

	router := http.NewServeMux()

	router.Handle("/assets/", http.FileServerFS(assets.AssetsFS))

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		templates.IndexTemplate.Execute(w, nil)
	})

	router.HandleFunc("GET /panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic test")
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		templates.ErrorTemplate.Execute(w,
			templates.ErrorPayload{
				StatusCode: http.StatusNotFound,
				Message:    "Page Not Found",
			},
		)
	})

	handler := middleware.Chain(
		router,
		middleware.Logger,
		middleware.RealIP,
		middleware.StripSlash,
		middleware.Recovery,
	)
	return handler
}
