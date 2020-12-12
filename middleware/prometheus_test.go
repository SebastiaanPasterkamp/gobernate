package middleware_test

import (
	mw "github.com/SebastiaanPasterkamp/gobernate/middleware"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestPrometheusMiddleware(t *testing.T) {
	mr := mux.NewRouter()
	mr.Use(mw.PrometheusMiddleware("test"))
	mr.HandleFunc("/", dummyHandler).Methods("GET")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	mr.ServeHTTP(w, r)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.",
			res.StatusCode, http.StatusOK)
	}
}
