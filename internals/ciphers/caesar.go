package ciphers

import "fmt"

// Take a string with a shift number and apply the caesar cipher to it
// @arg str = string that will be encoded/decoded
// @arg sn = shift value that will be applied to the encoding/decoding
func Caesar(str string, op uint8) (string, error) {
	var (
		sn        int8 = 3
		newstring string
	)

	fmt.Print("How many characters do you want to shift? (Default: 3 (a -> d): ")
	_, err := fmt.Scanln(&sn)

	if err != nil {
		return "", err
	}

	if op == 0 {
		sn = -sn
	}

	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			newstring += string(((int8(c) - 'a' + 26 + sn) % 26) + 'a')
		} else if c >= 'A' && c <= 'Z' {
			newstring += string(((int8(c) - 'A' + 26 + sn) % 26) + 'A')
		} else {
			newstring += string(c)
		}
	}

	return newstring, nil
}
