package cookies

import (
	"errors"
	"fmt"
	"net/http"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "auth",
		Value:    "Hello world",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	w.Write([]byte(fmt.Sprintf("cookie %v set", cookie.Name)))
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			http.Error(w, "server error", http.StatusInternalServerError)
		}
	}

	w.Write([]byte(cookie.Value))
}
