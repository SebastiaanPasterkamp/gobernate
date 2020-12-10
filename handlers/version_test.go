package handlers

import (
	"gobernate/version"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersion(t *testing.T) {
	expected := version.Info{
		Name:      "test-version",
		Release:   "1.0.0",
		Commit:    "f00b4r",
		BuildTime: "now",
	}

	w := httptest.NewRecorder()
	versionHandler(expected)(w, nil)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", have, want)
	}

	result := version.Info{}
	dec := json.NewDecoder(resp.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&result); err != nil {
		t.Error(err)
	}
	resp.Body.Close()

	if result.Name != expected.Name {
		t.Errorf("Name is wrong. Have: %s, want: %s.", result.Name, expected.Name)
	}
	if result.Release != expected.Release {
		t.Errorf("Release is wrong. Have: %s, want: %s.", result.Release, expected.Release)
	}
	if result.Commit != expected.Commit {
		t.Errorf("Commit is wrong. Have: %s, want: %s.", result.Commit, expected.Commit)
	}
	if result.BuildTime != expected.BuildTime {
		t.Errorf("BuildTime is wrong. Have: %s, want: %s.", result.BuildTime, expected.BuildTime)
	}
}
