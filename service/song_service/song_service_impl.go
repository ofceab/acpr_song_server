package song_service

import (
	"acpr_songs_server/core/constants"
	"acpr_songs_server/core/errors"
	"acpr_songs_server/dal/dal_interfaces"
	dataformat "acpr_songs_server/data_format"
	"acpr_songs_server/models"
	"acpr_songs_server/service/release_version_service"
	"net/http"
)

// An implementation of `ISongService`
type songServiceImpl struct {
	// base api for query songs
	songDal               dal_interfaces.ISongDatabaseAccessLayer
	releaseVersionService release_version_service.IReleaseVersionService
}

func (s *songServiceImpl) FetchSongs() ([]models.Song, errors.SongError) {
	songs, err := s.songDal.FetchSongs()
	if err.ErrorCode != 0 {
		return []models.Song{}, err
	}
	return songs, err // the error in this case is returned as zero value
}

func (s *songServiceImpl) FetchSongsPerSongUniqueId(snUID string) ([]models.Song, errors.SongError) {
	songs, err := s.songDal.FetchSongsPerSongUniqueId(snUID)
	if err.ErrorCode != 0 {
		return []models.Song{}, err
	}
	return songs, err
}

func (s *songServiceImpl) UpdateSong(p *dataformat.UpdateSong, releaseVersion uint) (models.Song, errors.SongError) {
	// Retrieve first releases versions
	_releaseVersions, _err := s.releaseVersionService.GetReleaseVersions()
	if _err.ErrorCode != 0 {
		return models.Song{}, errors.SongError(_err)
	}

	_rvExist := false
	for _, _rv := range _releaseVersions {
		if _rv.ID == releaseVersion {
			_rvExist = true
		}
	}

	if !_rvExist {
		return models.Song{}, errors.SongError{Message: errors.RELEASE_VERSION__OF_ID_DOESNT_EXIST_ERROR, ErrorCode: http.StatusBadRequest}
	}

	song, err := s.songDal.UpdateSong(p, releaseVersion)
	if err.ErrorCode != 0 {
		return models.Song{}, err
	}
	return song, err
}

func (s *songServiceImpl) FetchSongsPerVersionId(sv uint) ([]models.Song, errors.SongError) {
	_songs, err := s.songDal.FetchSongsPerVersionId(sv)

	if err.ErrorCode != 0 {
		return []models.Song{}, err
	}
	return _songs, err
}

func (s *songServiceImpl) AddSong(sn *dataformat.CreateSong, releaseVersion uint) (models.Song, errors.SongError) {

	if releaseVersion == constants.LATEST_RELEASE_KEY {
		_r, err := s.releaseVersionService.GetLatestReleaseVersion()
		if err.ErrorCode != 0 {
			return models.Song{}, errors.SongError(err)
		}
		_sn, _err := s.songDal.SaveSong(sn, _r.ID)
		if _err.ErrorCode != 0 {
			return models.Song{}, _err
		}
		return _sn, _err
	}
	_s, _err := s.songDal.SaveSong(sn, releaseVersion)
	if _err.ErrorCode != 0 {
		return models.Song{}, _err
	}
	return _s, _err
}

func (s *songServiceImpl) DeleteSong(sId uint) (dataformat.DeletedSong, errors.SongError) {
	_s, _err := s.songDal.DeleteSong(sId)
	if _err.ErrorCode != 0 {
		return dataformat.DeletedSong{}, _err
	}
	return _s, _err
}
