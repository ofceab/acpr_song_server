package main

import (
	controllers "acpr_songs_server/controllers/song_controller"

	"github.com/gin-gonic/gin"
)

func main() {
	// init song controller
	_controller := controllers.New()
	router := gin.Default()

	//Init Song routes
	// Get songs take latest version of songs
	router.GET("/songs", _controller.FetchSongs)
	// Get songs by release version
	router.GET("/songs/:releaseVersion", _controller.FetchSongsPerVersionId)
	// Add songs on a certain release version
	router.POST("/songs/:releaseVersion", _controller.AddSong)
	// Delete songs on a certain release
	router.DELETE("/songs/:releaseVersion/:songId", _controller.DeleteSong)

	//Init

	router.Run(":8080")
}
