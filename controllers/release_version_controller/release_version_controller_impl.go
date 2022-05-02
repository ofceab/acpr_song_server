package release_version_controller

import (
	"acpr_songs_server/core/constants"
	"acpr_songs_server/service/release_version_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type releaseVersionControllerImpl struct {
	releaseVersionService release_version_service.IReleaseVersionService
}

func (p *releaseVersionControllerImpl) GetReleaseVersions(c *gin.Context) {
	_versions, _err := p.releaseVersionService.GetReleaseVersions()
	if _err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": _err.Error()})
		return
	}
	c.JSON(http.StatusOK, _versions)
}

// Get current latest release version
func (p *releaseVersionControllerImpl) GetLatestReleaseVersion(c *gin.Context) {
	_cVersion, _err := p.releaseVersionService.GetLatestReleaseVersion()
	if _err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": _err.Error()})
		return
	}
	// check release version id if 0 then no last version yet
	if _cVersion.ID != 0 {
		c.JSON(http.StatusOK, _cVersion)
		return
	} else {
		// no latest version found
		c.JSON(http.StatusNoContent, nil)
	}
}

// Create a new Release Version
func (p *releaseVersionControllerImpl) CreateReleaseVersion(c *gin.Context) {
	_nVersion, _err := p.releaseVersionService.CreateReleaseVersion()

	if _err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": _err.Error()})
		return
	}

	c.JSON(http.StatusCreated, _nVersion)
}

func (p *releaseVersionControllerImpl) DeleteReleaseVersion(c *gin.Context) {
	_releaseVersionId, err := strconv.ParseUint(c.Param(constants.RELEASE_VERSION_KEY), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide a valid ReleaseVersionId"})
		return
	}

	_nVersion, _err := p.releaseVersionService.DeleteReleaseVersion(uint(_releaseVersionId))

	if _err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if _nVersion.ID != 0 {
		c.JSON(http.StatusOK, _nVersion)
		return
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}
