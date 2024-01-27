package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readIOString(delimiter byte) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please give me the text that you want to encode/decode: ")

	str, err := reader.ReadString(delimiter)
	if err != nil {
		return "", err
	}
	return str, nil
}

func openFile(filepath string) *os.File {
	f, err := os.Open(filepath)

	if err != nil {
		log.Fatal("Error opening your file: ", err)
	}
	return f
}

func scanFile(f *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(f)
	return scanner
}

// This will just create a file for now.
// Gotta find a way to fill this file with the encoded/decoded content.
func createFile() *os.File {
	fo, err := os.Create("output.txt")

	if err != nil {
		log.Fatal("Error creating output file: ", err)
	}
	return fo
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Fatal("Error closing file: ", err)
	}
}
