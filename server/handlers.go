package server

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() gin.Engine {
	core := gin.New()
	group := core.Group("/balance")
	{
		group.POST("/withdraw")
		group.POST("/deposit")
		group.POST("/transfer")
		group.GET("/")
	}
}
