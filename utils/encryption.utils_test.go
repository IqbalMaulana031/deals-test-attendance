package utils

import (
	"testing"
)

func TestEncryptAESCBC(t *testing.T) {
	plaintext := "Hello, World!"
	key := "UkXp2r5u8x/A?D(G+KbPeShVmYq3t6v9"
	iv := "1234567890123458"

	encrypted := EncryptAESCBC(plaintext, key, iv)
	if encrypted == plaintext {
		t.Errorf("EncryptAESCBC failed, expected encrypted string, got %s", encrypted)
	}
}

func TestDecryptAESCTR(t *testing.T) {
	key := []byte("-JaNdRgUkXp2s5v8")
	message := "8JOE0xvXxtgpyK4s4xOA"

	encrypted, err := EncryptAESCTR(key, message)
	if err != nil {
		t.Errorf("EncryptAESCTR failed: %v", err)
	}

	print(encrypted)

	decrypted, err := DecryptAESCTR(key, encrypted)
	if err != nil {
		t.Errorf("DecryptAESCTR failed: %v", err)
	}

	if decrypted != message {
		t.Errorf("DecryptAESCTR failed, expected %s, got %s", message, decrypted)
	}
}

func TestEncryptAESCTR(t *testing.T) {
	key := []byte("UkXp2r5u8x/A?D(G+KbPeShVmYq3t6v9")
	message := "Hello, World!"

	encrypted, err := EncryptAESCTR(key, message)
	if err != nil {
		t.Errorf("EncryptAESCTR failed: %v", err)
	}

	if encrypted == message {
		t.Errorf("EncryptAESCTR failed, expected encrypted string, got %s", encrypted)
	}
}

func TestSHAEncrypt(t *testing.T) {
	plaintext := "Hello, World!"
	expectedLength := 40 // SHA1 hash length in hex

	encrypted := SHAEncrypt(plaintext)
	if len(encrypted) != expectedLength {
		t.Errorf("SHAEncrypt failed, expected length %d, got %d", expectedLength, len(encrypted))
	}
}

func TestBcryptEncrypt(t *testing.T) {
	plaintext := "Hello, World!"

	encrypted, err := BcryptEncrypt(plaintext)
	if err != nil {
		t.Errorf("BcryptEncrypt failed: %v", err)
	}

	if encrypted == plaintext {
		t.Errorf("BcryptEncrypt failed, expected encrypted string, got %s", encrypted)
	}
}

func TestBcryptVerifyHash(t *testing.T) {
	plaintext := "Hello, World!"

	encrypted, err := BcryptEncrypt(plaintext)
	if err != nil {
		t.Errorf("BcryptEncrypt failed: %v", err)
	}

	if !BcryptVerifyHash(encrypted, plaintext) {
		t.Errorf("BcryptVerifyHash failed, expected true, got false")
	}
}
