package server

import (
	"avito_task/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.BalanceService
}

func NewHandler(service *service.BalanceService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	core := gin.New()
	group := core.Group("/balance")
	{
		group.POST("/withdraw", h.withdraw)
		group.POST("/deposit", h.deposit)
		group.POST("/transfer", h.transfer)
		group.GET("/:id", h.getInfo)
	}
	return core
}
