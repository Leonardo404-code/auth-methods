package main

import (
	"log"
	"net/http"
	"os"
	"time"

	basicauth "github.com/leonardo404-code/auth-methods/basic-auth"
)

func init() {
	os.Setenv("USERNAME", "leonardo")
	os.Setenv("PASSWORD", "12345678")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/basicAuth", basicauth.BasicAuth(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello world"))
		}),
	)

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
