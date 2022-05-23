package release_version_service

import (
	"acpr_songs_server/core/errors"
	"acpr_songs_server/dal/dal_interfaces"
	"acpr_songs_server/models"
)

type releaseVersionServiceImpl struct {
	// Store of `ReleaseVersion`
	releaseVersionDataAccessLayer dal_interfaces.IReleaseVersionDatabaseAccessLayer
	// Current cached `ReleaseVersion`
}

func (r *releaseVersionServiceImpl) GetReleaseVersions() ([]models.ReleaseVersion, errors.ReleaseVersionError) {
	_r, _err := r.releaseVersionDataAccessLayer.FetchReleaseVersions()
	if _err.ErrorCode != 0 {
		return []models.ReleaseVersion{}, _err
	}
	return _r, _err
}
func (r *releaseVersionServiceImpl) GetReleaseVersionById(id uint) (models.ReleaseVersion, errors.ReleaseVersionError) {
	_r, _err := r.releaseVersionDataAccessLayer.FetchReleaseVersionById(id)
	if _err.ErrorCode != 0 {
		return models.ReleaseVersion{}, _err
	}
	return _r, _err
}

func (r *releaseVersionServiceImpl) GetLatestReleaseVersion() (models.ReleaseVersion, errors.ReleaseVersionError) {
	// Query from DB
	_latestVersion, _err := r.releaseVersionDataAccessLayer.FetchLatestReleaseVersion()
	if _err.ErrorCode != 0 {
		return models.ReleaseVersion{}, _err
	}
	return _latestVersion, _err
}

func (r *releaseVersionServiceImpl) CreateReleaseVersion() (models.ReleaseVersion, errors.ReleaseVersionError) {

	_currentLatestVersion, err := r.releaseVersionDataAccessLayer.CreateReleaseVersion()
	if err.ErrorCode != 0 {
		return models.ReleaseVersion{}, err
	}

	return _currentLatestVersion, err
}

func (r *releaseVersionServiceImpl) DeleteReleaseVersion(releaseVersionId uint) (models.ReleaseVersion, errors.ReleaseVersionError) {
	_r, err := r.releaseVersionDataAccessLayer.DeleteReleaseVersion(releaseVersionId)

	if err.ErrorCode != 0 {
		return models.ReleaseVersion{}, err
	}

	return _r, err
}
