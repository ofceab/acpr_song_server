package errors

import "net/http"

const (
	INTERNAL_ERROR = "an exception on data occured, retry later"

	INVALID_RELEASE_VERSION_ERROR             = "invalid releaseVersion, provide release version that exist"
	NOT_HIGHER_RELEASE_VERSION_PROVIDED_ERROR = "provide a higher version for adding a new version of a song"
	SONG_TO_UPDATE_DONT_EXIST_ERROR           = "invalid song_unique_id. can't add a new version for a song that doesn't exist"
	SONG_TO_DELETE_DOESNT_EXIST_ERROR         = "song with provided id doesnt exist. Be sure of id"

	RELEASE_VERSION__OF_ID_DOESNT_EXIST_ERROR = "releaseVersion with provided id doesnt exist"
	RELEASE_VERSION_DOESNT_EXIST_ERROR        = "no release version yet, retry later"

	ROUTE_NOT_FOUND = "route you tried to reach is not found. Be sure of the path"
)

func GetInternalError() AppError {
	return AppError{Message: INTERNAL_ERROR, ErrorCode: http.StatusInternalServerError}
}
