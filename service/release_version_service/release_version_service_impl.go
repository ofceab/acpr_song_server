package release_version_service

import (
	"acpr_songs_server/models"
	"time"
)

type releaseVersionServiceImpl struct {
	// Store of `ReleaseVersion`
	localDB []models.ReleaseVersion
	// Current cached `ReleaseVersion`
	currentLatestVersion *models.ReleaseVersion
}

func (r *releaseVersionServiceImpl) GetReleaseVersions() []models.ReleaseVersion {
	return r.localDB
}

func (r *releaseVersionServiceImpl) GetLatestReleaseVersion() models.ReleaseVersion {
	if r.currentLatestVersion != nil {
		// use cached value
		return *r.currentLatestVersion
	} else {
		// Query from DB
		if len(r.localDB) > 0 {
			r.setCurrentReleaseVersion(r.localDB[0])

			return *r.currentLatestVersion
		}

		return models.ReleaseVersion{}
	}
}

func (r *releaseVersionServiceImpl) CreateReleaseVersion() models.ReleaseVersion {

	_currentLatestVersion := new(models.ReleaseVersion)

	if r.currentLatestVersion != nil {
		_currentLatestVersion = r.currentLatestVersion
		_currentLatestVersion.Id += 1
	} else {
		*_currentLatestVersion = r.GetLatestReleaseVersion()
		_currentLatestVersion.Id += 1
	}

	//TODO change it later
	_newReleaseVersion := models.ReleaseVersion{Id: _currentLatestVersion.Id, CreatedAt: time.Now()}
	// Add in store
	r.localDB = append(r.localDB, _newReleaseVersion)

	r.setCurrentReleaseVersion(*_currentLatestVersion)

	return _newReleaseVersion
}

func (r *releaseVersionServiceImpl) setCurrentReleaseVersion(p models.ReleaseVersion) int {
	*r.currentLatestVersion = p
	return 0
}
