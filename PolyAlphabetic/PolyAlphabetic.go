package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var Vigenère [26][26]rune

	fmt.Print("Generating Tabula Recta...\n")
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			Vigenère[i][j] = rune((i+j)%26 + 'A')
		}
	}

	var choice int
	reader := bufio.NewReader(os.Stdin)

	for {
		displayOptions()
		fmt.Scan(&choice)
		reader.ReadString('\n')

		switch choice {
		case 1:
			var secretKey string
			var keyStream []rune
			var cipherText string

			fmt.Print("Enter text to encrypt: ")
			plainText, _ := reader.ReadString('\n')
			plainText = strings.TrimSpace(plainText)

			fmt.Print("Enter key value: ")
			fmt.Scan(&secretKey)
			reader.ReadString('\n')

			plainText = normalizeText(plainText)
			secretKey = normalizeText(secretKey)
			keyStream = makeKeyStream(secretKey, plainText)

			fmt.Println("Plain Text (Normalized):", plainText)
			getSHA1(plainText)

			for i, char := range plainText {
				if char >= 'A' && char <= 'Z' {
					row := char - 'A'
					col := keyStream[i] - 'A'
					cipherText += string(Vigenère[row][col])
				} else {
					cipherText += string(char)
				}
			}

			fmt.Println("Cipher Text:", cipherText)
			fmt.Println()

		case 2:
			var secretKey string
			var keyStream []rune
			var plainText string

			fmt.Print("Enter text to decrypt:")
			cipherText, _ := reader.ReadString('\n')
			cipherText = strings.TrimSpace(cipherText)

			fmt.Print("Enter key value: ")
			fmt.Scan(&secretKey)
			reader.ReadString('\n')

			cipherText = normalizeText(cipherText)
			secretKey = normalizeText(secretKey)
			keyStream = makeKeyStream(secretKey, cipherText)

			for i, char := range cipherText {
				if char >= 'A' && char <= 'Z' {
					c := char - 'A'
					k := keyStream[i] - 'A'
					p := (c - k + 26) % 26
					plainText += string(p + 'A')
				} else {
					plainText += string(char)
				}
			}

			fmt.Println("Plain Text:", plainText)
			getSHA1(plainText)
			fmt.Println()

		case 3:
			return

		default:
			fmt.Println("Invalid choice.")
			fmt.Println()
		}
	}
}

func displayOptions() {
	fmt.Println("PolyAlphabetic Cryptography")
	fmt.Println("1. Encrypt A Text")
	fmt.Println("2. Decrypt A Text")
	fmt.Println("3. Exit")
	fmt.Print("Pick your choice: ")
}

func normalizeText(text string) string {
	text = strings.ToUpper(strings.TrimSpace(text))
	return text
}

func makeKeyStream(key string, text string) []rune {
	var keyStream []rune
	for i, char := range text {
		if char >= 'A' && char <= 'Z' {
			keyStream = append(keyStream, rune(key[i%len(key)]))
		} else {
			keyStream = append(keyStream, char)
		}
	}
	return keyStream
}

func getSHA1(text string) {
	data := []byte(text)
	fmt.Printf("SHA1: %x\n", data)
}
