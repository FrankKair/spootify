package main

import "fmt"

func main() {
	albumInfo := GetAlbumInfo()
	fmt.Println(albumInfo.URL)
	fmt.Println(albumInfo.Info)
}
