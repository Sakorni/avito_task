package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func sendErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, gin.H{"message": message})
}
