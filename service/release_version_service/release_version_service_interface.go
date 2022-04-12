package release_version_service

import "acpr_songs_server/models"

type IReleaseVersionService interface {
	// Get all release version
	GetReleaseVersions() []models.ReleaseVersion
	// Get current latest release version
	GetLatestReleaseVersion() models.ReleaseVersion
	// Create a new Release Version
	CreateReleaseVersion() models.ReleaseVersion
}

func New() IReleaseVersionService {
	return &releaseVersionServiceImpl{
		localDB:              []models.ReleaseVersion{},
		currentLatestVersion: new(models.ReleaseVersion),
	}
}
