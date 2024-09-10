package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Running a drone plugin")
	// Read Env
	input := os.Getenv("PLUGIN_INPUT")
	if input == "" {
		fmt.Println ("No input provided")
	}else {
		fmt.Printf("Input Received %s\n",input)
	}
	fmt.Println("plugin execution completed")
}

