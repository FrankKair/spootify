package info

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/FrankKair/spootify/track"
)

func TestGetLastfmURL(t *testing.T) {
	tests := []struct {
		name   string
		artist string
		album  string
		apiKey string
		check  func(t *testing.T, rawURL string)
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

// Intercept requests without contacting an external API.
type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

type trackedBody struct {
	io.Reader
	closed bool
}

func (b *trackedBody) Close() error { b.closed = true; return nil }

func TestResponses(t *testing.T) {
	original := httpClient
	t.Cleanup(func() { httpClient = original })
	for _, tt := range []struct {
		name    string
		status  int
		body    string
		want    string
		wantErr string
	}{
		{"success with HTML entities", 200, `{"album":{"wiki":{"summary":"A &amp; B <a href='x'>Read more</a>"}}}`, "A & B Read more", ""},
		{"API error", 200, `{"error":6,"message":"Album not found"}`, "", "Album not found"},
		{"missing summary", 200, `{"album":{}}`, "", "no album info"},
		{"empty summary", 200, `{"album":{"wiki":{"summary":" <b></b> "}}}`, "", "no album info"},
		{"malformed JSON", 200, `{`, "", "could not parse"},
		{"HTTP failure", 503, `unavailable`, "", "status 503"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			body := &trackedBody{Reader: strings.NewReader(tt.body)}
			httpClient = &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if r.Header.Get("User-Agent") != userAgent || r.URL.Query().Get("album") != "Album" || r.URL.Query().Get("api_key") != "key" {
					t.Errorf("unexpected request: %s, headers %v", r.URL, r.Header)
				}
				return &http.Response{StatusCode: tt.status, Body: body, Header: make(http.Header)}, nil
			})}
			got, err := Get(track.New("Artist", "Album (Deluxe Edition)", "Song (Remastered)"), "key")
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("got error %v, want %q", err, tt.wantErr)
				}
			} else if err != nil || !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %#v, %v; want %#v", got, err, tt.want)
			}
			if !body.closed {
				t.Error("response body was not closed")
			}
		})
	}
}

func TestTransportError(t *testing.T) {
	original := httpClient
	t.Cleanup(func() { httpClient = original })
	failure := errors.New("connection failed")
	httpClient = &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return nil, failure
	})}
	_, err := Get(track.New("Artist", "Album (Deluxe Edition)", "Song (Remastered)"), "key")
	if !errors.Is(err, failure) {
		t.Fatalf("expected wrapped transport error, got %v", err)
	}
}
