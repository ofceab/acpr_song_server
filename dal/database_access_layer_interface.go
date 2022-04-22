package dal

import (
	"acpr_songs_server/dal/mysql/release_version"
	"acpr_songs_server/models"

	"gorm.io/gorm"
)

// Define interfaction for getting song or deal with the song database
type ISongDatabaseAccessLayer interface {
	// Save song in store
	SaveSong(s *models.Song, releaseVersion uint) (models.Song, error)
	// Fetch songs
	FetchSongs() []models.Song
	// Fetch all sounds per version id for fetching release song of a certain `version Id`
	FetchSongsPerVersionId(releaseVersion uint) ([]models.Song, error)
	// Remove song from a certain release
	DeleteSongByReleaseVersion(releaseVersion uint, song models.Song) (models.Song, error)
	// Remove song no matter ReleasVersion
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
	// Fetch latest Release Version
	FetchLatestReleaseVersion() (models.ReleaseVersion, error)
}

func NewIReleaseVersionDatabaseAccessLayer(db *gorm.DB) IReleaseVersionDatabaseAccessLayer {
	return &release_version.MysqlReleaseVersionDataAccessLayer{DbConnection: db}
}
