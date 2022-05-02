package dal_interfaces

import (
	"acpr_songs_server/models"
	"strings"

	"gorm.io/gorm"
)

type MysqlSongDataAccessLayer struct {
	DbConnection      *gorm.DB
	ReleaseVersionDal IReleaseVersionDatabaseAccessLayer
}

// Save song in store
func (p *MysqlSongDataAccessLayer) SaveSong(s *models.Song, releaseVersion uint) (models.Song, error) {
	s.ReleaseVersionId = releaseVersion
	_result := p.DbConnection.Model(&models.Song{}).Create(s)
	if _result.Error != nil {
		return models.Song{}, _result.Error
	}
	return *s, nil
}

// Fetch songs
func (s *MysqlSongDataAccessLayer) FetchSongs() ([]models.Song, error) {
	// The idea for implementing the feature is to do a merge of all song's version

	// Fetch release version
	releaseVersions, err := s.ReleaseVersionDal.FetchReleaseVersions()
	if err != nil {
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
	_r := s.DbConnection.Find(songs, models.Song{ReleaseVersionId: releaseVersion})
	if _r.Error != nil {
		return []models.Song{}, _r.Error
	}
	return *songs, nil
}

func (s *MysqlSongDataAccessLayer) DeleteSong(songId uint) (models.Song, error) {
	newSong := new(models.Song)
	_r := s.DbConnection.Delete(newSong, songId)
	if _r.RowsAffected == 0 {
		return models.Song{}, _r.Error
	}
	return *newSong, nil
}

// Add into perform merge of song
func mergeSongs(s *[]models.Song, sn []models.Song) {
	for _, song := range sn {
		for _, savedSong := range *s {
			if song.ReleaseVersionId > savedSong.ReleaseVersionId && strings.EqualFold(song.Title, savedSong.Title) {
				// Then add that item into the list
				addSongInList(s, song)
			}

		}
	}
}

// Merge song
func addSongInList(s *[]models.Song, sn models.Song) {
	_tempSngs := []models.Song{}

	for _, _song := range *s {
		if _song.Id != sn.Id {
			_tempSngs = append(_tempSngs, _song)
		}
	}
	_tempSngs = append(_tempSngs, sn)

	// update the song slice
	*s = _tempSngs
}
