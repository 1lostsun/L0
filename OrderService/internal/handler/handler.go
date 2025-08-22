package handler

import (
	"OrderService/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	s *service.Service
}

func NewOrderHandler(orderService *service.Service) *Handler {
	return &Handler{orderService}
}

func (oh *Handler) GetOrderByUID(c *gin.Context) {
	ctx := c.Request.Context()
	orderUID := c.Param("order_uid")
	orderModel, err := oh.s.GetOrderByUID(ctx, orderUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orderModel})
}
