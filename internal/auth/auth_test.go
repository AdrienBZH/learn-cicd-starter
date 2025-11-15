package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeySuccess(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := "my-secret-key"
	if key != expected {
		t.Fatalf("expected %q, got %q", expected, key)
	}
}

func TestGetAPIKeyErrors(t *testing.T) {
	tests := []struct {
		name       string
		headerVal  string
		wantErrMsg string
	}{
		{
			name:       "missing header",
			headerVal:  "",
			wantErrMsg: ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name:       "wrong scheme",
			headerVal:  "Bearer token123",
			wantErrMsg: "malformed authorization header",
		},
		{
			name:       "missing token",
			headerVal:  "ApiKey",
			wantErrMsg: "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.headerVal != "" {
				headers.Set("Authorization", tt.headerVal)
			}

			_, err := GetAPIKey(headers)
			if err == nil {
				t.Fatalf("expected an error but got nil")
			}

			if err.Error() != tt.wantErrMsg {
				t.Fatalf("expected error %q, got %q", tt.wantErrMsg, err.Error())
			}
		})
	}
}
