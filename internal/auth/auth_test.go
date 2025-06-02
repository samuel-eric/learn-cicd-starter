package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	apiKey := "helloWorld"

	header1 := http.Header{}
	header1.Add("Authorization", fmt.Sprintf("ApiKey %s", apiKey))

	header2 := http.Header{}

	header3 := http.Header{}
	header3.Add("Authorization", apiKey)

	type test struct {
		header  http.Header
		apiKey  string
		wantErr bool
	}

	tests := []test{
		{
			header:  header1,
			apiKey:  apiKey,
			wantErr: false,
		},
		{
			header:  header2,
			apiKey:  apiKey,
			wantErr: true,
		},
		{
			header:  header3,
			apiKey:  apiKey,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		resultApiKey, err := GetAPIKey(tt.header)
		if err != nil {
			if !tt.wantErr {
				t.Errorf("unexpected error: %v", err)
				continue
			}
		} else {
			if resultApiKey != tt.apiKey {
				t.Errorf("resultApiKey (%v) != apiKey (%v)", resultApiKey, tt.apiKey)
				continue
			}
		}
	}
}
