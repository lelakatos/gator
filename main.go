package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	fmt.Printf("Home directory is: %s\n", homeDir)
}
