package main

import (

	"github.com/viniihds/api-go/router"
	"github.com/viniihds/api-go/config"
)

var(
	logger *config.Logger
)

func main() {
	//Inialização Router
	router.Initialize()


	logger = config.GetLogger("main")
	//Inialização Configs
	err := config.Init()
	if err != nil{
		logger.Errf("Erro na Inicialização da Config: %v", err)
		return
	}
}