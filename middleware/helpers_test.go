package middleware_test

import (
	"net/http"
)

func dummyHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
