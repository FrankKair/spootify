package info

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/FrankKair/spootify/track"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

var htmlTagRe = regexp.MustCompile(`<[^>]*>`)

type lastFmResponse struct {
	Album struct {
		Wiki struct {
			Summary string `json:"summary"`
			Content	string `json:"content"`
		} `json:"wiki"`
	} `json:"album"`
	Error	int	   `json:"error"`
	Message string `json:"message"`
}

const userAgent = "spootify"

func lastfm(track track.Track, apiKey string) (string, error) {
	req, err := http.NewRequest("GET", getLastfmURL(track, apiKey), nil)
	if err != nil {
		return "", fmt.Errorf("could not build Last.fm request: %w", err)
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not reach Last.fm API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Last.API returned status %d", resp.StatusCode)
	}

	var result lastFmResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("could not parse Last.fm response: %w", err)
	}

	if result.Error != 0 {
		return "", fmt.Errorf("Last.fm API error: %s", result.Message)
	}

	summary := strings.TrimSpace(htmlTagRe.ReplaceAllString(result.Album.Wiki.Summary, ""))
	if summary == "" {
		return "", fmt.Errorf("no album info found for %s - %s", track.Artist, track.Album)
	}

	return summary, nil
}

func getLastfmURL(track track.Track, apiKey string) string {
	v := url.Values{}
	v.Set("method", "album.getinfo")
	v.Set("api_key", apiKey)
	v.Set("artist", track.Artist)
	v.Set("album", track.Album)
	v.Set("format", "json")
	return "https://ws.audioscrobbler.com/2.0/?" + v.Encode()
}
