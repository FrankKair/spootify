package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/FrankKair/spootify/info"
	"github.com/FrankKair/spootify/lyrics"
	"github.com/FrankKair/spootify/spotify"
	"github.com/FrankKair/spootify/track"
)

func main() {
	lyricsOnly := flag.Bool("lyrics", false, "show only lyrics")
	infoOnly := flag.Bool("info", false, "show only album info")
	flag.Parse()

	song, err := spotify.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%s - %s - %s\n\n", song.Artist, song.Album, song.Title)

	lookupTrack := track.New(song.Artist, song.Album, song.Title)

	apiKey := os.Getenv("LASTFM_API_KEY")
	showInfo := !*lyricsOnly
	showLyrics := !*infoOnly

	if showInfo && apiKey == "" {
		fmt.Fprintln(os.Stderr, "warning: LASTFM_API_KEY not set, skipping album info (get one at https://www.last.fm/api/account/create)")
		showInfo = false
	}

	var (
		albumInfo    string
		albumInfoErr error
		lyricsLines  []string
		lyricsErr    error
	)

	// Fetch album info and lyrics in parallel
	var wg sync.WaitGroup

	if showInfo {
		wg.Go(func() {
			albumInfo, albumInfoErr = info.Get(lookupTrack, apiKey)
		})
	}

	if showLyrics {
		wg.Go(func() {
			lyricsLines, lyricsErr = lyrics.Get(lookupTrack)
		})
	}

	wg.Wait()

	if showInfo {
		if albumInfoErr != nil {
			fmt.Fprintf(os.Stderr, "warning: %v\n\n", albumInfoErr)
		} else {
			fmt.Println("* ALBUM INFORMATION:")
			fmt.Printf("%s\n\n", albumInfo)
		}
	}

	if showLyrics {
		if lyricsErr != nil {
			fmt.Fprintf(os.Stderr, "warning: %v\n\n", lyricsErr)
		} else {
			fmt.Println("* LYRICS:")
			for _, verse := range lyricsLines {
				fmt.Println(verse)
			}
		}
	}
}
