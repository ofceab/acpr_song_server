package main

import (
	"acpr_songs_server/controllers/release_version_controller"
	"acpr_songs_server/controllers/song_controller"
	"acpr_songs_server/dal"
	"acpr_songs_server/dal/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	_db := mysql.InitDb()
	// init song controller
	_songController := song_controller.New()
	// Init release version dependencies
	_releaseVersionDAL := dal.NewIReleaseVersionDatabaseAccessLayer(_db)

	_releaseVersionController := release_version_controller.New(_releaseVersionDAL)
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
	// Get all release versions
	router.GET("/v1/releaseVersions", _releaseVersionController.GetReleaseVersions)
	// Get latest release version
	router.GET("/v1/releaseVersions/latest", _releaseVersionController.GetLatestReleaseVersion)
	// Create a release version
	router.POST("/v1/releaseVersions", _releaseVersionController.CreateReleaseVersion)
	// Delete a release version
	router.DELETE("/v1/releaseVersions/:releaseVersionId")

	router.Run(":8080")
}
