package song_controller

import (
	"acpr_songs_server/service/song_service"

	"github.com/gin-gonic/gin"
)

// Controller getting songs
type ISongController interface {
	// Fetch songs all sounds
	FetchSongs(c *gin.Context)
	// Fetch all sounds per version id for fetching release song of a certain `version Id`
	FetchSongsPerVersionId(c *gin.Context)
	// Add song
	AddSong(c *gin.Context)
	// Delete song
	DeleteSong(c *gin.Context)
}

// Create new instance of `ISongController`
func New() ISongController {
	return &songControllerImpl{
		songService: song_service.New()}
}
