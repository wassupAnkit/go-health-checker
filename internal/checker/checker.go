package checker

import (
	"net/http"
	"time"
)

// Result holds the result of a health check
type Result struct {
	URL    string
	Alive  bool
	Status int
	Err    error
}

// Check performs a health check on one URL
func Check(url string) Result {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return Result{URL: url, Alive: false, Err: err}
	}
	defer resp.Body.Close()

	alive := resp.StatusCode >= 200 && resp.StatusCode < 400

	return Result{
		URL:    url,
		Alive:  alive,
		Status: resp.StatusCode,
		Err:    nil,
	}
}
