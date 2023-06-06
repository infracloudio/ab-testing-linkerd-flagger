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

	res := BackendBResponse{
		Status:  "success",
		Backend: backend,
		Version: version,
		Header:  header,
		BResponse: BResponse{
			Id:        "123456",
			Title:     "The Great Gatsby",
			Author:    "F. Scott Fitzgerald",
			Year:      1925,
			Genre:     "Fiction",
			Summary:   "The Great Gatsby is a novel by F. Scott Fitzgerald...",
			Publisher: "Scribner",
			Rating: Rating{
				Average:     4.8,
				ToatalVotes: 100,
			},
		},
	}

	ctx.JSON(200, res)

}
