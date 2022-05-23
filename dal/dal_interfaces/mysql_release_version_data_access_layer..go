package dal_interfaces

import (
	"acpr_songs_server/core/errors"
	"acpr_songs_server/models"
	"net/http"

	"gorm.io/gorm"
)

type MysqlReleaseVersionDataAccessLayer struct {
	DbConnection *gorm.DB
}

// Create a ReleaseVersion
func (p *MysqlReleaseVersionDataAccessLayer) CreateReleaseVersion() (models.ReleaseVersion, errors.ReleaseVersionError) {
	_newReleaseVersion := new(models.ReleaseVersion)
	_result := p.DbConnection.Model(&models.ReleaseVersion{}).Create(&_newReleaseVersion)
	if _result.Error != nil {
		return models.ReleaseVersion{}, errors.ReleaseVersionError(errors.GetInternalError())
	}
	return *_newReleaseVersion, errors.ReleaseVersionError{}
}

// Fetch all release version
func (p *MysqlReleaseVersionDataAccessLayer) FetchReleaseVersions() ([]models.ReleaseVersion, errors.ReleaseVersionError) {
	r := new([]models.ReleaseVersion)
	_r := p.DbConnection.Order("id DESC").Find(r)
	if _r.Error != nil {
		return nil, errors.ReleaseVersionError(errors.GetInternalError())
	}
	return *r, errors.ReleaseVersionError{}
}

// Delete a particular release version
func (p *MysqlReleaseVersionDataAccessLayer) DeleteReleaseVersion(songId uint) (models.ReleaseVersion, errors.ReleaseVersionError) {
	r := new(models.ReleaseVersion)
	_r := p.DbConnection.Delete(r, songId)

	if _r.RowsAffected == 0 {
		return models.ReleaseVersion{}, errors.ReleaseVersionError{Message: errors.RELEASE_VERSION__OF_ID_DOESNT_EXIST_ERROR, ErrorCode: http.StatusBadRequest}
	}
	r.ID = songId
	return *r, errors.ReleaseVersionError{}
}

// Fetch latest Release Version
func (p *MysqlReleaseVersionDataAccessLayer) FetchReleaseVersionById(id uint) (models.ReleaseVersion, errors.ReleaseVersionError) {
	r := new(models.ReleaseVersion)
	_r := p.DbConnection.First(r, id)
	if _r.RowsAffected == 0 {
		// no Record found
		return models.ReleaseVersion{}, errors.ReleaseVersionError{Message: errors.RELEASE_VERSION__OF_ID_DOESNT_EXIST_ERROR, ErrorCode: http.StatusNoContent}
	}
	return *r, errors.ReleaseVersionError{}
}

// Fetch latest Release Version
func (p *MysqlReleaseVersionDataAccessLayer) FetchLatestReleaseVersion() (models.ReleaseVersion, errors.ReleaseVersionError) {
	r := new(models.ReleaseVersion)
	_r := p.DbConnection.Last(r)
	if _r.RowsAffected == 0 {
		// no Record found
		return models.ReleaseVersion{}, errors.ReleaseVersionError{Message: errors.RELEASE_VERSION__OF_ID_DOESNT_EXIST_ERROR, ErrorCode: http.StatusNoContent}
	}
	return *r, errors.ReleaseVersionError{}
}
