package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/yhat/scrape"

	"golang.org/x/net/html"
)

func getLastfmURL(track Track) string {
	artist := strings.Replace(track.Artist, " ", "+", -1)
	album := strings.Replace(track.Album, " ", "+", -1)
	return fmt.Sprintf("https://www.last.fm/music/%s/%s/+wiki", artist, album)
}

func getLastfmInfo(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Could not access the URL", err)
		os.Exit(1)
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Could not parse the HTML body", err)
		os.Exit(1)
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
