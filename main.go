package main

import (
	notfoundcontroller "acpr_songs_server/controllers/not_found_controller"
	"acpr_songs_server/controllers/release_version_controller"
	"acpr_songs_server/controllers/song_controller"
	"acpr_songs_server/dal/dal_interfaces"
	"acpr_songs_server/dal/mysql"
	"acpr_songs_server/service/release_version_service"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func main() {
	uuid.Must(uuid.NewV4())

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
	v1 := router.Group("/v1")
	{
		// Get all song merged by version
		v1.GET("/songs", _songController.FetchSongs)
		// Get songs by release version
		v1.GET("/songs/:releaseVersionId", _songController.FetchSongsPerVersionId)
		// Add songs on a certain release version
		v1.POST("/songs/:releaseVersionId", _songController.AddSong)
		// Update song
		v1.PATCH("/songs/:releaseVersionId", _songController.UpdateSong)
		// Delete songs based on songId
		v1.DELETE("/songs/:songId", _songController.DeleteSong)

		//Init ReleaseVersion routes
		// Get all release versions
		v1.GET("/releaseVersions", _releaseVersionController.GetReleaseVersions)
		// Get latest release version
		v1.GET("/releaseVersions/latest", _releaseVersionController.GetLatestReleaseVersion)
		// Get release version info based on id
		v1.GET("/releaseVersions/:releaseVersionId", _releaseVersionController.GetReleaseVersionById)
		// Create a release version
		v1.POST("/releaseVersions", _releaseVersionController.CreateReleaseVersion)
		// Delete a release version
		v1.DELETE("/releaseVersions/:releaseVersionId", _releaseVersionController.DeleteReleaseVersion)
	}

	router.NoRoute(notfoundcontroller.NoFoundRouteHandler)

	router.Run(":8082")
}
