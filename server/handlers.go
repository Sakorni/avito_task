package server

import (
	"avito_task/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.BalanceService
	apiKey  string
}

func NewHandler(service *service.BalanceService, apiKey string) *Handler {
	return &Handler{service: service, apiKey: apiKey}
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
