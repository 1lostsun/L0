package router

import "github.com/gin-gonic/gin"

type Router struct {
	gin *gin.Engine
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) InitRoutes() {
	//v1 := r.gin.Group("/v1")
	//{
	//
	//}
}
