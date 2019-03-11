package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// AlbumInfo type
type AlbumInfo struct {
	URL  string
	Info string
}

// GetAlbumInfo returns album info and error
func GetAlbumInfo() AlbumInfo {
	// Output shape is: Artist-Album
	output, err := exec.Command("osascript", "track.applescript").Output()
	if err != nil {
		fmt.Println("Could not exectute AppleScript properly:", err)
		os.Exit(1)
	}

	if strings.TrimSpace(string(output)) == "Spotify is not running" {
		fmt.Println("Spotify is not running")
		os.Exit(1)
	}

	track := newTrack(string(output))
	url := getLastfmURL(track)
	info, err := getLastfmInfo(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return AlbumInfo{url, info}
}
