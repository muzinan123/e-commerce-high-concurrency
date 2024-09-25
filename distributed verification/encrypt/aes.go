package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// Advanced Encryption Standard (AES)

// 16, 24, 32 byte string corresponds to AES-128, AES-192, AES-256 encryption methods respectively
// Key must not be leaked
var PwdKey = []byte("DIS**#KKKDJJSKDI")

// PKCS7 padding mode
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	// Repeat() function duplicates the slice []byte{byte(padding)} 'padding' times, then merges into a new byte slice and returns
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// Reverse operation of padding, remove padding string
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	// Get data length
	length := len(origData)
	if length == 0 {
		return nil, errors.New("Encrypted string error!")
	} else {
		// Get padding string length
		unpadding := int(origData[length-1])
		// Slice the data, remove padding bytes, and return plaintext
		return origData[:(length - unpadding)], nil
	}
}

// Implement encryption
func AesEcrypt(origData []byte, key []byte) ([]byte, error) {
	// Create encryption algorithm instance
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// Get block size
	blockSize := block.BlockSize()
	// Pad the data to satisfy length requirement
	origData = PKCS7Padding(origData, blockSize)
	// Use CBC encryption mode in AES encryption method
	blocMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// Execute encryption
	blocMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// Implement decryption
func AesDeCrypt(cypted []byte, key []byte) ([]byte, error) {
	// Create encryption algorithm instance
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// Get block size
	blockSize := block.BlockSize()
	// Create encryption client instance
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cypted))
	// This function can also be used for decryption
	blockMode.CryptBlocks(origData, cypted)
	// Remove padding string
	origData, err = PKCS7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, err
}

// Encrypt to base64
func EnPwdCode(pwd []byte) (string, error) {
	result, err := AesEcrypt(pwd, PwdKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), err
}

// Decrypt
func DePwdCode(pwd string) ([]byte, error) {
	// Decode base64 string
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return nil, err
	}
	// Execute AES decryption
	return AesDeCrypt(pwdByte, PwdKey)
}
