package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Define a table of test cases
	tests := []struct {
		name          string      // The name of the test case
		headers       http.Header // The input headers for the function
		expectedKey   string      // The expected API key to be returned
		expectedError error       // The expected error to be returned
	}{
		{
			name: "Valid ApiKey",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-api-key"},
			},
			expectedKey:   "my-secret-api-key",
			expectedError: nil,
		},
		{
			name:          "No Authorization Header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Header - Wrong Scheme",
			headers: http.Header{
				"Authorization": []string{"Bearer my-secret-api-key"},
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "Malformed Header - No Space",
			headers: http.Header{
				"Authorization": []string{"ApiKeymy-secret-api-key"},
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
	}

	// Iterate over the test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function under test
			apiKey, err := GetAPIKey(tc.headers)

			// Assert the returned API key is what we expect
			if apiKey != tc.expectedKey {
				t.Errorf("expected key %q, but got %q", tc.expectedKey, apiKey)
			}

			// Assert the returned error is what we expect
			if (err != nil && tc.expectedError == nil) || (err == nil && tc.expectedError != nil) || (err != nil && tc.expectedError != nil && err.Error() != tc.expectedError.Error()) {
				t.Errorf("expected error %v, but got %v", tc.expectedError, err)
			}
		})
	}
}
