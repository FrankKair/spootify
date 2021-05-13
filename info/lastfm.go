package info

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/FrankKair/spootify/track"
	"github.com/yhat/scrape"

	"golang.org/x/net/html"
)

// Get returns information from Last.fm
func lastfm(track track.Track) (string, error) {
	url := getLastfmURL(track)

	resp, err := http.Get(url)
	if err != nil {
		e := fmt.Sprintf("Could not access the URL: %s", err)
		return "", errors.New(e)
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		e := fmt.Sprintf("Could not parse the HTML body: %s", err)
		return "", errors.New(e)
	}

	node, ok := scrape.Find(root, scrape.ByClass("wiki-content"))
	if ok {
		return scrape.Text(node), nil
	}

	// We don't have a wiki here yet...
	node, ok = scrape.Find(root, scrape.ByClass("wiki"))
	if ok {
		return scrape.Text(node), nil
	}

	return "", errors.New("Could not fetch album info")
}

func getLastfmURL(track track.Track) string {
	artist := strings.Replace(track.Artist, " ", "+", -1)
	album := strings.Replace(track.Album, " ", "+", -1)
	if strings.Contains(album, "/") {
		album = strings.Replace(album, "/", "%2F", -1)
	}
	return fmt.Sprintf("https://www.last.fm/music/%s/%s/+wiki", artist, album)
}
