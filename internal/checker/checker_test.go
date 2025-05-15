package checker

import (
	"testing"
)

// TestCheckURL tests the CheckURL function
func TestCheckURL(t *testing.T) {
	result := CheckURL("https://httpbin.org/status/200")
	if result.Error != nil {
		t.Errorf("Expected 200 OK, got %v and error %v", result.StatusCode, result.Error)
	}
}

func TestCheckURLFailure(t *testing.T) {
	result := CheckURL("https://random-url-that-does-not-exist.com")
	if result.Error == nil {
		t.Error("Expected an error for invalid URL, but got nil")
	}
}
