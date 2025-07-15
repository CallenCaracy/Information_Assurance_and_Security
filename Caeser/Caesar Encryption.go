package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const eName = "Caesar_Output_Encryption.txt"
	const dName = "Caesar_Output_Dencryption.txt"

	var choice int
	reader := bufio.NewReader(os.Stdin)

	for {
		displayOptions()
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

			getSHA1(plainText)

			fileCreate(eName)
			fileWrite(cipherText, eName)

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

			fmt.Println("Plain Text:", plainText)

			getSHA1(plainText)

			fileCreate("Caesar_Output_Dencryption.txt")
			fileWrite(plainText, dName)

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
	fmt.Println("Caesar Cryptography")
	fmt.Println("1. Encrypt A Text")
	fmt.Println("2. Decrypt A Text")
	fmt.Println("3. Exit")
	fmt.Print("Pick your choice: ")
}

func fileCreate(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
}

func fileWrite(cipherText string, fileName string) {
	data := []byte(cipherText)
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("File writing error", err)
		return
	}
	fmt.Println("Data successfully written to file")
}

func getSHA1(text string) {
	data := []byte(text)
	fmt.Printf("SHA1: %x\n", data)
}
