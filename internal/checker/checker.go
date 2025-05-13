package checker

import (
	"net/http"
	"time"
)

// Result holds the outcome of checking a single URL
type Result struct {
	URL        string
	StatusCode int
	Error      error
}

// CheckURL performs an HTTP GET with a timeout and returns the result
func CheckURL(url string) Result {
	client := http.Client{
		Timeout: 5 * time.Second, // timeout for requests
	}

	resp, err := client.Get(url)
	if err != nil {
		return Result{URL: url, Error: err}
	}
	defer resp.Body.Close()

	return Result{
		URL:        url,
		StatusCode: resp.StatusCode,
	}
}

// CheckAll runs checks concurrently and returns all results
func CheckAll(urls []string) []Result {
	results := make([]Result, 0, len(urls))
	ch := make(chan Result) // Create a channel to collect results

	// Loop through URLs and check them concurrently using goroutines
	for _, url := range urls {
		go func(u string) {
			ch <- CheckURL(u) // Send the result of CheckURL to the channel
		}(url)
	}

	// Collect all results from the channel
	for i := 0; i < len(urls); i++ {
		results = append(results, <-ch) // Receive the result and append to the results slice
	}

	return results
}
