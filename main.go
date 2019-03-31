package main

import (
	"fmt"
	"os"

	"github.com/FrankKair/spootify/applescript"
	"github.com/FrankKair/spootify/lastfm"
	"github.com/FrankKair/spootify/lyrics"
)

func main() {
	track, err := applescript.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s - %s - %s \n\n", track.Artist, track.Album, track.Title)
	printAlbumInfo(lastfm.GetInfo(track))
	printLyrics(lyrics.GetLyrics(track))
}

func printAlbumInfo(info string, err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(info)
	fmt.Println("")
}

func printLyrics(lyrics []string, err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for verse := range lyrics {
		fmt.Println(lyrics[verse])
	}
	fmt.Println("")
}
