package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func ShowVagaHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Vaga",
	})
}
