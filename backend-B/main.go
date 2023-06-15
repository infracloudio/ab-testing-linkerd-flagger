package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	VERSION = os.Getenv("VERSION")
	BACKEND = os.Getenv("BACKEND_ENV")
)

func main() {

	log.Print("---------------STARTED BACKEND SERVICE----------------\n")
	log.Printf("BACKEND :: %v\n", BACKEND)
	log.Printf("VERSION :: %v\n", VERSION)
	log.Print("------------------------------------------------------\n")

	router := gin.Default()
	router.GET("/", GetBook)
	router.Run(":8080")

}
