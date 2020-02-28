package applescript

import (
	"errors"
	"fmt"

	"github.com/FrankKair/spootify/track"
	"github.com/everdev/mack"
)

// Run AppleScript "tell" and returns a Track{artist, album}
func Run() (track.Track, error) {
	err := isRunning()
	if err != nil {
		return track.Track{}, err
	}

	artist, err := artist()
	if err != nil {
		return track.Track{}, err
	}

	album, err := album()
	if err != nil {
		return track.Track{}, err
	}

	song, err := song()
	if err != nil {
		return track.Track{}, err
	}

	track := track.New(artist, album, song)
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
		return errors.New("Spotify is either not installed or running")
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
		const errMsg =`
    		Could not fetch the artist/song from Spotify.
    		Are you logged in? Is there any music playing on the app?
    		`
		return "", errors.New(errMsg)
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

func song() (string, error) {
	const songCmd = `
	set ctrack to ""
	set ctrack to ctrack & (current track's name)
	`

	song, err := mack.Tell("Spotify", songCmd)
	if err != nil {
		e := fmt.Sprintf("Could not exectute AppleScript properly: %s", err)
		return "", errors.New(e)
	}

	return song, nil
}
