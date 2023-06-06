package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/",TestApiGetBook )
	router.Run(":8080")

}
