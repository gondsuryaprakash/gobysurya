package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

const (
	key = "my32digitkey12345678901234567890"
	iv  = "my16digitIvKey12"
)

func Encryption(plainText string) (encryptedStr string, err error) {
	var plainTextBlock []byte
	length := len(plainText)
	if length%16 != 0 {
		extendedBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendedBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendedBlock)}, extendedBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}
	copy(plainTextBlock, []byte(plainText))
	blocks, err1 := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("Err in the creating block : ", err)
		return encryptedStr, err1
	}
	chipherText := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(blocks, []byte(iv))
	mode.CryptBlocks(chipherText, plainTextBlock)
	encryptedStr = base64.StdEncoding.EncodeToString(chipherText)
	return

}
