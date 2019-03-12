package main

import (
	"errors"
	"fmt"

	"github.com/everdev/mack"
)

// Runs AppleScript "tell" and returns a Track{artist, album}
func runAppleScript() (Track, error) {
	err := isRunning()
	if err != nil {
		return Track{}, err
	}

	artist, err := artist()
	if err != nil {
		return Track{}, err
	}

	album, err := album()
	if err != nil {
		return Track{}, err
	}

	track := newTrack(artist, album)
	return track, nil
}

func isRunning() error {
	const isRunningCmd = `
	if it is running then
		return true
	else
		return false
	end if
	`

	isRunning, err := mack.Tell("Spotify", isRunningCmd)
	if err != nil {
		return errors.New("Could not exectute AppleScript properly")
	}

	if isRunning == "false" {
		return errors.New("Spotify is not running, please open the app")
	}

	return nil
}

func artist() (string, error) {
	const artistCmd = `
	set ctrack to ""
	set ctrack to ctrack & (current track's artist)
	`

	artist, err := mack.Tell("Spotify", artistCmd)
	if err != nil {
		e := fmt.Sprintf("Could not exectute AppleScript properly: %s", err)
		return "", errors.New(e)
	}

	return artist, nil
}

func album() (string, error) {
	const albumCmd = `
	set ctrack to ""
	set ctrack to ctrack & (current track's album)
	`

	album, err := mack.Tell("Spotify", albumCmd)
	if err != nil {
		e := fmt.Sprintf("Could not exectute AppleScript properly: %s", err)
		return "", errors.New(e)
	}

	return album, nil
}
