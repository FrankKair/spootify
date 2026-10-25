package lyrics

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

func TestGetLrclibURL(t *testing.T) {
	tests := []struct {
		name       string
		artist     string
		album      string
		title      string
		wantArtist string
		wantAlbum  string
		wantTitle  string
	}{
		{
			name:       "basic",
			artist:     "Led Zeppelin",
			album:      "Led Zeppelin IV",
			title:      "Stairway to Heaven",
			wantArtist: "Led Zeppelin",
			wantAlbum:  "Led Zeppelin IV",
			wantTitle:  "Stairway to Heaven",
		},
		{
			name:       "special characters",
			artist:     "AC/DC",
			album:      "High Voltage",
			title:      "It's a Long Way to the Top",
			wantArtist: "AC/DC",
			wantAlbum:  "High Voltage",
			wantTitle:  "It's a Long Way to the Top",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := track.New(tt.artist, tt.album, tt.title)
			rawURL := getLrclibURL(tr)

			u, err := url.Parse(rawURL)
			if err != nil {
				t.Fatalf("invalid URL: %v", err)
			}
			if got := u.Query().Get("artist_name"); got != tt.wantArtist {
				t.Errorf("artist_name = %q, want %q", got, tt.wantArtist)
			}
			if got := u.Query().Get("album_name"); got != tt.wantAlbum {
				t.Errorf("album_name = %q, want %q", got, tt.wantAlbum)
			}
			if got := u.Query().Get("track_name"); got != tt.wantTitle {
				t.Errorf("track_name = %q, want %q", got, tt.wantTitle)
			}
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
		want    []string
		wantErr string
	}{
		{"lyrics", 200, `{"plainLyrics":"First line\n\nSecond verse"}`, []string{"First line", "", "Second verse"}, ""},
		{"instrumental", 200, `{"instrumental":true}`, []string{"Instrumental"}, ""},
		{"missing lyrics", 200, `{}`, nil, "no lyrics found"},
		{"blank lyrics", 200, `{"plainLyrics":"  "}`, nil, "no lyrics found"},
		{"not found", 404, `{}`, nil, "no lyrics found"},
		{"HTTP failure", 503, `unavailable`, nil, "status 503"},
		{"malformed JSON", 200, `{`, nil, "could not parse"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			body := &trackedBody{Reader: strings.NewReader(tt.body)}
			httpClient = &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if r.Header.Get("User-Agent") != userAgent || r.URL.Query().Get("album_name") != "Album" || r.URL.Query().Get("track_name") != "Song" {
					t.Errorf("unexpected request: %s, headers %v", r.URL, r.Header)
				}
				return &http.Response{StatusCode: tt.status, Body: body, Header: make(http.Header)}, nil
			})}
			got, err := Get(track.New("Artist", "Album (Deluxe Edition)", "Song (Remastered)"))
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
	_, err := Get(track.New("Artist", "Album (Deluxe Edition)", "Song (Remastered)"))
	if !errors.Is(err, failure) {
		t.Fatalf("expected wrapped transport error, got %v", err)
	}
}
