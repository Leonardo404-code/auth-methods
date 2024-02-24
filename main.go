package main

import (
	"log"
	"net/http"
	"os"
	"time"

	basicauth "github.com/leonardo404-code/auth-methods/basicAuth"
	"github.com/leonardo404-code/auth-methods/cookies"
)

func init() {
	os.Setenv("USERNAME", "jhon")
	os.Setenv("PASSWORD", "12345678")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/basic_auth", basicauth.BasicAuth)
	mux.HandleFunc("/set_cookie", cookies.SetCookie)
	mux.HandleFunc("/get_cookie", cookies.GetCookie)

	server := &http.Server{
		Addr:         ":3000",
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("startin server on %s", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
