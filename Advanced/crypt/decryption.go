package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// Decrypt the encrypted byte data 
func Decription(encryptedText string) ([]byte, error) {

	cipherText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return nil, err
	}

	if len(cipherText)%aes.BlockSize != 0 {
		return nil, errors.New("block size can't be zero")
	}

	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(cipherText, cipherText)

	cipherText = PKCS5UnPadding(cipherText)

	return cipherText, nil
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
