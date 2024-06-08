package handler

import (
	"fmt"
	"net/http"

	"github.com/brunobotter/ecommerce-produto/scheamas"
	"github.com/gin-gonic/gin"
)

func DeleteOpenningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "Param").Error())
		return
	}
	produto := scheamas.Produto{}
	if err := db.First(&produto, id).Error; err != nil {
		sendError(ctx, http.StatusFound, fmt.Sprintf("produto with Id: %s not found", id))
		return
	}
	if err := db.Delete(&produto).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting produto with id: %s not found", id))
		return
	}
	sendSuccess(ctx, "delete produto", produto)
}
