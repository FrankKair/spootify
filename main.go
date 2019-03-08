package main

import "fmt"

func main() {
	info, err := GetAlbumInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)
}
