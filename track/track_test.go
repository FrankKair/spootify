package track

import "testing"

func TestCleanAlbumTitle(t *testing.T) {
	album := clean("Led Zeppelin II (Remastered)")
	if album != "Led Zeppelin II" {
		t.Error("clean is not working properly")
	}
}

func TestCleanSongTitle(t *testing.T) {
	song := clean("What Is and What Should Never Be - 2014 Remaster")
	if song != "What Is and What Should Never Be" {
		t.Error("clean is not working properly")
	}
}
