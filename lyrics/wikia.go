package lyrics

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/FrankKair/spootify/track"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
)

// Wikia returns string slice containing lyrics of the song
func wikia(track track.Track) ([]string, error) {
	url := getLyricsURL(track)

	resp, err := http.Get(url)
	if err != nil {
		e := fmt.Sprintf("Could not access the URL: %s", err)
		return []string{}, errors.New(e)
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		e := fmt.Sprintf("Could not parse the HTML body: %s", err)
		return []string{}, errors.New(e)
	}

	node, ok := scrape.Find(root, scrape.ByClass("lyricbox"))
	if ok {
		lyrics := buildLyrics(node.FirstChild)
		return lyrics, nil
	}

	return []string{}, errors.New("Could not fetch song lyrics")
}

func getLyricsURL(track track.Track) string {
	artist := strings.Replace(track.Artist, " ", "_", -1)
	song := strings.Replace(track.Title, " ", "_", -1)
	return fmt.Sprintf("http://lyrics.wikia.com/wiki/%s:%s", artist, song)
}

func buildLyrics(node *html.Node) []string {
	lyrics := []string{}

	for node.Data != "div" {
		if node.Data == "span" || node.Data == "b" {
			if node.FirstChild.Data == "Instrumental" {
				lyrics = append(lyrics, "Instrumental")
				return lyrics
			}
			node = node.NextSibling
			continue
		}

		if node.Data == "br" {
			node = node.NextSibling
			if node.Data == "br" {
				lyrics = append(lyrics, "")
				continue
			}
		}

		lyrics = append(lyrics, node.Data)
		node = node.NextSibling
	}

	return lyrics
}
