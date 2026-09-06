package lyrics

import "github.com/FrankKair/spootify/track"

// Get returns the track's lyrics from lrclib.net.
// No API key required.
func Get(track track.Track) ([]string, error) {
  return lrclib(track)
}
