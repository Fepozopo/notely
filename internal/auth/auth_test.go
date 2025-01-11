package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("with a valid auth header", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "ApiKey 123")
		apiKey, err := GetAPIKey(headers)
		if err != nil {
			t.Errorf("GetAPIKey() returned an error, but it shouldn't: %v", err)
		}
		if apiKey != "123" {
			t.Errorf("GetAPIKey() returned an unexpected api key: %v", apiKey)
		}
	})

	t.Run("with no auth header", func(t *testing.T) {
		headers := http.Header{}
		_, err := GetAPIKey(headers)
		if err == nil {
			t.Errorf("GetAPIKey() did not return an error, but it should")
		}
		if err != ErrNoAuthHeaderIncluded {
			t.Errorf("GetAPIKey() returned an unexpected error: %v", err)
		}
	})
}
