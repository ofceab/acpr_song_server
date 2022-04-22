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
	// check release version id if 0 then no last version yet
	if _cVersion.ID != 0 {
		c.JSON(http.StatusOK, _cVersion)
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
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

	if _nVersion.ID != 0 {
		c.JSON(http.StatusOK, _nVersion)
		return
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}
