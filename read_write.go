package main

import (
	"fmt"
	"os"
)

func main() {
	// name of the file we will write to/read
	filename := "sample.txt"
	// this is the text we want to save into the file
	contentToWrite := "Hello from GO!\nThis is a sample text file."

	// write to the file
	// os.WriteFile creates the file if it does not exist
	// it will overwrite the file if it already exists
	err := os.WriteFile(filename, []byte(contentToWrite), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File written successfully!")

	// read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// print what we read from the file
	fmt.Println("Content of the file:")
	fmt.Println(string(data))
}
