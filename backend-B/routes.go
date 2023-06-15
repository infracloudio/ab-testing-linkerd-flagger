package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetBook(ctx *gin.Context) {

	log.Printf("In the :: backend :%v :: version : %v", BACKEND, VERSION)

	res := BackendBResponse{
		Status:  "success",
		Backend: BACKEND,
		Version: VERSION,
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
