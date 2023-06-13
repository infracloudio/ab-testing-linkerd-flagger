package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetBook(ctx *gin.Context) {

	log.Printf("In the :: backend :%v :: version : %v", BACKEND, VERSION)
	header := ctx.Request.Header.Get(HEADER)

	log.Printf("Got the header request : %v", header)

	res := BackendAResponse{
		Status:  "success",
		Backend: BACKEND,
		Version: VERSION,
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
