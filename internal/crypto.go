package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

// Функция для генерации ключа на основе пароля (с помощью SHA-256)
func generateKey(password string) []byte {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hash.Sum(nil)
}

// Функция для шифрования сообщения
func Encrypt(message, password string) (string, error) {
	key := generateKey(password)

	// Генерация случайного IV для режима GCM
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Используем GCM для шифрования
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(message), nil)

	// Кодируем зашифрованное сообщение в base64 для удобства хранения
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Функция для дешифрования сообщения
func Decrypt(ciphertextBase64, password string) (string, error) {
	key := generateKey(password)

	// Декодируем из base64
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return "", err
	}

	// Генерация ключа и создание объекта GCM для дешифрования
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Извлекаем nonce из первого блока данных
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Дешифруем сообщение
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
