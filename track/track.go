package track

import (
	"strings"
)

// Track type
type Track struct {
	Artist string
	Album  string
	Title  string
}

// New creates a Track
func New(artist, album, title string) Track {
	albumTitle := clean(album)
	songTitle := clean(title)
	return Track{artist, albumTitle, songTitle}
}

func clean(title string) string {
	delimiter := []string{
		"-",
		"(",
	}

	extras := []string{
		"Remaster",
		"Remastered",
		"Single",
		"(Remastered)",
		"(Remastered Version)",
		"(Deluxe)",
		"(Deluxe Version)",
		"(Deluxe Edition)",
	}

	for i := range delimiter {
		if strings.Contains(title, delimiter[i]) {
			index := strings.Index(title, delimiter[i])
			subString := title[index:]
			for j := range extras {
				if strings.Contains(subString, extras[j]) {
					return strings.TrimSpace(title[:index])
				}
			}
		}
	}

	return title
}
