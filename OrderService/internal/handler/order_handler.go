package handler

import (
	"OrderService/OrderService/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	orderService *service.OrderService
}

func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{orderService}
}

func (oh *OrderHandler) GetOrderByUID(c *gin.Context) {
	ctx := c.Request.Context()
	orderUID := c.Param("order_uid")
	jsonModel, err := oh.orderService.GetOrderByUID(ctx, orderUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json", jsonModel)
}
