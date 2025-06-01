package server

import (
	"net/http"

	"github.com/yosa12978/kodama/internal/router"
)

func New(addr string) http.Server {
	router := router.New(nil)
	return http.Server{
		Addr:    addr,
		Handler: router,
	}
}
