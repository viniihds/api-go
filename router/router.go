package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	//Inicializa o Router com as configs padrões do gin
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
	 // Rodando na porta 8080
}