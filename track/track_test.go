package track

import "testing"

func TestCleanAlbumTitle(t *testing.T) {
	album := cleanAlbumTitle("Led Zeppelin II (Remastered)")
	if album != "Led Zeppelin II" {
		t.Error("cleanAlbumTitle is not working properly")
	}
}
