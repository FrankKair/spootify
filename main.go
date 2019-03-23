package main

import (
	"fmt"
	"os"

	"github.com/FrankKair/spootify/applescript"
	"github.com/FrankKair/spootify/lastfm"
)

func main() {
	track, err := applescript.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	info, err := lastfm.GetInfo(track)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(info.URL)
	fmt.Println(info.Info)
}
