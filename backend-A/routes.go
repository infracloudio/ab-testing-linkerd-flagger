package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GetBook(ctx *gin.Context) {

	log.Printf("In the :: backend :%v :: version : %v", BACKEND, VERSION)

	res := BackendAResponse{
		Status:  "success",
		Backend: BACKEND,
		Version: VERSION,
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
