package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListVagasHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Vagas",
	})
}
