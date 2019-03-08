package main

import (
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/yhat/scrape"

	"golang.org/x/net/html"
)

func cleanAlbumTitle(title string) string {
	title = strings.TrimSpace(title)
	extras := []string{
		"(Remastered)",
		"(Remastered+Version)",
		"(Deluxe)",
		"(Deluxe+Edition)",
	}

	for i := range extras {
		if strings.Contains(title, extras[i]) {
			sliceIndex := strings.Index(title, "(") - 1
			return title[:sliceIndex]
		}
	}

	return title
}

// GetAlbumInfo returns album info and error
func GetAlbumInfo() (string, error) {
	// Output shape is: Artist-Album
	output, err := exec.Command("osascript", "track.applescript").Output()
	if err != nil {
		fmt.Println(err)
	}

	info := strings.Split(string(output), "-")
	artist := strings.Replace(info[0], " ", "+", -1)
	album := strings.Replace(info[1], " ", "+", -1)
	albumTitle := cleanAlbumTitle(album)

	url := fmt.Sprintf("https://www.last.fm/music/%s/%s/+wiki", artist, albumTitle)
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
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
