package dal_interfaces

import (
	dataformat "acpr_songs_server/data_format"
	"acpr_songs_server/models"

	"gorm.io/gorm"
)

// Define interfaction for getting song or deal with the song database
type ISongDatabaseAccessLayer interface {
	// Save song in store
	SaveSong(s *dataformat.CreateSong, releaseVersion uint) (models.Song, error)
	// Update song
	UpdateSong(s *dataformat.UpdateSong, releaseVersion uint) (models.Song, error)
	// Fetch songs
	FetchSongs() ([]models.Song, error)
	// Fetch all sounds per version id for fetching release song of a certain `version Id`
	FetchSongsPerVersionId(releaseVersion uint) ([]models.Song, error)
	// Fetch song by SongUnqueId
	FetchSongsPerSongUniqueId(snUID string) ([]models.Song, error)
	// Remove song
	DeleteSong(songId uint) (models.Song, error)
}

// Define method in use for access song database
type IReleaseVersionDatabaseAccessLayer interface {
	// Create a ReleaseVersion
	CreateReleaseVersion() (models.ReleaseVersion, error)
	// Fetch all release version
	FetchReleaseVersions() ([]models.ReleaseVersion, error)
	// Delete a particular release version
	DeleteReleaseVersion(releaseVersion uint) (models.ReleaseVersion, error)
	// Fetch release version per Id
	FetchReleaseVersionById(id uint) (models.ReleaseVersion, error)
	// Fetch latest Release Version
	FetchLatestReleaseVersion() (models.ReleaseVersion, error)
}

func NewIReleaseVersionDatabaseAccessLayer(db *gorm.DB) IReleaseVersionDatabaseAccessLayer {
	return &MysqlReleaseVersionDataAccessLayer{DbConnection: db}
}

func NewISongDatabaseAccessLayer(db *gorm.DB, rDal IReleaseVersionDatabaseAccessLayer) ISongDatabaseAccessLayer {
	return &MysqlSongDataAccessLayer{DbConnection: db, ReleaseVersionDal: rDal}
}
