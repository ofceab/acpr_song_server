package release_version_service

import (
	"acpr_songs_server/dal/dal_interfaces"
	"acpr_songs_server/models"
)

type releaseVersionServiceImpl struct {
	// Store of `ReleaseVersion`
	releaseVersionDataAccessLayer dal_interfaces.IReleaseVersionDatabaseAccessLayer
	// Current cached `ReleaseVersion`
}

func (r *releaseVersionServiceImpl) GetReleaseVersions() ([]models.ReleaseVersion, error) {
	_r, _err := r.releaseVersionDataAccessLayer.FetchReleaseVersions()
	if _err != nil {
		return []models.ReleaseVersion{}, _err
	}
	return _r, nil
}
func (r *releaseVersionServiceImpl) GetReleaseVersionById(id uint) (models.ReleaseVersion, error) {
	_r, _err := r.releaseVersionDataAccessLayer.FetchReleaseVersionById(id)
	if _err != nil {
		return models.ReleaseVersion{}, _err
	}
	return _r, nil
}

func (r *releaseVersionServiceImpl) GetLatestReleaseVersion() (models.ReleaseVersion, error) {
	// Query from DB
	_latestVersion, _err := r.releaseVersionDataAccessLayer.FetchLatestReleaseVersion()
	if _err != nil {
		return models.ReleaseVersion{}, _err
	}
	return _latestVersion, nil
}

func (r *releaseVersionServiceImpl) CreateReleaseVersion() (models.ReleaseVersion, error) {

	_currentLatestVersion, err := r.releaseVersionDataAccessLayer.CreateReleaseVersion()
	if err != nil {
		return models.ReleaseVersion{}, err
	}

	return _currentLatestVersion, nil
}

func (r *releaseVersionServiceImpl) DeleteReleaseVersion(releaseVersionId uint) (models.ReleaseVersion, error) {
	_r, err := r.releaseVersionDataAccessLayer.DeleteReleaseVersion(releaseVersionId)

	if err != nil {
		return models.ReleaseVersion{}, err
	}

	return _r, nil
}
