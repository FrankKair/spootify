package track

import "testing"

func TestClean(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"parenthesized remastered", "Led Zeppelin II (Remastered)", "Led Zeppelin II"},
		{"dash remaster", "What Is and What Should Never Be - 2014 Remaster", "What Is and What Should Never Be"},
		{"deluxe version", "Random Acces Memories (Deluxe Version)", "Random Acces Memories"},
		{"deluxe edition", "Abbey Road (Deluxe Edition)", "Abbey Road"},
		{"bracket remastered", "OK Computer [Remastered]", "OK Computer"},
		{"single", "Shape of You (Single)", "Shape of You"},
		{"bonus track", "Hidden Track - Bonus Track", "Hidden Track"},
		{"anniversary", "Thriller (25th Anniversary Edition)", "Thriller"},
		{"expanded edition", "Rumours (Expanded Edition)", "Rumours"},
		{"special edition", "The Dark Side of the Moon (Special Edition)", "The Dark Side of the Moon"},
		{"super deluxe", "Abbey Road (Super Deluxe)", "Abbey Road"},
		{"no suffix", "OK Computer", "OK Computer"},
		{"dash no keyword", "Hide and Seek - Live", "Hide and Seek - Live"},
		{"parenthesized no keyword", "Song (feat. Someone)", "Song (feat. Someone)"},
		{"empty string", "", ""},
		{"case insensitive", "Album (REMASTERED)", "Album"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clean(tt.input)
			if got != tt.want {
				t.Errorf("clean(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tr := New("Led Zeppelin", "Led Zeppelin II (Remastered)", "Whole Lotta Love - 2014 Remaster")
	if tr.Artist != "Led Zeppelin" {
		t.Errorf("Artist = %q, want %q", tr.Artist, "Led Zeppelin")
	}
	if tr.Album!= "Led Zeppelin II" {
		t.Errorf("Album = %q, want %q", tr.Album, "Led Zeppelin II")
	}
	if tr.Title != "Whole Lotta Love" {
		t.Errorf("Title = %q, want %q", tr.Title, "Whole Lotta Love")
	}
}
