package jwt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

func CreateToken(w http.ResponseWriter, r *http.Request) {
	user := User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("error in decode body: %v", err), http.StatusBadRequest)
		return
	}

	token, err := generateJWTAcessToken(user.ID, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
}

func GetTokenData(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		http.Error(w, "invalid Authorization header format", http.StatusUnauthorized)
		return
	}

	token := parts[1]

	tokenParsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return signedKey, nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var userID, userEmail interface{}

	if claims, ok := tokenParsed.Claims.(jwt.MapClaims); ok {
		userID = claims["user_id"]
		userEmail = claims["email"]
	} else {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "ID: %q\n", userID)
	fmt.Fprintf(w, "E-mail: %s\n", userEmail)
}
