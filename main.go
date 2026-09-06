package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/FrankKair/spootify/applescript"
	"github.com/FrankKair/spootify/info"
	"github.com/FrankKair/spootify/lyrics"
)

var version = "dev"

func main() {
	showVersion := flag.Bool("version", false, "print version and exit")
	lyricsOnly := flag.Bool("lyrics", false, "show only lyrics")
	infoOnly := flag.Bool("info", false, "show only album info")
	flag.Parse()

	if *showVersion {
		fmt.Println("spootify", version)
		return
	}

	track, err := applescript.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%s - %s - %s\n\n", track.Artist, track.Album, track.Title)

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
			albumInfo, albumInfoErr = info.Get(track, apiKey)
		})
	}

	if showLyrics {
		wg.Go(func() {
			lyricsLines, lyricsErr = lyrics.Get(track)
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
