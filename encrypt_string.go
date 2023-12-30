package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

// DecryptedString is a decrypted string
type DecryptedString string

// NewDecryptedString creates a decrypted string
func NewDecryptedString(value string) DecryptedString {
	return DecryptedString(value)
}

// generateIV generates an Initialization Vector of AES block size
func generateIV() []byte {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		panic(err)
	}

	return iv
}

// pkcs7Padding adds padding to the data using PKCS7 method
func pkcs7Padding(data []byte, blockSize int) []byte {
	padSize := blockSize - len(data)%blockSize
	padBytes := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(data, padBytes...)
}

// pkcs7UnPadding removes padding from the data using PKCS7 method
func pkcs7UnPadding(data []byte) []byte {
	dataLength := len(data)
	padLength := int(data[dataLength-1])
	return data[:dataLength-padLength]
}

// Encrypt converts a DecryptedString to an EncryptedBytes
func (ds DecryptedString) Encrypt() (*EncryptedBytes, error) {
	encryptKey := getConfig().EncryptKey
	encryptKeyBytes := []byte(encryptKey)
	if len(encryptKeyBytes) != 32 {
		return nil, fmt.Errorf("invalid encrypt key length")
	}

	iv := generateIV()
	block, err := aes.NewCipher(encryptKeyBytes)
	if err != nil {
		return nil, err
	}

	paddedBytes := pkcs7Padding([]byte(ds), aes.BlockSize)
	encrypted := make([]byte, len(paddedBytes))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(encrypted, paddedBytes)

	es := EncryptedBytes(append(encrypted, iv...))
	return &es, nil
}

// EncryptedBytes is an encrypted bytes
type EncryptedBytes []byte

// Decrypt converts an EncryptedBytes to a DecryptedString
func (es EncryptedBytes) Decrypt() (*DecryptedString, error) {
	encryptKey := getConfig().EncryptKey
	encryptKeyBytes := []byte(encryptKey)
	if len(encryptKeyBytes) != 32 {
		return nil, fmt.Errorf("invalid encrypt key length")
	}

	block, err := aes.NewCipher([]byte(encryptKey))
	if err != nil {
		panic(err)
	}

	encrypted, iv := es[:len(es)-aes.BlockSize], es[len(es)-aes.BlockSize:]

	decrypted := make([]byte, len([]byte(encrypted)))
	cbcDecrypter := cipher.NewCBCDecrypter(block, []byte(iv))
	cbcDecrypter.CryptBlocks(decrypted, []byte(encrypted))

	unpadedDecrypted := pkcs7UnPadding(decrypted)
	if unpadedDecrypted == nil {
		return nil, fmt.Errorf("invalid padding size")
	}

	ds := DecryptedString(string(unpadedDecrypted))
	return &ds, nil
}
