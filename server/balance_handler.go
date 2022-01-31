package server

import (
	"avito_task/errors"
	"avito_task/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
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
	var model TransactionModel
	if err := c.ShouldBindWith(&model, binding.JSON); err != nil {
		sendErrorResponse(c, http.StatusBadRequest,
			fmt.Sprintf("invalid request body: %s", err.Error()))
		return
	}
	if model.Amount < 1 {
		sendErrorResponse(c, http.StatusBadRequest,
			fmt.Sprint("invalid request body: Amount should be greater than 0"))
		return
	}
	err := h.service.Transfer(&models.Transaction{
		Amount: int(model.Amount),
		To:     model.ToId,
		From:   model.FromId,
	})
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
	c.Writer.WriteHeader(http.StatusOK)
}

func getUid(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, err
	}
	if id < 0 {
		return 0, fmt.Errorf("")
	}
	return uint(id), nil
}

func (h *Handler) getInfo(c *gin.Context) {
	id, err := getUid(c)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid id in params")
		return
	}
	balance, err := h.service.GetBalance(id)
	if err != nil {
		if err == errors.NoSuchUser {
			sendErrorResponse(c, http.StatusNotFound, err.Error())
		} else {
			sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	c.JSON(http.StatusOK, &balance)
}
