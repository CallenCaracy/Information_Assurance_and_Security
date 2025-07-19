package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var choice int
	reader := bufio.NewReader(os.Stdin)

	for {
		displayOptions()
		fmt.Scan(&choice)
		reader.ReadString('\n')

		switch choice {
		case 1:
			plainText := readLine("Enter text to encrypt: ")

			getSHA1(plainText)

			keyText := readLine("Enter key value (Strictly 32 characters in length): ")
			if len(keyText) != 32 {
				fmt.Println("WARNING!, Key must be exactly 32 characters.")
				continue
			}

			key := []byte(keyText)

			ciphertextHex, nonceHex, err := encryptAESGCM(plainText, key)
			if err != nil {
				fmt.Println("Encryption failed:", err)
				continue
			}

			fmt.Println("Ciphertext (hex):", ciphertextHex)
			fmt.Println("Nonce (hex):", nonceHex)

		case 2:
			ciphertextHex := readLine("Enter ciphertext (hex): ")
			nonceHex := readLine("Enter nonce (hex): ")
			keyText := readLine("Enter key value (Strictly 32 characters in length): ")
			if len(keyText) != 32 {
				fmt.Println("WARNING!, Key must be exactly 32 characters.")
				continue
			}

			key := []byte(keyText)

			plainText, err := decryptAESGCM(ciphertextHex, nonceHex, key)
			if err != nil {
				fmt.Println("Decryption failed:", err)
				continue
			}

			fmt.Println("Decrypted text:", plainText)
			getSHA1(plainText)
		case 3:
			return

		default:
			fmt.Println("Invalid choice.")
			fmt.Println()
		}
	}
}

func displayOptions() {
	fmt.Println("AES Cryptography")
	fmt.Println("1. Encrypt A Text")
	fmt.Println("2. Decrypt A Text")
	fmt.Println("3. Exit")
	fmt.Print("Pick your choice: ")
}

func encryptAESGCM(plaintext string, key []byte) (ciphertextHex, nonceHex string, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", "", err
	}

	ciphertext := aesGCM.Seal(nil, nonce, []byte(plaintext), nil)

	return hex.EncodeToString(ciphertext), hex.EncodeToString(nonce), nil
}

func decryptAESGCM(ciphertextHex, nonceHex string, key []byte) (string, error) {
	ciphertext, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return "", err
	}
	nonce, err := hex.DecodeString(nonceHex)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func getSHA1(text string) {
	data := []byte(text)
	fmt.Printf("SHA1: %x\n", data)
}

func readLine(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
