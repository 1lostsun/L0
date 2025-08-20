package router

import (
	"OrderService/ProducerService/internal/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
	gin *gin.Engine
	h   *handler.Handler
}

func NewRouter(gin *gin.Engine, h *handler.Handler) *Router {
	return &Router{
		gin: gin,
		h:   h,
	}
}

func (r *Router) InitRoutes() {
	v1 := r.gin.Group("/v1")
	{
		v1.POST("/order", r.h.CreateOrderAndValidateHandler)
	}
}
