package handlers

import (
	"gobernate/version"
	"sync/atomic"

	"github.com/gorilla/mux"
)

// Router register necessary routes and returns an instance of a router.
func Router(info version.Info, isReady *atomic.Value, shutdown chan bool) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/version", versionHandler(info)).
		Methods("GET")
	r.HandleFunc("/health", healthHandler).
		Methods("GET")
	r.HandleFunc("/readiness", readinessHandler(isReady, shutdown)).
		Methods("GET")
	return r
}
