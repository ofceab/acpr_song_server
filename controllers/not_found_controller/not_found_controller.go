package notfoundcontroller

import (
	"acpr_songs_server/core/errors"
	"acpr_songs_server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoFoundRouteHandler(c *gin.Context) {
	utils.SendResponse(c, errors.AppError{ErrorCode: http.StatusNotFound, Message: errors.ROUTE_NOT_FOUND})
}
