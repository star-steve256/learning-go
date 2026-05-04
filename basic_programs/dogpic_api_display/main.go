package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/blacktop/go-termimg"
)

type RequestApiResp struct {
	Message string `json:"message"`
	Status string `json:"status"`
}

// NOTE: resolution may be poor because of terminal display
func main() {
	var requestApiResp RequestApiResp
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		log.Fatal("Error API call:", err)
	}
	defer resp.Body.Close()

	byteData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading API:", err)
	}

	err = json.Unmarshal(byteData, &requestApiResp)
	if err != nil {
		log.Fatal("Error unmarshaling json:", err)
	}

	fmt.Println("Result image: ", requestApiResp.Message)
	fmt.Println("Fetching image ...")

	resp, err = http.Get(requestApiResp.Message)
	if err != nil {
		log.Fatal("Error calling dog image API:", err)
	}
	defer resp.Body.Close()

	byteData, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading API:", err)
	}

	err = os.WriteFile("random_dog_img.jpg", byteData, 0644)
	if err != nil {
		log.Fatal("Error saving file:", err)
	}

	err = termimg.PrintFile("random_dog_img.jpg");
	if err != nil {
		log.Fatal("Error displaying image with termimg:", err)
	}

	err = os.Remove("random_dog_img.jpg")
	if err != nil {
		log.Fatal("Error deleting file:", err)
	}
}
