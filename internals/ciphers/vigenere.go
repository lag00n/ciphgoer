package ciphers

import (
	"strings"
)

func Vigenere(text, key string, op uint8) string {
	text = strings.ToUpper(text)
	key = strings.ToUpper(key)

	// Key and text lenght
	textL := len(text)
	keyL := len(key)

	result := make([]byte, textL)

	keyIndex := 0

	for i := 0; i < textL; i++ {
		char := text[i]

		if char >= 'A' && char <= 'Z' {
			shift := int(key[keyIndex] - 'A')

			// If encrypting...
			if op == 1 {
				char = (char-'A'+byte(shift))%26 + 'A'
			} else {
				char = (char-'A'+byte(26-shift))%26 + 'A'
			}

			keyIndex = (keyIndex + 1) % keyL
		}

		result[i] = char
	}

	return string(result)
}
