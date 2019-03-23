package track

import "strings"

// Track type
type Track struct {
	Artist string
	Album  string
}

// New creates a Track
func New(artist, album string) Track {
	albumTitle := cleanAlbumTitle(album)
	return Track{artist, albumTitle}
}

func cleanAlbumTitle(title string) string {
	title = strings.TrimSpace(title)
	extras := []string{
		"(Remastered)",
		"(Remastered Version)",
		"(Deluxe)",
		"(Deluxe Edition)",
	}

	for i := range extras {
		if strings.Contains(title, extras[i]) {
			sliceIndex := strings.Index(title, "(") - 1
			return title[:sliceIndex]
		}
	}

	return title
}
