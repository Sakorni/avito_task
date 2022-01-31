package server

import (
	"avito_task/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func parseChangeAmountBody(c *gin.Context) ChangeAmountModel {
	var changeModel ChangeAmountModel
	if err := c.ShouldBindWith(&changeModel, binding.JSON); err != nil {
		sendErrorResponse(c, http.StatusBadRequest,
			fmt.Sprintf("invalid request body: %s", err.Error()))
		return changeModel
	}
	if changeModel.Amount < 1 {
		sendErrorResponse(c, http.StatusBadRequest,
			fmt.Sprint("invalid request body: Amount should be greater than 0"))
	}
	return changeModel
}

func (h *Handler) withdraw(c *gin.Context) {
	changeModel := parseChangeAmountBody(c)
	if c.IsAborted() {
		return
	}
	balance, err := h.service.ChangeAmount(changeModel.UserId, -changeModel.Amount)
	if err != nil {
		if err == errors.NoSuchUser {
			sendErrorResponse(c, http.StatusNotFound, err.Error())
		} else if err == errors.NotEnoughMoney {
			sendErrorResponse(c, http.StatusMethodNotAllowed, err.Error())
		} else {
			sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, &balance)
}

func (h *Handler) deposit(c *gin.Context) {
	changeModel := parseChangeAmountBody(c)
	if c.IsAborted() {
		return
	}
	balance, err := h.service.ChangeAmount(changeModel.UserId, changeModel.Amount)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, balance)
}

func (h *Handler) transfer(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"molodec": "sosi konec transfer"})
}
func (h *Handler) getInfo(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"molodec": "sosi konec get info"})
}
