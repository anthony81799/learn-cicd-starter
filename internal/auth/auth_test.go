package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testHeader := http.Header{}
	testHeader["Authorization"] = []string{"ApiKey", "data"}

	got, err := GetAPIKey(testHeader)
	want := []string{"data", err.Error()}
	if got != want[0] && err.Error() != want[1] {
		t.Fatalf("expected: %v, got: %v", want, []string{got, err.Error()})
	}
}

func TestGetAPIKeyError(t *testing.T) {
	testHeader := http.Header{}
	testHeader["Authorization"] = []string{"test", "data"}

	got, err := GetAPIKey(testHeader)
	want := []string{"", "malformed authorization header"}
	if got != want[0] && err.Error() != want[1] {
		t.Fatalf("expected: %v, got: %v", want, []string{got, err.Error()})
	}
}
