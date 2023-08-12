package router

import (

	"github.com/gin-gonic/gin"
	"github.com/viniihds/api-go/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/vaga", handler.ShowVagaHandler)
		v1.POST("/vaga", handler.CreateVagaHandler )
		v1.DELETE("/vaga", handler.DeleteVagaHandler)
		v1.PUT("/vaga", handler.UpdateVagaHandler)
		v1.GET("/vagas", handler.ListVagasHandler)
	}
}