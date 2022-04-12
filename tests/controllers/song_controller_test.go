package controllers

import (
	"acpr_songs_server/controllers/song_controller"
	"acpr_songs_server/core/constants"
	"testing"
)

func TestValidateDeleteSongParam(t *testing.T) {
	releaseVersion, songId := constants.EMPTY_SPACE, constants.EMPTY_SPACE

	msg, err := song_controller.ValidateDeleteSongParam(&releaseVersion, &songId)

	if err == nil {
		t.Error(msg)
	}

	// Test with all values satisfy
	releaseVersion, songId = "version-1", "song1"

	msg, err = song_controller.ValidateDeleteSongParam(&releaseVersion, &songId)

	if err != nil {
		t.Error(msg)
	}

	// Test with on value satisfy
	releaseVersion, songId = "version-1", constants.EMPTY_SPACE

	msg, err = song_controller.ValidateDeleteSongParam(&releaseVersion, &songId)

	if err == nil {
		t.Error(msg)
	}
}
