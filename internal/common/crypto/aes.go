package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"os"
)

func getSecretKey() ([]byte, error) {
	key := []byte(os.Getenv("AES_SECRET_KEY"))
	length := len(key)
	if length != 16 && length != 24 && length != 32 {
		return nil, errors.New("invalid AES key length")
	}
	return key, nil
}

func EncryptAES(plainText string) (string, error) {
	key, err := getSecretKey()
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

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func DecryptAES(cipherText string) (string, error) {
	key, err := getSecretKey()
	if err != nil {
		return "", err
	}

	decoded, err := base64.StdEncoding.DecodeString(cipherText)
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

	nonceSize := aesGCM.NonceSize()
	if len(decoded) < nonceSize {
		return "", errors.New("cipher text too short")
	}

	nonce, cipherTextBytes := decoded[:nonceSize], decoded[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherTextBytes, nil)
	return string(plainText), err
}
