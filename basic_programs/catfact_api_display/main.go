package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// DISCLAIMER
/* There may be a better way to fetch APIs with concurrency,
   but for now, sequential call and wait will be used. */

func main() {
	fmt.Println("===== Catfact API display =====")
	var optionInput int

	for {
		fmt.Print(`Enter option:
1. Call API once
2. Call API multiple times.
0. Exit
>> `)

		fmt.Scan(&optionInput)

		if optionInput == 1 {
			requestApi(1)
		} else if optionInput == 2 {
			var numOfApiCalls int
			fmt.Print("Number of API calls >> ")
			fmt.Scan(&numOfApiCalls)

			for i := 0; i < numOfApiCalls; i++ {
				requestApi(i + 1)
			}
		} else if optionInput == 0 {
			break
		}
		fmt.Println("===========================")
	}

	fmt.Println("Program terminated.")
}

func requestApi(index int) {
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

	fmt.Printf("#%d Fact: %s\n", index, responseBody.Fact)
}
