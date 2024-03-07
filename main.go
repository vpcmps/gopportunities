package main

import (
	"github.com/vpcmps/gopportunities/config"
	"github.com/vpcmps/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}
	//Initialize the router
	router.Initialize()
}
