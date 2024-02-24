package cookies

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name: "Jhon Doe",
		Age:  21,
	}

	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(&user); err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "auth",
		Value:    buf.String(),
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	if err := writeEncrypted(w, cookie, secretKey); err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("cookie %v set", cookie.Name)))
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	auth := cookies[0]

	cookie, err := readEncrypted(r, auth.Name, secretKey)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
			return
		default:
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
	}

	var user User

	reader := strings.NewReader(cookie)

	if err := gob.NewDecoder(reader).Decode(&user); err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Name: %q\n", user.Name)
	fmt.Fprintf(w, "Age: %d\n", user.Age)
}
