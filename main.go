package main

import (
	"acpr_songs_server/controllers/release_version_controller"
	"acpr_songs_server/controllers/song_controller"

	"github.com/gin-gonic/gin"
)

func main() {
	// init song controller
	_songController := song_controller.New()
	_releaseVersionController := release_version_controller.New()
	router := gin.Default()

	//Init Song routes
	// Get songs take latest version of songs

	router.GET("/songs", _songController.FetchSongs)
	// Get songs by release version
	router.GET("/songs/:releaseVersion", _songController.FetchSongsPerVersionId)
	// Add songs on a certain release version
	router.POST("/songs/:releaseVersion", _songController.AddSong)
	// Delete songs on a certain release
	router.DELETE("/songs/:releaseVersion/:songId", _songController.DeleteSong)

	//Init ReleaseVersion routes
	router.GET("/releaseVersion", _releaseVersionController.GetReleaseVersions)
	router.GET("/releaseVersion/latest", _releaseVersionController.GetLatestReleaseVersion)
	router.POST("/releaseVersion", _releaseVersionController.CreateReleaseVersion)

	router.Run(":8080")
}
