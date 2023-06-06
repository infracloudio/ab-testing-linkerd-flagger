package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// type ResponseData struct {
// 	Backend string `json:"backend"`
// 	Version string `json:"version"`
// }

func main() {

	router := gin.Default()
	router.GET("/", Forworder)
	router.Run(":8080")
	log.Printf("---------- started forworded server --------------")

}

func Forworder(ctx *gin.Context) {

	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", "http://backend-a:8080", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}


	// Set headers on the request
	req.Header = ctx.Request.Header


	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	// Parse the response body into a JSON object
	var b interface{}
	err = json.Unmarshal(body, &b)
	if err != nil {
		log.Printf("Error %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshell resp body"})
		return
	}
	log.Printf("Response-->%v", b)

	ctx.JSON(resp.StatusCode, b)

}
