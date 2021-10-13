package main

import (
	"log"
	"main/util"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	setUpLogging()

	_, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	util.InitializeRedis()

}

func main() {

	log.Println("Application running in environment: ", viper.GetString("RUNTIME_SETUP"),
		" and on port: ", viper.GetString("APP_PORT"))

	var router *gin.Engine = gin.Default()
	router.Static("/", "./static")
	router.Run(viper.GetString("SERVER_ADDRESS") + ":" +
		viper.GetString("APP_PORT"))

}

func setUpLogging() {
	file, err := os.OpenFile(
		"logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}
