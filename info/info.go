package info

import "github.com/FrankKair/spootify/track"

// Get returns album information
func Get(track track.Track) (string, error) {
  // TODO: iterate over multiple info backends
  return lastfm(track)
}
