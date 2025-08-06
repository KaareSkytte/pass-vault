package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/argon2"
)

func GenerateEncryptionKey(password string, salt []byte) []byte {
	key := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	return key
}

func Encrypt(entries, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("couldn't create cipher: %v", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("couldn't create GCM: %v", err)
	}

	nonce := make([]byte, aesgcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, fmt.Errorf("couldn't generate nonce: %v", err)
	}

	encryptedData := aesgcm.Seal(nonce, nonce, entries, nil)

	return encryptedData, nil
}

func Decrypt(data, key []byte) ([]byte, error) {
	if len(data) < 12 {
		return nil, fmt.Errorf("data too short: need at least 12 bytes, got %d", len(data))
	}
	nonce := data[:12]
	encrypted_data := data[12:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("couldn't create cipher: %v", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("couldn't create GCM: %v", err)
	}

	decryptedData, err := aesgcm.Open(nil, nonce, encrypted_data, nil)
	if err != nil {
		return nil, fmt.Errorf("couldn't decrypt data: %v", err)
	}

	return decryptedData, nil
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, fmt.Errorf("couldn't generate salt: %v", err)
	}
	return salt, nil
}
