package release_version_controller

import (
	"acpr_songs_server/service/release_version_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type releaseVersionControllerImpl struct {
	releaseVersionService release_version_service.IReleaseVersionService
}

func (p *releaseVersionControllerImpl) GetReleaseVersions(c *gin.Context) {
	_versions := p.releaseVersionService.GetReleaseVersions()

	c.JSON(http.StatusOK, _versions)
}

// Get current latest release version
func (p *releaseVersionControllerImpl) GetLatestReleaseVersion(c *gin.Context) {
	_cVersion := p.releaseVersionService.GetLatestReleaseVersion()

	c.JSON(http.StatusOK, _cVersion)
}

// Create a new Release Version
func (p *releaseVersionControllerImpl) CreateReleaseVersion(c *gin.Context) {
	_nVersion := p.releaseVersionService.CreateReleaseVersion()

	c.JSON(http.StatusCreated, _nVersion)
}
