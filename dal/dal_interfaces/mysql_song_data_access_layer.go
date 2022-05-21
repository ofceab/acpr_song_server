package dal_interfaces

import (
	dataformat "acpr_songs_server/data_format"
	"acpr_songs_server/models"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

var OMIT_SONG_FIELD = []string{"ID"}

type MysqlSongDataAccessLayer struct {
	DbConnection      *gorm.DB
	ReleaseVersionDal IReleaseVersionDatabaseAccessLayer
}

// Save song in store
func (p *MysqlSongDataAccessLayer) SaveSong(s *dataformat.CreateSong, releaseVersion uint) (models.Song, error) {
	_songUUID, _err := uuid.NewV4()
	if _err != nil {
		return models.Song{}, _err
	}

	_s := models.Song{Title: s.Title, Lyrics: s.Lyrics, AudioUrl: s.AudioUrl, ReleaseVersionId: releaseVersion, CreatedAt: time.Now(), SongUniqueId: _songUUID.String()}
	_result := p.DbConnection.Omit(OMIT_SONG_FIELD...).Create(&_s)
	if _result.Error != nil {
		return models.Song{}, _result.Error
	}
	return _s, nil
}

func (p *MysqlSongDataAccessLayer) UpdateSong(s *dataformat.UpdateSong, releaseVersion uint) (models.Song, error) {
	_s := models.Song{Title: s.Title, Lyrics: s.Lyrics, AudioUrl: s.AudioUrl, ReleaseVersionId: releaseVersion, CreatedAt: time.Now(), SongUniqueId: s.SongUniqueId}
	_result := p.DbConnection.Omit(OMIT_SONG_FIELD...).Create(&_s)
	if _result.Error != nil {
		return models.Song{}, _result.Error
	}
	return _s, nil
}

// Fetch songs
func (s *MysqlSongDataAccessLayer) FetchSongs() ([]models.Song, error) {
	// The idea for implementing the feature is to do a merge of all song's version

	// Retrieve all songs
	fullSongs, _err := s.fetchAllSongs()
	if _err != nil {
		return []models.Song{}, _err
	}

	_merge_songs := []models.Song{}

	_songs := mergeSongs(&fullSongs, _merge_songs)

	return _songs, nil
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
func mergeSongs(s *[]models.Song, sn []models.Song) []models.Song {
	for _, _song := range *s {
		_status, _index := checkIfSongInSlice(sn, _song)

		if _status {
			if _song.ReleaseVersionId > sn[_index].ReleaseVersionId {
				// Replace the existing song within the list
				replaceSongs(sn, _song, _index)
			}
		} else {
			sn = append(sn, _song)
		}
	}

	return sn
}

func checkIfSongInSlice(s []models.Song, song models.Song) (bool, int) {
	for _i, _song := range s {
		if _song.SongUniqueId == song.SongUniqueId {
			return true, _i
		}
	}
	return false, -1
}

func replaceSongs(s []models.Song, song models.Song, index int) {
	s[index] = song
}

// // Merge song
// func addSongInList(s *[]models.Song, sn models.Song) {
// 	_tempSngs := []models.Song{}

// 	for _, _song := range *s {
// 		if _song.ID != sn.ID {
// 			_tempSngs = append(_tempSngs, _song)
// 		}
// 	}
// 	_tempSngs = append(_tempSngs, sn)

// 	// update the song slice
// 	*s = _tempSngs
// }

// Fetch all songs from storage
func (s *MysqlSongDataAccessLayer) fetchAllSongs() ([]models.Song, error) {
	_songs := new([]models.Song)

	_result := s.DbConnection.Find(_songs)
	if _result.Error != nil {
		return []models.Song{}, _result.Error
	}
	return *_songs, nil
}
