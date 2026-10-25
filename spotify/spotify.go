package spotify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/FrankKair/spootify/track"
)

// JXA serializes metadata as JSON so punctuation and newlines remain unambiguous.
const tellSpotify = `
const spotify = Application("Spotify");
if (!spotify.running()) throw new Error("Spotify is not running");
const song = spotify.currentTrack();
JSON.stringify({artist: song.artist(), album: song.album(), title: song.name()});
`

// Run fetches the current Spotify track with a bounded execution time.
func Run() (track.Track, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return run(ctx, "osascript")
}

func run(ctx context.Context, executable string) (track.Track, error) {
	cmd := exec.CommandContext(ctx, executable, "-l", "JavaScript", "-e", tellSpotify)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		if ctx.Err() != nil {
			err = ctx.Err()
		}
		return track.Track{}, executionError(err, stderr.String())
	}
	return parse(out)
}

type spotifyError struct {
	message string
	cause   error
}

// Satisfies type error interface { Error() string }
func (e *spotifyError) Error() string { return e.message }
func (e *spotifyError) Unwrap() error { return e.cause }

func executionError(cause error, diagnostic string) error {
	diagnostic = strings.ToLower(diagnostic)
	message := "Could not read the current track. Open Spotify and play a song."
	switch {
	case errors.Is(cause, context.DeadlineExceeded):
		message = "Spotify took too long to respond. Try again."
	case errors.Is(cause, context.Canceled):
		message = "Reading Spotify was canceled."
	case errors.Is(cause, os.ErrNotExist), errors.Is(cause, exec.ErrNotFound):
		message = "The macOS scripting tool (osascript) is unavailable."
	case strings.Contains(diagnostic, "application can't be found"):
		message = "Spotify could not be found. Install the Spotify desktop app and try again."
	case strings.Contains(diagnostic, "spotify is not running"):
		message = "Spotify is not running. Open Spotify and play a song."
	case strings.Contains(diagnostic, "(-1743)"), strings.Contains(diagnostic, "not authorized"), strings.Contains(diagnostic, "permission denied"):
		message = "Allow your terminal to control Spotify in System Settings > Privacy & Security > Automation."
	}
	return &spotifyError{message: message, cause: cause}
}

func parse(out []byte) (track.Track, error) {
	var result track.Track
	if err := json.Unmarshal(out, &result); err != nil {
		return track.Track{}, &spotifyError{message: "Spotify returned unreadable track information. Try playing another song.", cause: err}
	}
	if strings.TrimSpace(result.Artist) == "" || strings.TrimSpace(result.Title) == "" {
		return track.Track{}, fmt.Errorf("No music track is selected. Play a song in Spotify and try again.")
	}
	return result, nil
}
