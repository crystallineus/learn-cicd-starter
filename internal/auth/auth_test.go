package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectedErr error
	}{
		{
			name:        "Valid Authorization Header",
			headers:     http.Header{"Authorization": []string{"ApiKey abc123"}},
			expectedKey: "abc123",
			expectedErr: nil,
		},
		{
			name:        "Missing Authorization Header",
			headers:     http.Header{},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:        "Empty Authorization Header",
			headers:     http.Header{"Authorization": []string{""}},
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)

			// Check if the returned API key matches the expected value
			if apiKey != tt.expectedKey {
				t.Errorf("expected API key %q, got %q", tt.expectedKey, apiKey)
			}

			// Check if the error matches the expected error
			if (err == nil && tt.expectedErr != nil) || (err != nil && tt.expectedErr == nil) || (err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error()) {
				t.Errorf("expected error %q, got %q", tt.expectedErr, err)
			}
		})
	}
}
