package main

import "strings"

// Track type
type Track struct {
	Artist string
	Album  string
}

func newTrack(appleScriptInfo string) Track {
	info := strings.Split(string(appleScriptInfo), "-")
	artist := info[0]
	album := info[1]
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
