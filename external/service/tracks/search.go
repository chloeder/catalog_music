package tracks

import (
	"catalog-music/external/models/spotify"
	externalSpotify "catalog-music/external/spotify"
	"context"

	"github.com/rs/zerolog/log"
)

func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int) (*spotify.SearchResponse, error) {
	offset := (pageIndex - 1) * pageSize
	limit := pageSize

	response, err := s.spotifyOutbound.Search(ctx, query, offset, limit)
	if err != nil {
		log.Error().Err(err).Msg("failed to search tracks")
		return nil, err
	}

	return modelToResponse(response), nil
}

func modelToResponse(data *externalSpotify.SpotifySearchResponse) *spotify.SearchResponse {
	if data == nil {
		return nil
	}

	items := make([]spotify.SpotifyTrackItem, 0, len(data.Tracks.Items))

	for _, item := range data.Tracks.Items {
		albumImagesURL := make([]string, 0, len(item.Album.Images))
		for _, image := range item.Album.Images {
			albumImagesURL = append(albumImagesURL, image.URL)
		}

		items = append(items, spotify.SpotifyTrackItem{
			ID:               item.ID,
			Name:             item.Name,
			AlbumID:          item.Album.ID,
			AlbumName:        item.Album.Name,
			AlbumImagesURL:   albumImagesURL,
			AlbumTotalTracks: item.Album.TotalTracks,
			ArtistID:         item.Artists[0].ID,
			ArtistName:       item.Artists[0].Name,
			ArtistHref:       item.Artists[0].Href,
			Explicit:         item.Explicit,
		})
	}

	return &spotify.SearchResponse{
		Limit:  data.Tracks.Limit,
		Offset: data.Tracks.Offset,
		Total:  data.Tracks.Total,
		Items:  items,
	}
}
