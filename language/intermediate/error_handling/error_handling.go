package errorhandling

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func Init() {

	// error handling is unique in Go
	// treated as values instead of Object or Class
	_, err := http.Get("Non-existent-api...")

	// verbose error handling with nil check
	if err != nil {
		fmt.Println("Error calling API:", err)
	}

	// errors can be created
	var errOutOfIdeas = errors.New("No more ideas.")
	fmt.Println("Error:", errOutOfIdeas.Error())
	
	// Errorf supports variable insert in error msg
	err = fmt.Errorf("Just a simple secret: %s", "helloworld")
	fmt.Println("Error:", err.Error())


	// usage of openFile
	file, err := openFile("error_handling/file.txt")
	if err != nil {
		fmt.Println("Error opening file.txt: ", err)
	}
	defer file.Close()

	file2, err := openFile("gibberish.txt")
	if err != nil {
		fmt.Println("Error opening gibberish.txt:", err)
	}
	defer file2.Close()
}

// what if a function can give errors?
// return error as a parameter
func openFile(fileName string) (file *os.File, err error)  {
	file, err = os.Open(fileName)
	return
}