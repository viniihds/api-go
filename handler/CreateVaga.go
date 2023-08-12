package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateVagaHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Vaga",
	})
}