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

func ShowVagaHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Vaga",
	})
}

func DeleteVagaHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Vaga",
	})
}

func UpdateVagaHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Vaga",
	})
}

func ListVagasHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Vagas",
	})
}
