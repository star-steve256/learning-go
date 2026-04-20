package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("===== Catfact API display =====")

	var responseBody struct {
		Fact string `json:"fact"`
	}

	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		log.Fatal("ERROR at calling API:", err)
	}
	defer resp.Body.Close()

	byteData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ERROR at reading IO:", err)
	}


	err = json.Unmarshal(byteData, &responseBody)
	if err != nil {
		log.Fatal("ERROR at json decoding:", err)
	}

	fmt.Println("Fact:", responseBody.Fact)
}

