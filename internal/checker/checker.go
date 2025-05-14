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
// CheckAll runs checks concurrently and returns all results
func CheckAll(urls []string) []Result {
	results := make([]Result, 0, len(urls))
	ch := make(chan Result, len(urls)) // ✅ Buffered channel

	// Launch goroutines to check URLs concurrently
	for _, url := range urls {
		go func(u string) {
			ch <- CheckWithRetry(u, 3)
		}(url)
	}

	// Collect results from the channel
	for i := 0; i < len(urls); i++ {
		results = append(results, <-ch)
	}

	return results
}

// Check with retires after a timeout

func CheckWithRetry(url string, attempts int) Result {

	var lastError error

	//the loop will run for the number of attempts
	// and will retry the request if it fails
	// it will return the result of the last attempt
	for i := 0; i < attempts; i++ { // not for range because attempts is an int not a slice or channel and range cannot iterate over an int

		result := CheckURL(url)

		if result.Error == nil {
			return result // Return the successful result
		}

		lastError = result.Error
		time.Sleep(2 * time.Second) // Wait before retrying
	}
	return Result{
		URL:   url,
		Error: lastError,
	} // Return the last error if all attempts fail

}
