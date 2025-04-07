package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

var SecretKey []byte

func init() {
	secretKeyHex := os.Getenv("ENCRYPTION_SECRET_KEY")
	if secretKeyHex == "" {
		log.Fatal("ENCRYPTION_SECRET_KEY environment variable not set.")
	}
	var err error
	SecretKey, err = hex.DecodeString(secretKeyHex)
	if err != nil {
		log.Fatalf("Failed to decode secret key from hex: %v", err)
	}
	if len(SecretKey) != 32 { // AES-256 requires a 32-byte key
		log.Fatalf("Secret key must be 32 bytes (256 bits). Current length: %d", len(SecretKey))
	}
}

// EncryptPassword encrypts the given password using AES-GCM.
// It returns the ciphertext as a hexadecimal string.
func EncryptPassword(plainTextPassword string) (string, error) {
	block, err := aes.NewCipher(SecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plainTextPassword), nil)

	// Return the nonce prepended to the ciphertext for decryption later
	return hex.EncodeToString(ciphertext), nil
}

// DecryptPassword decrypts the given ciphertext (hexadecimal string) using AES-GCM.
// It returns the original plaintext password.
func DecryptPassword(encryptedPasswordHex string) (string, error) {
	ciphertext, err := hex.DecodeString(encryptedPasswordHex)
	if err != nil {
		return "", fmt.Errorf("failed to decode ciphertext from hex: %w", err)
	}

	block, err := aes.NewCipher(SecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to create AES cipher: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("invalid ciphertext size")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt: %w", err)
	}

	return string(plaintext), nil
}
