package release_version_controller

import (
	"acpr_songs_server/service/release_version_service"

	"github.com/gin-gonic/gin"
)

type IReleaseVersionController interface {
	// Get all release version
	GetReleaseVersions(c *gin.Context)
	GetReleaseVersionById(c *gin.Context)
	// Get current latest release version
	GetLatestReleaseVersion(c *gin.Context)
	// Create a new Release Version
	CreateReleaseVersion(c *gin.Context)
	// Delete a relase version
	DeleteReleaseVersion(c *gin.Context)
}

func New(releaseVersionService release_version_service.IReleaseVersionService) IReleaseVersionController {
	return &releaseVersionControllerImpl{
		releaseVersionService: releaseVersionService,
	}
}
