package handler

import (
	"net/http"

	"github.com/brunobotter/ecommerce-produto/scheamas"
	"github.com/gin-gonic/gin"
)

func ShowOpenningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	logger.Debugf("id %s", id)
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "Param").Error())
		return
	}
	produto := scheamas.Produto{}
	if err := db.First(&produto, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "produto not found")
		return
	}
	sendSuccess(ctx, "show-opening", produto)
}
