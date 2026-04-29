package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	type ApiResp struct {
		Data struct {
			AreaMetadata []struct {
				Name          string `json:"name"`
				LabelLocation struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"label_location"`
			} `json:"area_metadata"`
		} `json:"data"`
	}

	var apiResp ApiResp

	resp, err := http.Get("https://api-open.data.gov.sg/v2/real-time/api/two-hr-forecast")
	if err != nil {
		log.Fatal("Error calling API:", err)
	}
	defer resp.Body.Close()

	byteData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading data:", err)
	}

	err = json.Unmarshal(byteData, &apiResp)
	if err != nil {
		log.Fatal("Error decoding data:", err)
	}

	for i, region := range apiResp.Data.AreaMetadata {
		fmt.Printf("===== Region %d =====\n", i+1)
		fmt.Println("Name:", region.Name)
		fmt.Printf("Location: %.2f, %.2f\n\n", region.LabelLocation.Latitude, region.LabelLocation.Longitude)
	}
}
