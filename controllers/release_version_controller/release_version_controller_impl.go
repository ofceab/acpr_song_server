package release_version_controller

import (
	"acpr_songs_server/service/release_version_service"
	"net/http"
	"strconv"

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

func (p *releaseVersionControllerImpl) DeleteReleaseVersion(c *gin.Context) {
	_releaseVersionId := c.Param("releaseVersionId")

	_relCon, err := strconv.ParseUint(_releaseVersionId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Provide a valid ReleaseVersionId"})
		return
	}
	_nVersion := p.releaseVersionService.DeleteReleaseVersion(uint(_relCon))
	c.JSON(http.StatusCreated, _nVersion)
}
