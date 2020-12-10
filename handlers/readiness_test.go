package handlers

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

var testReadinessCases = []struct {
	name           string
	isReady        bool
	shutdown       bool
	expectedStatus int
	expectedMime   string
}{
	{"Not ready", false, false, http.StatusServiceUnavailable, "text/plain; charset=utf-8"},
	{"Ready", true, false, http.StatusOK, ""},
	{"Shutting down", false, false, http.StatusServiceUnavailable, "text/plain; charset=utf-8"},
}

func TestReadiness(t *testing.T) {
	isReady := &atomic.Value{}

	for _, tt := range testReadinessCases {
		t.Run(tt.name, func(t *testing.T) {
			isReady.Store(tt.isReady)
			shutdown := make(chan bool)
			if tt.shutdown {
				close(shutdown)
			}

			w := httptest.NewRecorder()
			readinessHandler(isReady, shutdown)(w, nil)

			res := w.Result()
			if res.StatusCode != tt.expectedStatus {
				t.Errorf("Status code is wrong. Have: %d, want: %d.",
					res.StatusCode, tt.expectedStatus)
			}

			if ctype := res.Header.Get("Content-Type"); ctype != tt.expectedMime {
				t.Errorf("Content-Type is wrong. Have: '%s', want: '%s'.",
					ctype, tt.expectedMime)
			}
		})
	}
}
