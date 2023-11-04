package router

import (
	"github.com/RianIhsan/go-postgres-redis/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine, bookController *controller.BookController) *gin.Engine {
	api := router.Group("/api/v1")
	api.POST("/book", bookController.CreateBook)
	return router
}
