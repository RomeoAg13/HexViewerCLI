package main

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

const MIN_ASCII = 32

func PrintHexWithASCII(fileName *os.File, offset int) {
	var buffer = make([]byte, 16)
	for {
		n, err := fileName.Read(buffer)
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
