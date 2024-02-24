package cookies

import (
	"crypto/hmac"
	"crypto/sha256"
	"net/http"
)

// Write a signature HMAC in a cookie value
func writeSigned(
	w http.ResponseWriter,
	cookie http.Cookie,
	secretKey []byte,
) error {
	signature := createHmacSignature(cookie.Name, cookie.Value, secretKey)
	cookie.Value = string(signature) + cookie.Value
	return write(w, cookie)
}

// Read a signature HMAC in a cookie value
func readSigned(r *http.Request, name string, secretKey []byte) (string, error) {
	signedValue, err := read(r, name)
	if err != nil {
		return "", err
	}

	if len(signedValue) < sha256.Size {
		return "", ErrInvalidValue
	}

	signature := signedValue[:sha256.Size]
	value := signedValue[sha256.Size:]

	expectedSignature := createHmacSignature(name, value, secretKey)

	if !hmac.Equal([]byte(signature), expectedSignature) {
		return "", ErrInvalidValue
	}

	return value, nil
}

func createHmacSignature(name, value string, secretKey []byte) []byte {
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(name))
	mac.Write([]byte(value))
	signature := mac.Sum(nil)

	return signature
}
