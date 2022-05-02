package song_controller

import (
	"acpr_songs_server/core/constants"
	"acpr_songs_server/models"
	"acpr_songs_server/service/song_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type songControllerImpl struct {
	songService song_service.ISongService
}

// Implementing ISongController

func (s *songControllerImpl) FetchSongs(c *gin.Context) {
	songs, err := s.songService.FetchSongs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Send result
	c.JSON(http.StatusOK, songs)
}

func (s *songControllerImpl) FetchSongsPerVersionId(c *gin.Context) {
	reVersion, _err := strconv.ParseUint(c.Param(constants.RELEASE_VERSION_KEY), 10, 32)
	if _err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid releaseVersion"})
		return
	}

	songs, err := s.songService.FetchSongsPerVersionId(uint(reVersion))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Send result
	c.JSON(http.StatusOK, songs)
}

func (s *songControllerImpl) AddSong(c *gin.Context) {
	releaseVersion, _convErr := strconv.ParseUint(c.Param(constants.RELEASE_VERSION_KEY), 10, 32)
	if _convErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid releaseVersion"})
		return
	}

	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result, _err := s.songService.AddSong(&song, uint(releaseVersion))
	if _err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": _err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (s *songControllerImpl) DeleteSong(c *gin.Context) {
	songId, _conErr := strconv.ParseUint(c.Param(constants.SONG_ID_KEY), 10, 32)
	if _conErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid songId"})
		return
	}

	result, _err := s.songService.DeleteSong(uint(songId))
	if _err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": _err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}
