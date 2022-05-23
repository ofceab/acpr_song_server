package dal_interfaces

import (
	"acpr_songs_server/core/errors"
	dataformat "acpr_songs_server/data_format"
	"acpr_songs_server/models"

	"gorm.io/gorm"
)

// Define interfaction for getting song or deal with the song database
type ISongDatabaseAccessLayer interface {
	// Save song in store
	SaveSong(s *dataformat.CreateSong, releaseVersion uint) (models.Song, errors.SongError)
	// Update song
	UpdateSong(s *dataformat.UpdateSong, releaseVersion uint) (models.Song, errors.SongError)
	// Fetch songs
	FetchSongs() ([]models.Song, errors.SongError)
	// Fetch all sounds per version id for fetching release song of a certain `version Id`
	FetchSongsPerVersionId(releaseVersion uint) ([]models.Song, errors.SongError)
	// Fetch song by SongUnqueId
	FetchSongsPerSongUniqueId(snUID string) ([]models.Song, errors.SongError)
	// Remove song
	DeleteSong(songId uint) (models.Song, errors.SongError)
}

// Define method in use for access song database
type IReleaseVersionDatabaseAccessLayer interface {
	// Create a ReleaseVersion
	CreateReleaseVersion() (models.ReleaseVersion, errors.ReleaseVersionError)
	// Fetch all release version
	FetchReleaseVersions() ([]models.ReleaseVersion, errors.ReleaseVersionError)
	// Delete a particular release version
	DeleteReleaseVersion(releaseVersion uint) (models.ReleaseVersion, errors.ReleaseVersionError)
	// Fetch release version per Id
	FetchReleaseVersionById(id uint) (models.ReleaseVersion, errors.ReleaseVersionError)
	// Fetch latest Release Version
	FetchLatestReleaseVersion() (models.ReleaseVersion, errors.ReleaseVersionError)
}

func NewIReleaseVersionDatabaseAccessLayer(db *gorm.DB) IReleaseVersionDatabaseAccessLayer {
	return &MysqlReleaseVersionDataAccessLayer{DbConnection: db}
}

func NewISongDatabaseAccessLayer(db *gorm.DB, rDal IReleaseVersionDatabaseAccessLayer) ISongDatabaseAccessLayer {
	return &MysqlSongDataAccessLayer{DbConnection: db, ReleaseVersionDal: rDal}
}
