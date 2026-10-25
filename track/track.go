package track

import (
	"regexp"
	"strings"
)

// Track holds the currently playing song's metadata.
type Track struct {
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Title  string `json:"title"`
}

// New creates a Track with common edition suffixes removed for API lookups.
func New(artist, album, title string) Track {
	return Track{Artist: artist, Album: clean(album), Title: clean(title)}
}

// Only strip complete edition labels at the end, preserving meaningful subtitles.
var editionSuffix = regexp.MustCompile(`(?i)^(?:[0-9]{4} )?(?:remaster(?:ed)?(?: [0-9]{4})?|(?:super )?deluxe(?: edition| version)?|single(?: version| edit)?|bonus track|expanded(?: edition)?|(?:[0-9]+(?:st|nd|rd|th)? )?anniversary(?: edition)?|special edition)$`)

func clean(title string) string {
	for {
		start := -1
		label := ""
		for _, delimiter := range []string{" - ", " (", " ["} {
			idx := strings.LastIndex(title, delimiter)
			if idx <= start {
				continue
			}
			candidate := title[idx+len(delimiter):]
			if delimiter == " (" || delimiter == " [" {
				closer := ")"
				if delimiter == " [" {
					closer = "]"
				}
				if !strings.HasSuffix(candidate, closer) {
					continue
				}
				candidate = strings.TrimSuffix(candidate, closer)
			}
			start, label = idx, candidate
		}
		if start <= 0 || !editionSuffix.MatchString(strings.TrimSpace(label)) {
			return title
		}
		title = strings.TrimSpace(title[:start])
	}
}
