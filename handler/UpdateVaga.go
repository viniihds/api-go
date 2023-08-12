package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateVagaHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Vaga",
	})
}