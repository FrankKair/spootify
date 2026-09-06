package lyrics

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "strings"
    "time"
    
    "github.com/FrankKair/spootify/track"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

type lrclibResponse struct {
    PlainLyrics  string `json:"plainLyrics"`
    Instrumental bool   `json:"instrumental"`
    StatusCode	 int	`json:"statusCode"`
}

const userAgent = "spootify"

func lrclib(track track.Track) ([]string, error) {
    req, err := http.NewRequest("GET", getLrclibURL(track), nil)
    if err != nil {
    	return nil, fmt.Errorf("could not build lrclib request: %w", err)
    }
    req.Header.Set("User-Agent", userAgent)

    resp, err := httpClient.Do(req)
    if err != nil {
	return nil, fmt.Errorf("could not reach lrclib.net: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusNotFound {
    	return nil, fmt.Errorf("no lyrics found for %s - %s", track.Artist, track.Title)
    }

    if resp.StatusCode != http.StatusOK {
	return nil, fmt.Errorf("lrclib.net returned status %d", resp.StatusCode)
    }
 
    var result lrclibResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
    	return nil, fmt.Errorf("could not parse lrclib response: %w", err)
    }

    if result.Instrumental {
	return []string{"Instrumental"}, nil
    }

    if strings.TrimSpace(result.PlainLyrics) == "" {
    	return nil, fmt.Errorf("no lyrics found for %s - %s", track.Artist, track.Title)
    }

    return strings.Split(result.PlainLyrics, "\n"), nil
}

func getLrclibURL(track track.Track) string {
    v := url.Values{}
    v.Set("artist_name", track.Artist)
    v.Set("track_name", track.Title)
    v.Set("album_name", track.Album)
    return "https://lrclib.net/api/get?" + v.Encode()
}
