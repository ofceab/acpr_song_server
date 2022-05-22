package dal_interfaces

import (
	dataformat "acpr_songs_server/data_format"
	"acpr_songs_server/models"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

var OMIT_SONG_FIELD = []string{"ID"}

const (
	canBeAdd    = 0
	cannotBeAdd = 1
	doesntExist = 3
)

type MysqlSongDataAccessLayer struct {
	DbConnection      *gorm.DB
	ReleaseVersionDal IReleaseVersionDatabaseAccessLayer
}

// Save song in store
func (p *MysqlSongDataAccessLayer) SaveSong(s *dataformat.CreateSong, releaseVersion uint) (models.Song, error) {
	// Ensure release version Exist
	_err := p.checkExistenceOfVersionRelease(releaseVersion)
	if _err != nil {
		return models.Song{}, _err
	}

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
	// Ensure release version Exist
	_err := p.checkExistenceOfVersionRelease(releaseVersion)
	if _err != nil {
		return models.Song{}, _err
	}

	// Ensure that release version has a different version compare to existence song
	_songs, _err := p.FetchSongsPerSongUniqueId(s.SongUniqueId)
	if _err != nil {
		return models.Song{}, _err
	}

	// Check if a song can be add / check if provided version is high that existing
	_status := compareSongReleaseVersion(_songs, releaseVersion, s.SongUniqueId)
	if _status == cannotBeAdd {
		return models.Song{}, fmt.Errorf("provide a higher version for adding a new version of a song")
	} else if _status == doesntExist {
		return models.Song{}, fmt.Errorf("invalid song_unique_id. can't add a new version for a song that doesn't exist")
	}

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

// Return `true` for knowing if a song with a certain release version can be add to store
func compareSongReleaseVersion(sn []models.Song, rv uint, snUID string) int {
	_status, _index := checkIfSongInSlice(sn, models.Song{SongUniqueId: snUID})
	if !_status {
		return doesntExist
	}

	_tempSong := sn[_index]

	if _tempSong.ReleaseVersionId < rv {
		return canBeAdd
	} else {
		return cannotBeAdd
	}
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

// Fetch all sounds per version id for fetching release song of a certain `version Id`
func (s *MysqlSongDataAccessLayer) FetchSongsPerSongUniqueId(snUID string) ([]models.Song, error) {
	songs := new([]models.Song)
	_r := s.DbConnection.Find(songs, models.Song{SongUniqueId: snUID})
	if _r.Error != nil {
		return []models.Song{}, _r.Error
	}
	return *songs, nil
}

func (s *MysqlSongDataAccessLayer) DeleteSong(songId uint) (models.Song, error) {
	newSong := new(models.Song)
	_r := s.DbConnection.Delete(newSong, songId)
	if _r.RowsAffected == 0 {
		return models.Song{}, fmt.Errorf("song with provided id doesnt exist. Be sure of id")
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

func (p *MysqlSongDataAccessLayer) checkExistenceOfVersionRelease(rv uint) error {
	_releasesVersions, _err := p.ReleaseVersionDal.FetchReleaseVersions()

	if _err != nil {
		// TODO for later
		return fmt.Errorf("an exception on data occured, retry later")
	}

	_status, _ := checkIfReleaseVersionInSlice(_releasesVersions, rv)
	if !_status {
		return fmt.Errorf("invalid releaseVersion, provide release version that exist")
	}
	return nil
}

func checkIfReleaseVersionInSlice(s []models.ReleaseVersion, rv uint) (bool, int) {
	for _i, _rv := range s {
		if _rv.ID == rv {
			return true, _i
		}
	}
	return false, -1
}

func replaceSongs(s []models.Song, song models.Song, index int) {
	s[index] = song
}

// Fetch all songs from storage
func (s *MysqlSongDataAccessLayer) fetchAllSongs() ([]models.Song, error) {
	_songs := new([]models.Song)

	_result := s.DbConnection.Find(_songs)
	if _result.Error != nil {
		return []models.Song{}, _result.Error
	}
	return *_songs, nil
}
