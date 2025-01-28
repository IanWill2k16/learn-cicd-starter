package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name: "malformedHeader",
			headers: func() http.Header {
				h := make(http.Header)
				h.Set("Authorization", "womp_womp")
				return h
			}(),
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name: "noAuthHeader",
			headers: func() http.Header {
				h := make(http.Header)
				return h
			}(),
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "validAPIKey",
			headers: func() http.Header {
				h := make(http.Header)
				h.Set("Authorization", "ApiKey abc123")
				return h
			}(),
			want:    "abc123",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if tt.wantErr != nil {
				if err != nil && got != tt.want || err.Error() != tt.wantErr.Error() {
					t.Fatalf("expected: %v, got: %v", tt.wantErr, err)
				}
				return
			}
			if got != tt.want {
				t.Fatalf("expected: %v, got: %v", tt.want, got)
			}
		})
	}
}
