package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Call a given API endpoint and read response
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

// Generic that accepts and decodes the Body of a JSON Response and returns the "T" type output
func DecodeJSON[T any](body []byte) (T, error) {
	var decodedJsonBody T

	err := json.Unmarshal([]byte(body), &decodedJsonBody)
	if err != nil {
		return decodedJsonBody, fmt.Errorf("error occurred decoding json response:\n%s", err)
	}
	return decodedJsonBody, nil
}
