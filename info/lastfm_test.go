package info

import (
	"testing"

	"github.com/FrankKair/spootify/track"
)

func TestGetLastfmURL(t *testing.T) {
	track := track.New("Justin Timberlake", "The 20/20 Experience", "Pusher Love Girl")
	url := getLastfmURL(track)
	if url != "https://www.last.fm/music/Justin+Timberlake/The+20%2F20+Experience/+wiki" {
		t.Error("getLastfmURL is not working properly")
	}
}
