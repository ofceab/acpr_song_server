package song_service

import (
	"acpr_songs_server/core/constants"
	"acpr_songs_server/models"
)

// Service for providing songs
type ISongService interface {
	// Fetch songs all sounds
	FetchSongs() []models.Song
	// Fetch all sounds per version id for fetching release song of a certain `version Id`
	FetchSongsPerVersionId() []models.Song

	// Add song with
	AddSong(s *models.Song, relaseVersion string) models.Song
	// Remove song from a certain release
	DeleteSong(s string) (models.Song, error)
}

func New() ISongService {
	customApiBaseUrl := constants.API_BASE_URL
	db := []models.Song{}

	return &songServiceImpl{
		apiBaseUrl: customApiBaseUrl,
		lastIndex:  1,
		localDB:    db,
	}
}
