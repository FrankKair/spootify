package main

import (
	"fmt"
	"os"

	"github.com/FrankKair/spootify/applescript"
	"github.com/FrankKair/spootify/info"
	"github.com/FrankKair/spootify/lyrics"
)

func main() {
	track, err := applescript.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s - %s - %s \n\n", track.Artist, track.Album, track.Title)
	info, err := info.Get(track)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("* ALBUM INFORMATION:")
	fmt.Printf("%s\n\n", info)

	lyrics, err := lyrics.Get(track)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("* LYRICS:")
	for verse := range lyrics {
		fmt.Println(lyrics[verse])
	}
}
