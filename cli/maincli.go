package cli

import (
	"ciphgoer/internals/ciphers"
	"fmt"
)

func checkPanic(e error) {
	if e != nil {
		panic(e)
	}
}

func StartCli(filepath ...string) {
	var (
		operation uint8
		cipher    uint8
		key       string
	)

	fmt.Print("Which operation do you want to do?\n0 -> Decode\n1 -> Encode (Default: 0 -> Decode): ")
	n, err := fmt.Scanln(&operation)
	if n == 0 {
		operation = 0
	} else {
		checkPanic(err)
	}

	fmt.Print("Which cipher should be used?\n0 -> Caesar Cipher\n1 -> Vigenere Cipher (Default: 0 -> Caesar Cipher): ")
	n, err = fmt.Scanln(&cipher)
	if n == 0 {
		cipher = 0
	} else {
		checkPanic(err)
	}

	if cipher == 1 {
		fmt.Print("Please give your encoding/decoding key (Default: 'KEY'): ")
		x, err := fmt.Scanln(&key)
		if x == 0 {
			key = "key"
		} else {
			checkPanic(err)
		}
	}

	// @TODO: fix that. That is absolutely horrible. I'll leave it as it is now just for
	// testing purposes, but I gotta do this right urgently.
	if len(filepath) != 1 {
		str, err := readIOString('\n') // Read everything until newline is found
		checkPanic(err)
		switch cipher {
		case 0:
			str, err = ciphers.Caesar(str, operation)
			checkPanic(err)
		case 1:
			str = ciphers.Vigenere(str, key, operation)
			checkPanic(err)
		default:
			panic("Cipher unknown!")
		}
		fmt.Printf("\nHere's your final string: %s", str)
	} else {
		f := openFile(filepath[0])
		defer closeFile(f)

		scanner := scanFile(f)
		out := createFile()
		defer closeFile(out)

		for scanner.Scan() {
			line := scanner.Text()
			switch cipher {
			case 0:
				line, err = ciphers.Caesar(line, operation)
				checkPanic(err)
			case 1:
				line = ciphers.Vigenere(line, key, operation)
				checkPanic(err)
			default:
				panic("Cipher unknown!")
			}
			_, err := out.WriteString(line + "\n")
			checkPanic(err)
		}
	}
}
