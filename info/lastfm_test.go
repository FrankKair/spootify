package info

import (
	"net/url"
	"testing"

	"github.com/FrankKair/spootify/track"
)

func TestGetLastfmURL(t *testing.T) {
	tests := []struct {
		name   string
		artist string
		album  string
		apiKey string
		check func(t *testing.T, rawURL string)
	}{
		{
			name:   "basic artist and album",
			artist: "Led Zeppelin",
			album:  "Led Zeppelin IV",
			apiKey: "test-key",
			check: func(t *testing.T, rawURL string) {
				u, err := url.Parse(rawURL)
				if err != nil {
					t.Fatalf("invalid URL: %v", err)
				}
				q := u.Query()
				if got := q.Get("artist"); got != "Led Zeppelin" {
					t.Errorf("artist = %q, want %q", got, "Led Zeppelin")
				}
				if got := q.Get("album"); got != "Led Zeppelin IV" {
					t.Errorf("album = %q, want %q", got, "Led Zeppelin IV")
				}
				if got := q.Get("api_key"); got != "test-key" {
					t.Errorf("api_key = %q, want %q", got, "test-key")
				}
				if got := q.Get("format"); got != "json" {
					t.Errorf("artist = %q, want %q", got, "json")
				}
			},
		},
		{
			name:   "special characters in album",
			artist: "Justin Timberlake",
			album:  "The 20/20 Experience",
			apiKey: "key123",
			check: func(t *testing.T, rawURL string) {
				u, err := url.Parse(rawURL)
				if err != nil {
					t.Fatalf("invalid URL: %v", err)
				}
				if got := u.Query().Get("album"); got != "The 20/20 Experience" {
					t.Errorf("album = %q, want %q", got, "The 20/20 Experience")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := track.New(tt.artist, tt.album, "SomeTitle")
			u := getLastfmURL(tr, tt.apiKey)
			tt.check(t, u)
		})
	}
}
