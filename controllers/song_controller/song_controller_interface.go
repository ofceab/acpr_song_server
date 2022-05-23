package song_controller

import (
	"acpr_songs_server/dal/dal_interfaces"
	"acpr_songs_server/service/release_version_service"
	"acpr_songs_server/service/song_service"

	"github.com/gin-gonic/gin"
)

// Controller getting songs
type ISongController interface {
	// Fetch songs all sounds
	FetchSongs(c *gin.Context)
	// Fetch all sounds per version id for fetching release song of a certain `version Id`
	FetchSongsPerVersionId(c *gin.Context)
	FetchSongsPerSongUniqueId(c gin.Context)

	// Add song
	AddSong(c *gin.Context)
	UpdateSong(c *gin.Context)
	// Delete song
	DeleteSong(c *gin.Context)
}

// Create new instance of `ISongController`
func New(songDal dal_interfaces.ISongDatabaseAccessLayer, releaseVersionService release_version_service.IReleaseVersionService) ISongController {
	return &songControllerImpl{
		songService: song_service.New(songDal, releaseVersionService)}
}
