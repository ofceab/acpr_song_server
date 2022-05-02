package dal_interfaces

import (
	"acpr_songs_server/models"

	"gorm.io/gorm"
)

type MysqlReleaseVersionDataAccessLayer struct {
	DbConnection *gorm.DB
}

// Create a ReleaseVersion
func (p *MysqlReleaseVersionDataAccessLayer) CreateReleaseVersion() (models.ReleaseVersion, error) {
	_newReleaseVersion := new(models.ReleaseVersion)
	_result := p.DbConnection.Model(&models.ReleaseVersion{}).Create(&_newReleaseVersion)
	if _result.Error != nil {
		return models.ReleaseVersion{}, _result.Error
	}
	return *_newReleaseVersion, nil
}

// Fetch all release version
func (p *MysqlReleaseVersionDataAccessLayer) FetchReleaseVersions() ([]models.ReleaseVersion, error) {
	r := new([]models.ReleaseVersion)
	_r := p.DbConnection.Order("id DESC").Find(r)
	if _r.Error != nil {
		return nil, _r.Error
	}
	return *r, nil
}

// Delete a particular release version
func (p *MysqlReleaseVersionDataAccessLayer) DeleteReleaseVersion(releaseVersion uint) (models.ReleaseVersion, error) {
	r := new(models.ReleaseVersion)
	_r := p.DbConnection.Delete(r, releaseVersion)

	if _r.RowsAffected == 0 {
		return models.ReleaseVersion{}, _r.Error
	}
	r.ID = releaseVersion
	return *r, nil
}

// Fetch latest Release Version
func (p *MysqlReleaseVersionDataAccessLayer) FetchLatestReleaseVersion() (models.ReleaseVersion, error) {
	r := new(models.ReleaseVersion)
	_r := p.DbConnection.Last(r)
	if _r.RowsAffected == 0 {
		// no Record found
		return models.ReleaseVersion{}, _r.Error
	}
	return *r, nil
}
