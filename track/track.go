package track

import (
	"strings"
)

// Track holds the currently playing song's metadata.
type Track struct {
	Artist string
	Album  string
	Title  string
}

// New creates a Track, stripping common Spotify suffixes like
// "Remastered", "Deluxe Edition", etc. from album and song titles.
func New(artist, album, title string) Track {
	return Track{
		Artist: artist,
		Album:  clean(album),
		Title:  clean(title),
	}
}

// delimiters marks where a suffix region might begin.
var delimiters = []string{" - ", " (", " ["}

// suffixKeywords are matched case-insensitively inside the suffix region.
var suffixKeywords = []string{
	"remaster",
	"deluxe",
	"single",
	"bonus track",
	"expanded",
	"anniversary",
	"special edition",
	"super deluxe",
}

func clean(title string) string {
	lower := strings.ToLower(title)
	for _, d := range delimiters {
		idx := strings.Index(lower, d)
		if idx < 0 {
			continue
		}
		suffix := lower[idx:]
		for _, kw := range suffixKeywords {
			if strings.Contains(suffix, kw) {
				return strings.TrimSpace(title[:idx])
			}
		}
	}
	return title
}
