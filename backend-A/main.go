package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	VERSION = os.Getenv("VERSION")
	BACKEND = os.Getenv("BACKEND_ENV")
	HEADER = os.Getenv("HEADER")
)

func main() {

	fmt.Print("---------------STARTED BACKEND SERVICE----------------\n")
	fmt.Printf("BACKEND :: %v\n", BACKEND)
	fmt.Printf("VERSION :: %v\n", VERSION)
	fmt.Print("------------------------------------------------------\n")
	router := gin.Default()
	router.GET("/", TestApiGetBook)

	router.Run(":8080")

}
