package spotify

type SearchResponse struct {
	Limit  int                `json:"limit"`
	Offset int                `json:"offset"`
	Total  int                `json:"total"`
	Items  []SpotifyTrackItem `json:"items"`
}

type SpotifyTrackItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	AlbumID          string   `json:"albumId"`
	AlbumName        string   `json:"albumName"`
	AlbumImagesURL   []string `json:"albumImagesURL"`
	AlbumTotalTracks int      `json:"albumTotalTracks"`

	ArtistID   string `json:"artistId"`
	ArtistName string `json:"artistName"`
	ArtistHref string `json:"artistHref"`

	Explicit bool `json:"explicit"`
}
