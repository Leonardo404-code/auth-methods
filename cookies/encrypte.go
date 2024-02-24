package cookies

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func writeEncrypted(w http.ResponseWriter, cookie http.Cookie, secreteKey []byte) error {
	block, err := aes.NewCipher(secreteKey)
	if err != nil {
		return err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	plaintext := fmt.Sprintf("%s:%s", cookie.Name, cookie.Value)

	encryptedValue := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	cookie.Value = string(encryptedValue)

	return write(w, cookie)
}

func readEncrypted(r *http.Request, name string, secretKey []byte) (string, error) {
	encryptedValue, err := read(r, name)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()

	if len(encryptedValue) < nonceSize {
		return "", ErrInvalidValue
	}

	nonce := encryptedValue[:nonceSize]
	cipherText := encryptedValue[nonceSize:]

	plainText, err := aesGCM.Open(nil, []byte(nonce), []byte(cipherText), nil)
	if err != nil {
		return "", err
	}

	expectedName, value, ok := strings.Cut(string(plainText), ":")
	if !ok {
		return "", ErrInvalidValue
	}

	if expectedName != name {
		return "", ErrInvalidValue
	}

	return value, nil
}
