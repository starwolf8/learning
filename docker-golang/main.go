package main

import (
	"main/logging"
	"main/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var config *util.Config

func init() {
	logging.SetUpLogging()

	myconfig, err := util.LoadConfig(".")
	if err != nil {
		logging.AppLog.WriteLogError("cannot load config:", map[string]interface{}{"source": config, "err": err})
	}

	config = &myconfig

	util.InitializeRedis()

}

func main() {

	logging.AppLog.WriteLogInfo("Application running in environment: ",
		map[string]interface{}{"runtime_setup": viper.GetString("RUNTIME_SETUP"),
			"app_port": viper.GetString("APP_PORT")})

	var router *gin.Engine = gin.Default()
	// router.Static("/", "./static")
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	router.Run(viper.GetString("SERVER_ADDRESS") + ":" +
		viper.GetString("APP_PORT"))

}
