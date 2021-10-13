package main

import (
	"log"
	"main/util"

	"github.com/gin-gonic/gin"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	log.Println("Application running in environment: ",
		config.RuntimeSetup, " and on port: ", config.AppPort)

	var router *gin.Engine
	router = gin.Default()
	router.Static("/", "./static")
	router.Run(config.ServerAddress + ":" + config.AppPort)

}
