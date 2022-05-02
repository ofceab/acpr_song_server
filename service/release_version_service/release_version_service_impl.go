package release_version_service

import (
	"acpr_songs_server/dal/dal_interfaces"
	"acpr_songs_server/models"
	"os"
)

type releaseVersionServiceImpl struct {
	// Store of `ReleaseVersion`
	releaseVersionDataAccessLayer dal_interfaces.IReleaseVersionDatabaseAccessLayer
	// Current cached `ReleaseVersion`
}

func (r *releaseVersionServiceImpl) GetReleaseVersions() []models.ReleaseVersion {
	_r, _err := r.releaseVersionDataAccessLayer.FetchReleaseVersions()
	if _err != nil {
		return []models.ReleaseVersion{}
	}
	return _r
}

func (r *releaseVersionServiceImpl) GetLatestReleaseVersion() models.ReleaseVersion {
	// Query from DB
	_latestVersion, _ := r.releaseVersionDataAccessLayer.FetchLatestReleaseVersion()

	return _latestVersion
}

func (r *releaseVersionServiceImpl) CreateReleaseVersion() models.ReleaseVersion {

	_currentLatestVersion, err := r.releaseVersionDataAccessLayer.CreateReleaseVersion()
	if err != nil {
		os.Exit(1)
	}

	return _currentLatestVersion
}

func (r *releaseVersionServiceImpl) DeleteReleaseVersion(releaseVersionId uint) models.ReleaseVersion {
	_r, err := r.releaseVersionDataAccessLayer.DeleteReleaseVersion(releaseVersionId)

	if err != nil {
		os.Exit(1)
	}

	return _r
}
