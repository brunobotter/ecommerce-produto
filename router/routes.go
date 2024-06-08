package router

import (
	"github.com/brunobotter/ecommerce-produto/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/produto/:id", handler.ShowProdutoHandler)
		v1.GET("/produtos", handler.ListProdutoHandler)
		v1.POST("/produto", handler.CreateProdutoHandler)
		v1.PUT("/produto/:id", handler.UpdateProdutoHandler)
		v1.PATCH("/produto/venda", handler.VendaProdutoHandler)
		v1.DELETE("/produto/:id", handler.DeleteProdutoHandler)
	}
}
