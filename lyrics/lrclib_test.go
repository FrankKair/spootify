package lyrics

import (
    "net/url"
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
	    if got := u.Query().Get("artist_name"); got != tt.wantArtist{
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
