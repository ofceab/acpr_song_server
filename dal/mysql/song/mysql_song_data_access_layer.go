package dal

import (
	"acpr_songs_server/dal"
	"acpr_songs_server/models"
	"strings"

	"gorm.io/gorm"
)

type MysqlSongDataAccessLayer struct {
	dbConnection *gorm.DB
	rDal         dal.IReleaseVersionDatabaseAccessLayer
}

// Save song in store
func (p *MysqlSongDataAccessLayer) SaveSong(s *models.Song, releaseVersion uint) (models.Song, error) {
	_result := p.dbConnection.Model(&models.Song{}).Create(s)
	if _result.Error != nil {
		return models.Song{}, _result.Error
	}
	return *s, nil
}

// Fetch songs
func (s *MysqlSongDataAccessLayer) FetchSongs() ([]models.Song, error) {
	// The idea for implementing the feature is to do a merge of all song's version

	// Fetch release version
	releaseVersions, err := s.rDal.FetchReleaseVersions()
	if err == nil {
		return []models.Song{}, err
	}

	fullSongs := make([]models.Song, 99999)

	// Fetch songs by release version
	for _, _releaseVersion := range releaseVersions {
		_fetchedSongs, _ := s.FetchSongsPerVersionId(_releaseVersion.ID)
		mergeSongs(&fullSongs, _fetchedSongs)
	}
	return fullSongs, nil
}

// Fetch all sounds per version id for fetching release song of a certain `version Id`
func (s *MysqlSongDataAccessLayer) FetchSongsPerVersionId(releaseVersion uint) ([]models.Song, error) {
	songs := new([]models.Song)
	_r := s.dbConnection.Where("ReleaseVersion = ?", releaseVersion).Find(songs)
	if _r.Error != nil {
		return []models.Song{}, _r.Error
	}
	return *songs, nil
}

func (s *MysqlSongDataAccessLayer) DeleteSong(songId uint) (models.Song, error) {
	newSong := new(models.Song)
	_r := s.dbConnection.Where("Id = ?", songId).Delete(newSong)
	if _r.Error != nil {
		return models.Song{}, _r.Error
	}
	return *newSong, nil
}

// Add into perform merge of song
func mergeSongs(s *[]models.Song, sn []models.Song) error {
	for _, song := range sn {
		for _, savedSong := range *s {
			if song.ReleaseVersion > savedSong.ReleaseVersion && strings.EqualFold(song.Title, savedSong.Title) {
				// Then add that item into the list
				*s = append(*s, song)
			}
		}
	}
	return nil
}
