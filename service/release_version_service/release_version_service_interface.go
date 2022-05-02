package release_version_service

import (
	"acpr_songs_server/dal/dal_interfaces"
	"acpr_songs_server/models"
)

type IReleaseVersionService interface {
	// Get all release version
	GetReleaseVersions() ([]models.ReleaseVersion, error)
	// Get current latest release version
	GetLatestReleaseVersion() (models.ReleaseVersion, error)
	// Create a new Release Version
	CreateReleaseVersion() (models.ReleaseVersion, error)
	// Delete a release version
	DeleteReleaseVersion(releaseVersionId uint) (models.ReleaseVersion, error)
}

func New(releaseVersionDal dal_interfaces.IReleaseVersionDatabaseAccessLayer) IReleaseVersionService {
	return &releaseVersionServiceImpl{
		releaseVersionDataAccessLayer: releaseVersionDal,
	}
}
