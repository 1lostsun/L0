package handler

import (
	"ProducerService/internal/model"
	"ProducerService/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Handler struct {
	uc *usecase.Usecase
}

func NewHandler(uc *usecase.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) CreateOrder(c *gin.Context) {
	ctx := c.Request.Context()
	var req *model.OrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	val := validator.New()

	if err := val.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.SendKafkaMsg(ctx, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
