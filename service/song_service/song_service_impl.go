package song_service

import (
	"acpr_songs_server/core/constants"
	"acpr_songs_server/models"
	"fmt"
)

// An implementation of `ISongService`
type songServiceImpl struct {
	// base api for query songs
	apiBaseUrl string
	localDB    []models.Song
	lastIndex  int
}

func (s *songServiceImpl) FetchSongs() []models.Song {
	return s.localDB // for now an empty list
}

func (s *songServiceImpl) FetchSongsPerVersionId() []models.Song {
	return (*s).localDB // for now an empty list
}

func (s *songServiceImpl) AddSong(sn *models.Song, releaseVersion string) models.Song {
	// Generate key
	//TODO generate key
	sn.Id = "index_" + fmt.Sprint(s.lastIndex)
	s.lastIndex++

	if releaseVersion == constants.LATEST_RELEASE_KEY {
		// Use latest version
		//TODO
	}

	s.localDB = append(s.localDB, *sn)
	return *sn
}

func (s *songServiceImpl) DeleteSong(i string) (models.Song, error) {
	song := new(models.Song)

	var newSongList []models.Song

	for _, _song := range s.localDB {
		if _song.Id != i {
			newSongList = append(newSongList, _song)
		} else {
			*song = _song
		}
	}
	if song != nil {
		return *song, nil
	}

	return *song, fmt.Errorf(constants.ITEM_NOT_FOUND)
}
