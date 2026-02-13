// Package utils provides REST API client utilities for making HTTP requests
// and decoding JSON responses. These functions handle API endpoint calls
// and generic JSON deserialization used throughout the ETL pipeline.
package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CallEndpoint performs an HTTP GET request to the specified endpoint URL
// and returns the response body as a byte slice. It validates that the
// response status code is 200 and returns an error for any network failures,
// non-200 status codes, or issues reading the response body.
//
// Parameters:
//   - endpoint: The full URL of the API endpoint to call
//
// Returns:
//   - []byte: The raw response body bytes, or nil on error
//   - error: An error describing any failures during the request or response processing
func CallEndpoint(endpoint string) ([]byte, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error occurred calling endpoint: %s: \n%s", endpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("non 200 response code returned %s:\n%d", endpoint, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error occurred parsing API Response: \n%s", err)
	}

	return body, nil
}

// DecodeJSON unmarshals JSON data from a byte slice into a value of type T.
// It is a generic function that accepts any type T and attempts to decode the
// JSON bytes into that type. If the JSON cannot be decoded into the target type,
// an error is returned.
//
// Type Parameters:
//   - T: The target type to decode the JSON into (must be a concrete type)
//
// Parameters:
//   - body: The JSON data as a byte slice to decode
//
// Returns:
//   - T: The decoded value of type T, or the zero value of T on error
//   - error: An error describing any failures during JSON unmarshaling
func DecodeJSON[T any](body []byte) (T, error) {
	var decodedJsonBody T

	err := json.Unmarshal([]byte(body), &decodedJsonBody)
	if err != nil {
		return decodedJsonBody, fmt.Errorf("error occurred decoding json response:\n%s", err)
	}
	return decodedJsonBody, nil
}
