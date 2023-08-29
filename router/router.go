package router


import "github.com/gin-gonic/gin"

func Initialize() {

// Inicializa o Router
router := gin.Default()

initializeRoutes(router)

// Roda o Servidor
	router.Run(":8080")
}