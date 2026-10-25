package spotify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
		return track.Track{}, fmt.Errorf("could not fetch track from Spotify (is it running and playing music?): %w (%s)", err, strings.TrimSpace(stderr.String()))
	}
	return parse(out)
}

func parse(out []byte) (track.Track, error) {
	var result track.Track
	if err := json.Unmarshal(out, &result); err != nil {
		return track.Track{}, fmt.Errorf("could not parse Spotify metadata: %w", err)
	}
	if strings.TrimSpace(result.Artist) == "" || strings.TrimSpace(result.Title) == "" {
		return track.Track{}, fmt.Errorf("Spotify returned no artist or title; is a music track selected?")
	}
	return result, nil
}
