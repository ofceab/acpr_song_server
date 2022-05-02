package song_service

import (
	"acpr_songs_server/core/constants"
	"acpr_songs_server/dal/dal_interfaces"
	"acpr_songs_server/models"
	"acpr_songs_server/service/release_version_service"
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

func (s *songServiceImpl) FetchSongsPerVersionId(sv uint) ([]models.Song, error) {
	_songs, err := s.songDal.FetchSongsPerVersionId(sv)

	if err != nil {
		return []models.Song{}, err
	}
	return _songs, nil
}

func (s *songServiceImpl) AddSong(sn *models.Song, releaseVersion uint) (models.Song, error) {

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
