package info

import "github.com/FrankKair/spootify/track"

// Get returns album information from Last.fm.
// Requires a valid Last.fm API key.
func Get(track track.Track, apiKey string) (string, error) {
  return lastfm(track, apiKey)
}
