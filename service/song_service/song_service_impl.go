package song_service

import (
	"acpr_songs_server/core/constants"
	"acpr_songs_server/dal/dal_interfaces"
	dataformat "acpr_songs_server/data_format"
	"acpr_songs_server/models"
	"acpr_songs_server/service/release_version_service"
	"fmt"
)

// An implementation of `ISongService`
type songServiceImpl struct {
	// base api for query songs
	songDal               dal_interfaces.ISongDatabaseAccessLayer
	releaseVersionService release_version_service.IReleaseVersionService
}

func (s *songServiceImpl) FetchSongs() ([]models.Song, error) {
	songs, err := s.songDal.FetchSongs()
	if err != nil {
		return []models.Song{}, err
	}
	return songs, nil
}

func (s *songServiceImpl) FetchSongsPerSongUniqueId(snUID string) ([]models.Song, error) {
	songs, err := s.songDal.FetchSongsPerSongUniqueId(snUID)
	if err != nil {
		return []models.Song{}, err
	}
	return songs, nil
}

func (s *songServiceImpl) UpdateSong(p *dataformat.UpdateSong, releaseVersion uint) (models.Song, error) {
	// Retrieve first releases versions
	_releaseVersions, _err := s.releaseVersionService.GetReleaseVersions()
	if _err != nil {
		return models.Song{}, _err
	}

	_rvExist := false
	for _, _rv := range _releaseVersions {
		if _rv.ID == releaseVersion {
			_rvExist = true
		}
	}

	if !_rvExist {
		return models.Song{}, fmt.Errorf("release version provided doesn't exist")
	}

	song, err := s.songDal.UpdateSong(p, releaseVersion)
	if err != nil {
		return models.Song{}, err
	}
	return song, nil
}

func (s *songServiceImpl) FetchSongsPerVersionId(sv uint) ([]models.Song, error) {
	_songs, err := s.songDal.FetchSongsPerVersionId(sv)

	if err != nil {
		return []models.Song{}, err
	}
	return _songs, nil
}

func (s *songServiceImpl) AddSong(sn *dataformat.CreateSong, releaseVersion uint) (models.Song, error) {

	if releaseVersion == constants.LATEST_RELEASE_KEY {
		_r, err := s.releaseVersionService.GetLatestReleaseVersion()
		if err != nil {
			return models.Song{}, err
		}
		_sn, _err := s.songDal.SaveSong(sn, _r.ID)
		if _err != nil {
			return models.Song{}, _err
		}
		return _sn, nil
	}
	_s, _err := s.songDal.SaveSong(sn, releaseVersion)
	if _err != nil {
		return models.Song{}, _err
	}
	return _s, nil
}

func (s *songServiceImpl) DeleteSong(sId uint) (models.Song, error) {
	_s, _err := s.songDal.DeleteSong(sId)
	if _err != nil {
		return models.Song{}, _err
	}
	return _s, nil
}
