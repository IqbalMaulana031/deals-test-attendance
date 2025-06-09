package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1" // #nosec
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"starter-go-gin/common/constant"
)

const (
	cost = 10
)

// EncryptAESCBC encrypts a string using AES 256 (CBC) with a given key and iv.
func EncryptAESCBC(plaintext, key, iv string) string {
	blockSize := constant.ThirtyTwo
	bKey := []byte(key)
	bIV := []byte(iv)
	bPlaintext := PKCS5Padding([]byte(plaintext), blockSize, len(plaintext))
	block, _ := aes.NewCipher(bKey)
	ciphertext := make([]byte, len(bPlaintext))
	mode := cipher.NewCBCEncrypter(block, bIV)
	mode.CryptBlocks(ciphertext, bPlaintext)
	return hex.EncodeToString(ciphertext)
}

// PKCS5Padding adds padding to a string.
func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// DecryptAESCTR decrypts a string.
func DecryptAESCTR(key []byte, secure string) (decoded string, err error) {
	cipherText, err := base64.StdEncoding.DecodeString(secure)
	if err != nil {
		return
	}

	// Create a new AES cipher with the key and encrypted message
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	iv := make([]byte, aes.BlockSize)

	// Decrypt the message
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), err
}

// EncryptAESCTR encrypts a string.
func EncryptAESCTR(key []byte, message string) (encoded string, err error) {
	// Create byte array from the input string
	plainText := []byte(message)

	// Create a new AES cipher using the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	// Make the cipher text a byte array of the same size as message
	cipherText := make([]byte, len(plainText))

	// iv is null byte array of size BlockSize
	iv := make([]byte, aes.BlockSize)

	// Encrypt the data:
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherText, plainText)

	// Return string encoded in base64
	return base64.StdEncoding.EncodeToString(cipherText), err
}

// Decrypt encrypts a string.
func Decrypt(encryptedString string, keyString string) string {
	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintext)
}

// SHAEncrypt encrypts a string
func SHAEncrypt(plainText string) string {
	sha := sha1.New() // #nosec
	sha.Write([]byte(plainText))
	encrypted := sha.Sum(nil)
	encryptedString := fmt.Sprintf("%x", encrypted)

	return encryptedString
}

// BcryptEncrypt encrypts a string.
func BcryptEncrypt(plainText string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plainText), cost)
	return string(hashed), err
}

// BcryptVerifyHash compares hashed and plain string.
func BcryptVerifyHash(encrypted, plain string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(plain)); err != nil {
		return false
	}
	return true
}
