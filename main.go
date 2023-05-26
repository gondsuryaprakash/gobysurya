package main

import (
	"fmt"

	"gotutorial.com/Advanced/crypt"
)

func main() {
	input := "Hello Surya"
	encryptedData, err := crypt.Encryption(input)
	if err != nil {
		fmt.Println(err)
	}
	// Now decryption of the encoded value
	fmt.Println("Encrypted Data : ", encryptedData)
	decryptByte, err := crypt.Decription(encryptedData)
	if err != nil {
		fmt.Println("Err in Decription ")
	}
	fmt.Println("Decrypted Data :", string(decryptByte))
}
