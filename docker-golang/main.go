package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	log.Println("Application running in environment: ",
		os.Getenv("RUNTIME_SETUP"), " and on port: ", os.Getenv("PORT"))

	var router *gin.Engine
	router = gin.Default()
	router.Static("/", "./static")
	router.Run()

}
