package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var choice int
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Encrypt A Text")
		fmt.Println("2. Decrypt A Text")
		fmt.Println("3. Exit")
		fmt.Print("Pick your choice: ")
		fmt.Scan(&choice)
		reader.ReadString('\n')

		switch choice {
		case 1:
			var shifting int
			var cipherText string

			fmt.Print("Enter text to encrypt: ")
			plainText, _ := reader.ReadString('\n')
			plainText = strings.TrimSpace(plainText)

			fmt.Print("Enter shifting value: ")
			fmt.Scan(&shifting)
			reader.ReadString('\n')

			for _, char := range plainText {
				if char >= 'a' && char <= 'z' {
					cipherText += string((char-'a'+rune(shifting))%26 + 'a')
				} else if char >= 'A' && char <= 'Z' {
					cipherText += string((char-'A'+rune(shifting))%26 + 'A')
				} else {
					cipherText += string(char)
				}
			}

			fmt.Println("Cipher Text:", cipherText)
			fmt.Println()

		case 2:
			var shifting int
			var plainText string

			fmt.Print("Enter text to decrypt:")
			cipherText, _ := reader.ReadString('\n')
			cipherText = strings.TrimSpace(cipherText)

			fmt.Print("Enter shifting value: ")
			fmt.Scan(&shifting)
			reader.ReadString('\n')

			for _, char := range cipherText {
				if char >= 'a' && char <= 'z' {
					plainText += string((char-'a'-rune(shifting)+26)%26 + 'a')
				} else if char >= 'A' && char <= 'Z' {
					plainText += string((char-'A'-rune(shifting)+26)%26 + 'A')
				} else {
					plainText += string(char)
				}
			}

			fmt.Println("Cipher Text:", plainText)
			fmt.Println()

		case 3:
			return

		default:
			fmt.Println("Invalid choice.")
			fmt.Println()
		}
	}
}
