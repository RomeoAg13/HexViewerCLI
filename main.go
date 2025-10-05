package main

import (
	"fmt"
	"os"
)

const MIN_ARGS = 2

func main() {
	var bytes = 16
	fmt.Println("Enter the number of bytes to read at a time (default is 16):")
	fmt.Scan(&bytes)

	args := os.Args[1]

	if len(os.Args) < MIN_ARGS {
		fmt.Println("Please provide a file name")
		os.Exit(1)
	}

	openFile, err := os.Open(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer openFile.Close()
	// an offset is the position of the byte in the file
	// - it's incremented by the number of bytes read
	// - it's printed in hexadecimal format
	offset := 0
	PrintHexWithASCII(openFile, offset)
}
