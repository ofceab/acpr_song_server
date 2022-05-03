package main

import (
	"acpr_songs_server/controllers/release_version_controller"
	"acpr_songs_server/controllers/song_controller"
	"acpr_songs_server/dal/dal_interfaces"
	"acpr_songs_server/dal/mysql"
	"acpr_songs_server/service/release_version_service"

	"github.com/gin-gonic/gin"
)

func main() {
	_db := mysql.InitDb()
	// init song controller

	//Release version dependencies
	_releaseVersionDAL := dal_interfaces.NewIReleaseVersionDatabaseAccessLayer(_db)
	_releaseVersionService := release_version_service.New(_releaseVersionDAL)
	_releaseVersionController := release_version_controller.New(_releaseVersionService)

	// Song dependencies
	_songDal := dal_interfaces.NewISongDatabaseAccessLayer(_db, _releaseVersionDAL)
	_songController := song_controller.New(_songDal, _releaseVersionService)

	router := gin.Default()
	//Init Song routes
	// Get songs take latest version of songs

	router.GET("/v1/songs", _songController.FetchSongs)
	// Get songs by release version
	router.GET("/v1/songs/:releaseVersion", _songController.FetchSongsPerVersionId)
	// Add songs on a certain release version
	router.POST("/v1/songs/:releaseVersion", _songController.AddSong)
	// Delete songs on a certain release
	router.DELETE("/v1/songs/:songId", _songController.DeleteSong)

	//Init ReleaseVersion routes
	// Get all release versions
	router.GET("/v1/releaseVersions", _releaseVersionController.GetReleaseVersions)
	// Get latest release version
	router.GET("/v1/releaseVersions/latest", _releaseVersionController.GetLatestReleaseVersion)
	// Get release version info based on id
	router.GET("/v1/releaseVersions/:releaseVersion", _releaseVersionController.GetReleaseVersionById)
	// Create a release version
	router.POST("/v1/releaseVersions", _releaseVersionController.CreateReleaseVersion)
	// Delete a release version
	router.DELETE("/v1/releaseVersions/:releaseVersionId", _releaseVersionController.DeleteReleaseVersion)

	router.Run(":8081")
}
