package applescript

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/FrankKair/spootify/track"
)

const tellSpotify = `tell application "Spotify"
set info to (current track's artist) & "|||" & (current track's album) & "|||" & (current track's name)
end tell`

// Run executes an AppleScript command to get the current Spotify track.
func Run() (track.Track, error) {
	out, err := exec.Command("osascript", "-e", tellSpotify).CombinedOutput()
	if err != nil {
		return track.Track{}, fmt.Errorf(
			"could not fetch track from Spotify -- is it running and playing music? (%s)",
			strings.TrimSpace(string(out)),
		)
	}

	result := strings.TrimSpace(string(out))
	parts := strings.SplitN(result, "|||", 3)
	if len(parts) != 3 {
		return track.Track{}, fmt.Errorf("unexpected AppleScript output: %s", result)
	}

	return track.New(parts[0], parts[1], parts[2]), nil
}
