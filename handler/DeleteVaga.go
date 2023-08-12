package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteVagaHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Vaga",
	})
}