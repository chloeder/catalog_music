package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SpotifySearchResponse struct {
	Tracks SpotifyTrack `json:"tracks"`
}

type SpotifyTrack struct {
	Href     string             `json:"href"`
	Limit    int                `json:"limit"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Offset   int                `json:"offset"`
	Total    int                `json:"total"`
	Items    []SpotifyTrackItem `json:"items"`
}

type SpotifyTrackItem struct {
	ID       string               `json:"id"`
	Name     string               `json:"name"`
	Album    SpotifyTrackAlbum    `json:"album"`
	Artists  []SpotifyTrackArtist `json:"artists"`
	Explicit bool                 `json:"explicit"`
}

type SpotifyTrackAlbum struct {
	ID          string                   `json:"id"`
	Name        string                   `json:"name"`
	Images      []SpotifyTrackAlbumImage `json:"images"`
	TotalTracks int                      `json:"total_tracks"`
}

type SpotifyTrackAlbumImage struct {
	URL string `json:"url"`
}

type SpotifyTrackArtist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Href string `json:"href"`
}

func (o *outbound) Search(ctx context.Context, query string, offset, limit int) (*SpotifySearchResponse, error) {
	token, _, err := o.GetTokenDetails()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track&offset=%d&limit=%d", query, offset, limit)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	bearerToken := fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Authorization", bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response SpotifySearchResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
