package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) withdraw(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"molodec": "sosi konec withdraw"})
}

func (h *Handler) deposit(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"molodec": "sosi konec deposit"})
}

func (h *Handler) transfer(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"molodec": "sosi konec transfer"})
}
func (h *Handler) getInfo(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"molodec": "sosi konec get info"})
}
