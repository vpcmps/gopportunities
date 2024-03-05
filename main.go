package main

import (
	"github.com/vpcmps/gopportunities/config"
	"github.com/vpcmps/gopportunities/router"
)

var (
	logger = config.NewLogger("")
)

func main() {
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %s", err.Error())
		return
	}
	//Initialize the router
	router.Initialize()
}
