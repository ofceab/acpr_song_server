package release_version_controller

import (
	"acpr_songs_server/core/constants"
	"acpr_songs_server/core/errors"
	"acpr_songs_server/service/release_version_service"
	"acpr_songs_server/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type releaseVersionControllerImpl struct {
	releaseVersionService release_version_service.IReleaseVersionService
}

func (p *releaseVersionControllerImpl) GetReleaseVersions(c *gin.Context) {
	_versions, _err := p.releaseVersionService.GetReleaseVersions()
	if _err.ErrorCode != 0 {
		utils.SendResponse(c, errors.AppError(_err))
		return
	}
	c.JSON(http.StatusOK, _versions)
}

func (p *releaseVersionControllerImpl) GetReleaseVersionById(c *gin.Context) {
	_rId, _convErr := strconv.ParseUint(c.Param(constants.RELEASE_VERSION_KEY), 10, 32)

	if _convErr != nil {
		//TODO migrate this to service layer
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide a valid ReleaseVersionId"})
	}
	_versions, _err := p.releaseVersionService.GetReleaseVersionById(uint(_rId))
	if _err.ErrorCode != 0 {
		utils.SendResponse(c, errors.AppError(_err))
		return
	}
	c.JSON(http.StatusOK, _versions)
}

// Get current latest release version
func (p *releaseVersionControllerImpl) GetLatestReleaseVersion(c *gin.Context) {
	_cVersion, _err := p.releaseVersionService.GetLatestReleaseVersion()
	if _err.ErrorCode != 0 {
		utils.SendResponse(c, errors.AppError(_err))
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

	if _err.ErrorCode != 0 {
		utils.SendResponse(c, errors.AppError(_err))
		return
	}

	c.JSON(http.StatusCreated, _nVersion)
}

func (p *releaseVersionControllerImpl) DeleteReleaseVersion(c *gin.Context) {
	_releaseVersionId, err := strconv.ParseUint(c.Param(constants.RELEASE_VERSION_KEY), 10, 32)

	if err != nil {
		//TODO
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provide a valid ReleaseVersionId"})
		return
	}

	_nVersion, _err := p.releaseVersionService.DeleteReleaseVersion(uint(_releaseVersionId))

	if _err.ErrorCode != 0 {
		utils.SendResponse(c, errors.AppError(_err))
		return
	}
	if _nVersion.ID != 0 {
		c.JSON(http.StatusOK, _nVersion)
		return
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}
