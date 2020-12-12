package handlers_test

import (
	"github.com/SebastiaanPasterkamp/gobernate/handlers"
	"github.com/SebastiaanPasterkamp/gobernate/version"

	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

var testRouterCases = []struct {
	name           string
	path           string
	method         string
	expectedStatus int
	expectedMime   string
}{
	{"GET /version", "/version", "GET", http.StatusOK, "application/json"},
	{"POST /version", "/version", "POST", http.StatusMethodNotAllowed, ""},
	{"GET /health", "/health", "GET", http.StatusOK, ""},
	{"POST /health", "/health", "POST", http.StatusMethodNotAllowed, ""},
	{"GET /readiness", "/readiness", "GET", http.StatusOK, ""},
	{"POST /readiness", "/readiness", "POST", http.StatusMethodNotAllowed, ""},
	{"GET /metrics", "/metrics", "GET", http.StatusOK, "text/plain; version=0.0.4; charset=utf-8"},
	{"POST /metrics", "/metrics", "POST", http.StatusMethodNotAllowed, ""},
	{"GET /nonexistent", "/nonexistent", "GET", http.StatusNotFound, "text/plain; charset=utf-8"},
}

func TestRouter(t *testing.T) {
	isReady := &atomic.Value{}
	isReady.Store(true)
	shutdown := make(chan bool)

	r := handlers.Router(version.Info{
		Name:      "router-version",
		Release:   "1.0.0",
		Commit:    "f00b4r",
		BuildTime: "now",
	}, isReady, shutdown)
	ts := httptest.NewServer(r)
	defer ts.Close()

	for _, tt := range testRouterCases {
		t.Run(tt.name, func(t *testing.T) {
			var res *http.Response
			var err error
			switch tt.method {
			case "GET":
				res, err = http.Get(ts.URL + tt.path)
			case "POST":
				res, err = http.Post(ts.URL+tt.path, "text/plain", nil)
			default:
				t.Fatalf("Unknown method: %s", tt.method)
			}
			if err != nil {
				t.Fatal(err)
			}

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("Status code for %s %s is wrong. Have: %d, want: %d.",
					tt.method, tt.path, res.StatusCode, http.StatusOK)
			}

			if ctype := res.Header.Get("Content-Type"); ctype != tt.expectedMime {
				t.Errorf("Content-Type for %s %s is wrong. Have: '%s', want: '%s'.",
					tt.method, tt.path, ctype, tt.expectedMime)
			}
		})
	}
}
