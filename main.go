package main

import (
	"github.com/brunobotter/ecommerce-produto/config"
	"github.com/brunobotter/ecommerce-produto/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initializate error: %v", err)
		return
	}
	//initialize router
	router.Initialize()
}
