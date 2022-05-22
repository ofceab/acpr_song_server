package song_service

import (
	"acpr_songs_server/dal/dal_interfaces"
	dataformat "acpr_songs_server/data_format"
	"acpr_songs_server/models"
	"acpr_songs_server/service/release_version_service"
)

// Service for providing songs
type ISongService interface {
	// Fetch songs all sounds
	FetchSongs() ([]models.Song, error)
	// Fetch all sounds per version id for fetching release song of a certain `version Id`
	FetchSongsPerVersionId(sv uint) ([]models.Song, error)
	// Fetch by SongUnique Id
	FetchSongsPerSongUniqueId(snUID string) ([]models.Song, error)
	// Add song with
	AddSong(s *dataformat.CreateSong, relaseVersion uint) (models.Song, error)
	// Update song
	UpdateSong(s *dataformat.UpdateSong, releaseVersion uint) (models.Song, error)
	// Remove song from a certain release
	DeleteSong(sId uint) (models.Song, error)
}

func New(songDal dal_interfaces.ISongDatabaseAccessLayer, releaseVersionService release_version_service.IReleaseVersionService) ISongService {

	return &songServiceImpl{
		songDal:               songDal,
		releaseVersionService: releaseVersionService,
	}
}
