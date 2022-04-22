package release_version_service

import (
	"acpr_songs_server/dal"
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

func New(releaseVersionDal dal.IReleaseVersionDatabaseAccessLayer) IReleaseVersionService {
	return &releaseVersionServiceImpl{
		releaseVersionDataAccessLayer: releaseVersionDal,
		currentLatestVersion:          new(models.ReleaseVersion),
	}
}
