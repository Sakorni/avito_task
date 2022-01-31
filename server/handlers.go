package server

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	core := gin.New()
	group := core.Group("/balance")
	{
		group.POST("/withdraw", h.withdraw)
		group.POST("/deposit", h.deposit)
		group.POST("/transfer", h.transfer)
		group.GET("/", h.getInfo)
	}
	return core
}
