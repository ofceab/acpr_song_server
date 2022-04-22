package release_version_service

import (
	"acpr_songs_server/dal"
	"acpr_songs_server/models"
	"os"
)

type releaseVersionServiceImpl struct {
	// Store of `ReleaseVersion`
	releaseVersionDataAccessLayer dal.IReleaseVersionDatabaseAccessLayer
	// Current cached `ReleaseVersion`
	currentLatestVersion *models.ReleaseVersion
}

func (r *releaseVersionServiceImpl) GetReleaseVersions() []models.ReleaseVersion {
	_r, _err := r.releaseVersionDataAccessLayer.FetchReleaseVersions()
	if _err != nil {
		return []models.ReleaseVersion{}
	}
	return _r
}

func (r *releaseVersionServiceImpl) GetLatestReleaseVersion() models.ReleaseVersion {
	if r.currentLatestVersion != nil {
		// use cached value
		return *r.currentLatestVersion
	} else {
		// Query from DB
		_latestVersion, _ := r.releaseVersionDataAccessLayer.FetchLatestReleaseVersion()
		// Set cached value
		r.setCurrentReleaseVersion(_latestVersion)

		return *r.currentLatestVersion
	}
}

func (r *releaseVersionServiceImpl) CreateReleaseVersion() models.ReleaseVersion {

	_currentLatestVersion, err := r.releaseVersionDataAccessLayer.CreateReleaseVersion()
	if err != nil {
		os.Exit(1)
	}
	r.setCurrentReleaseVersion(_currentLatestVersion)

	return _currentLatestVersion
}

func (r *releaseVersionServiceImpl) DeleteReleaseVersion(releaseVersionId uint) models.ReleaseVersion {
	_r, err := r.releaseVersionDataAccessLayer.DeleteReleaseVersion(releaseVersionId)

	if err != nil {
		os.Exit(1)
	}
	return _r
}

func (r *releaseVersionServiceImpl) setCurrentReleaseVersion(p models.ReleaseVersion) int {
	*r.currentLatestVersion = p
	return 0
}
