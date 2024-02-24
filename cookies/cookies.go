package cookies

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "auth",
		Value:    "Hello Zoë!",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	if err := write(w, cookie); err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("cookie %v set", cookie.Name)))
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := read(r, "auth")
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

	w.Write([]byte(cookie))
}
