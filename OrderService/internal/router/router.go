package router

import (
	"OrderService/internal/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
	gin *gin.Engine
	h   *handler.Handler
}

func NewRouter(gin *gin.Engine, oh *handler.Handler) *Router {
	return &Router{gin: gin, h: oh}
}

func (r *Router) InitRoutes() {
	v1 := r.gin.Group("/v1")
	{
		v1.GET("order/:order_uid", r.h.GetOrderByUID)
	}
}
