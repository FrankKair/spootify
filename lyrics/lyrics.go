package lyrics

import "github.com/FrankKair/spootify/track"

// Get returns the track's lyrics
func Get(track track.Track) ([]string, error) {
  // TODO: iterate over multiple lyrics backends
  // Check this out -> https://github.com/ddddxxx/LyricsKit
  return wikia(track)
}
