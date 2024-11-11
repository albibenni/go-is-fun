package main

import (
	"fmt"
	"io"
	"net/http"
)

// getIssueData fetches issue data from the specified API endpoint.
// It sends a GET request to the API and returns the response body as a byte slice.
// If an error occurs during the request or reading the response body, it returns the error.
//
// Returns:
//   - []byte: The response body from the API.
//   - error: An error if one occurred during the request or reading the response body.
func getIssueData() ([]byte, error) {
	res, err := http.Get("https://api.boot.dev/v1/courses_rest_api/learn-http/issues")
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	return data, err
}

// - http.Get uses the http.DefaultClient to make a request to the given url
// - res is the HTTP response that comes back from the server
// - defer res.Body.Close() ensures that the response body is properly closed after reading. Not doing so can cause memory issues.
// - io.ReadAll reads the response body into a slice of bytes []byte called data
func getProjects() ([]byte, error) {
	res, err := http.Get("https://api.jello.com/projects")
	if err != nil {
		return []byte{}, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading response: %w", err)
	}
	return data, nil
}
