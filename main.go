package main

import (
	"fmt"
	"os"
)

func main() {
	track, err := runAppleScript()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	info, err := getLastfmInfo(track)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(info.URL)
	fmt.Println(info.Info)
}
