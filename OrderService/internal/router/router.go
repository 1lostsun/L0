package router

import (
	"OrderService/OrderService/internal/handler"
	"github.com/gin-gonic/gin"
)

type Router struct {
	gin *gin.Engine
	oh  *handler.OrderHandler
}

func NewRouter(gin *gin.Engine, oh *handler.OrderHandler) *Router {
	return &Router{gin: gin, oh: oh}
}

func (r *Router) InitRoutes() {
	v1 := r.gin.Group("/v1")
	{
		v1.GET("order/:order_uid", r.oh.GetOrderByUID)
	}
}
