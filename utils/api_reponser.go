package utils

import (
	"acpr_songs_server/core/errors"

	"github.com/gin-gonic/gin"
)

// Send response
func SendResponse(c *gin.Context, msg errors.AppError) {
	if msg.ErrorCode < 300 && msg.ErrorCode >= 200 {
		c.JSON(msg.ErrorCode, msg.Error())
		return
	}
	c.JSON(msg.ErrorCode, gin.H{"error": msg.Error()})
}
