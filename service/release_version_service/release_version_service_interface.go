package release_version_service

import (
	"acpr_songs_server/core/errors"
	"acpr_songs_server/dal/dal_interfaces"
	"acpr_songs_server/models"
)

type IReleaseVersionService interface {
	// Get all release version
	GetReleaseVersions() ([]models.ReleaseVersion, errors.ReleaseVersionError)
	// Get release version of a certain Id
	GetReleaseVersionById(id uint) (models.ReleaseVersion, errors.ReleaseVersionError)
	// Get current latest release version
	GetLatestReleaseVersion() (models.ReleaseVersion, errors.ReleaseVersionError)
	// Create a new Release Version
	CreateReleaseVersion() (models.ReleaseVersion, errors.ReleaseVersionError)
	// Delete a release version
	DeleteReleaseVersion(releaseVersionId uint) (models.ReleaseVersion, errors.ReleaseVersionError)
}

func New(releaseVersionDal dal_interfaces.IReleaseVersionDatabaseAccessLayer) IReleaseVersionService {
	return &releaseVersionServiceImpl{
		releaseVersionDataAccessLayer: releaseVersionDal,
	}
}
