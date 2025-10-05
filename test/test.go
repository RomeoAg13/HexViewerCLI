package test

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

const MIN_ARGS = 2
const MIN_ASCII = 32

func main() {
	var buffer = make([]byte, 16)
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
	for {
		n, err := openFile.Read(buffer)
		t := ""
		// we loop over the number of bytes read
		for i := 0; i < n; i++ {
			// we're looking if the i-th byte is a printable character (> 33 and < 127)
			if buffer[i] >= MIN_ASCII && buffer[i] < unicode.MaxASCII {
				t += string(buffer[i])
			} else {
				t += "."
			}
		}
		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF REACHED")
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%08x %x |%s|\n", offset, buffer[:n], t)
		offset += n
	}
}
