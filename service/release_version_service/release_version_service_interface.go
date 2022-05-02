package release_version_service

import (
	"acpr_songs_server/dal/dal_interfaces"
	"acpr_songs_server/models"
)

type IReleaseVersionService interface {
	// Get all release version
	GetReleaseVersions() []models.ReleaseVersion
	// Get current latest release version
	GetLatestReleaseVersion() models.ReleaseVersion
	// Create a new Release Version
	CreateReleaseVersion() models.ReleaseVersion
	// Delete a release version
	DeleteReleaseVersion(releaseVersionId uint) models.ReleaseVersion
}

func New(releaseVersionDal dal_interfaces.IReleaseVersionDatabaseAccessLayer) IReleaseVersionService {
	return &releaseVersionServiceImpl{
		releaseVersionDataAccessLayer: releaseVersionDal,
	}
}
