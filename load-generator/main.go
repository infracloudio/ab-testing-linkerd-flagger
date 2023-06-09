package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	log.Print("--------------------- LOAD GENERATOR STARTED --------------------")

	//generate load for CALL_DURATION mins
	cd, err1 := strconv.Atoi(os.Getenv("CALL_DURATION"))
	//after generating load for CALL_DURATION it sleeps for SLEEP_DURATION mins
	sd, err2 := strconv.Atoi(os.Getenv("SLEEP_DURATION"))

	if err1 != nil || err2 != nil {
		log.Panic("Error while convering call or sleep duration env variable to int")
	}

	canaryHeaderKey := os.Getenv("NEW_VERSION_HEADER_KEY")
	canaryHeaderVal := os.Getenv("NEW_VERSION_HEADER_VAL")
	endpoint := os.Getenv("ENDPOINT")

	log.Printf("CALL DURATION :: %v", cd)
	log.Printf("SLEEP DURATION :: %v", sd)
	log.Printf("NEW_VERSION_HEADER_KEY :: %v", canaryHeaderKey)
	log.Printf("NEW_VERSION_HEADER_VAL :: %v", canaryHeaderVal)
	log.Printf("ENDPOINT :: %v", endpoint)

	log.Print("-----------------------------------------------------------------------")

	//load-generator with header when we use header based A/B testing
	if canaryHeaderKey != "" && canaryHeaderVal != "" {
		header := map[string]string{canaryHeaderKey: canaryHeaderVal}
		go callBackend(endpoint, header, cd, sd)

	}
	//load-generator
	callBackend(endpoint, map[string]string{}, cd, sd)

}

func callBackend(endpoint string, header map[string]string, callDurationENV int, sleepDurationENV int) {

	callDuration := time.Duration(callDurationENV) * time.Minute
	sleepDuration := time.Duration(sleepDurationENV) * time.Minute

	for {
		// Perform API calls for the specified duration
		log.Printf("load generator start for %v minutes for endpoint %v with header %v", callDurationENV, endpoint, header)
		startTime := time.Now()
		endTime := startTime.Add(callDuration)

		for time.Now().Before(endTime) {
			// Make the HTTP request
			DoGetHttp(endpoint, header)
			time.Sleep(1 * time.Second)
		}

		// Sleep for the specified duration
		log.Printf("Sleeping load generator for %v minutes", sleepDurationENV)
		time.Sleep(sleepDuration)
	}
}

func DoGetHttp(endpoint string, newVersionHeader map[string]string) {
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers on the request
	for key, val := range newVersionHeader {
		req.Header.Set(key, val)
	}

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
		return
	}

	// Parse the response body into a JSON object
	var b interface{}
	err = json.Unmarshal(body, &b)
	if err != nil {
		log.Printf("Error %v\n", err)
		return
	}
	log.Printf("Response------->%v", b)
}
