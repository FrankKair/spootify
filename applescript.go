package main

import (
	"errors"
	"fmt"

	"github.com/everdev/mack"
)

const isRunningCmd = `
if it is running then 
	return true 
else 
	return false 
end if
`
const artistAndAlbumCmd = `
set ctrack to ""
set ctrack to ctrack & (current track's artist) & "-"
set ctrack to ctrack & (current track's album)
`

// Runs AppleScript "tell" and returns string shaped Artist-Album
func runAppleScript() (string, error) {
	isRunning, err := mack.Tell("Spotify", isRunningCmd)
	if err != nil {
		return "", errors.New("Could not exectute AppleScript properly")
	}
	if isRunning == "false" {
		return "", errors.New("Spotify is not running, please open the app")
	}

	output, err := mack.Tell("Spotify", artistAndAlbumCmd)
	if err != nil {
		e := fmt.Sprintf("Could not exectute AppleScript properly: %s", err)
		return "", errors.New(e)
	}

	return output, nil
}
