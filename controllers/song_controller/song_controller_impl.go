package song_controller

import (
	"acpr_songs_server/core/constants"
	customErros "acpr_songs_server/errors"
	"acpr_songs_server/models"
	"acpr_songs_server/service/song_service"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type songControllerImpl struct {
	songService song_service.ISongService
}

// Implementing ISongController

func (s *songControllerImpl) FetchSongs(c *gin.Context) {
	songs := s.songService.FetchSongs()
	// Send result
	c.JSON(200, songs)
}

func (s *songControllerImpl) FetchSongsPerVersionId(c *gin.Context) {
	songs := s.songService.FetchSongsPerVersionId()
	// Send result
	c.JSON(200, songs)
}

func (s *songControllerImpl) AddSong(c *gin.Context) {
	releaseVersion := c.Param(constants.RELEASE_VERSION_KEY)
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(401, err)
		return
	}

	result := s.songService.AddSong(&song, releaseVersion)
	c.JSON(201, result)
}

func (s *songControllerImpl) DeleteSong(c *gin.Context) {
	releaseVersion, songId := c.Param(constants.RELEASE_VERSION_KEY), c.Param(constants.SONG_ID_KEY)

	fmt.Printf("rel %s, sonId %s", releaseVersion, songId)

	msg, err := ValidateDeleteSongParam(&releaseVersion, &songId)
	if err != nil {
		c.JSON(401, msg)
		return
	}

	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(401, err)
		return
	}

	result := s.songService.AddSong(&song, releaseVersion)
	c.JSON(201, result)
}

func ValidateDeleteSongParam(releaseVersion *string, songId *string) (message string, err error) {
	var _msg []string
	if *releaseVersion == constants.EMPTY_SPACE {
		_msg = append(_msg, "releaseVersion is missing")
	}
	if *songId == constants.EMPTY_SPACE {
		_msg = append(_msg, "songId is missing")
	}

	if len(_msg) != 0 {
		_msgStringify := strings.Join(_msg, ",")
		_error := customErros.InvalidateDeleteRequest{}
		return _msgStringify, &_error
	}
	return "", nil
}
