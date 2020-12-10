package gobernate_test

import (
	"gobernate"

	"net/http"
	"testing"
	"time"
)

var testGobernateCases = []struct {
	name           string
	path           string
	method         string
	expectedStatus int
	expectedMime   string
}{
	{"GET /version", "/version", "GET", http.StatusOK, "application/json"},
	{"POST /version", "/version", "POST", http.StatusMethodNotAllowed, ""},
	{"GET /health", "/health", "GET", http.StatusOK, ""},
	{"GET /readiness", "/readiness", "GET", http.StatusOK, ""},
	{"GET /nonexistent", "/nonexistent", "GET", http.StatusNotFound, ""},
}

func TestGobernate(t *testing.T) {
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 100

	g := gobernate.New("0", "gobernate-version", "1.0.0", "f00b4r", "now")

	shutdown := g.Launch()
	g.Ready()

	time.Sleep(500 * time.Millisecond)

	for _, tt := range testGobernateCases {
		t.Run(tt.name, func(t *testing.T) {

			var res *http.Response
			var err error
			switch tt.method {
			case "GET":
				res, err = http.Get(g.URL() + tt.path)
			case "POST":
				res, err = http.Post(g.URL()+tt.path, "text/plain", nil)
			default:
				t.Fatalf("Unknown method: %s", tt.method)
			}
			if err != nil {
				t.Fatal(err)
			}

			if res.StatusCode != tt.expectedStatus {
				t.Errorf("Status code for %s %s is wrong. Have: %d, want: %d.",
					tt.method, tt.path, res.StatusCode, tt.expectedStatus)
			}
		})
	}

	g.Shutdown()

	select {
	case <-shutdown:
		// expected
		break
	case <-time.After(1 * time.Second):
		t.Errorf("Shutdown took too long. Expected termination within 1 second.")
	}
}
