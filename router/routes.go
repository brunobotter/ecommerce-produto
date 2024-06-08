package router

import (
	"github.com/brunobotter/ecommerce-produto/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/produto/:id", handler.ShowOpenningHandler)
		v1.GET("/produtos", handler.ListOpenningHandler)
		v1.POST("/produto", handler.CreateOpenningHandler)
		v1.PUT("/produto/:id", handler.UpdateOpenningHandler)
		v1.DELETE("/produto/:id", handler.DeleteOpenningHandler)
	}
}
