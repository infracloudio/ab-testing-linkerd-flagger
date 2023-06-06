package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func TestApiGetBook(ctx *gin.Context) {

	backend := os.Getenv("BACKEND_ENV")
	version := os.Getenv("VERSION")
	log.Printf("In the backend %v", backend)
	header := ctx.Request.Header.Get("x-backend")

	log.Printf("Got %v in the header", backend)

	res := BackendAResponse{
		Status:  "success",
		Backend: backend,
		Version: version,
		Header:  header,
		AResponse: AResponse{
			Id:     "123456",
			Title:  "The Great Gatsby",
			Author: "F. Scott Fitzgerald",
			Year:   1925,
			Genre:  "Fiction",
			Rating: 4.5,
		},
	}

	ctx.JSON(200, res)

}
