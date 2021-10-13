package main

import (
	"log"
	"main/util"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	_, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

}

func main() {

	log.Println("Application running in environment: ", viper.GetString("RUNTIME_SETUP"),
		" and on port: ", viper.GetString("APP_PORT"))

	var router *gin.Engine = gin.Default()
	router.Static("/", "./static")
	router.Run(viper.GetString("SERVER_ADDRESS") + ":" +
		viper.GetString("APP_PORT"))

}
