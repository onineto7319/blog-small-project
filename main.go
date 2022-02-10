package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/blog-small-project/internal/router"
)

var (
	port = "8080"
)

// @title 部落格系統
// @version v1.0
// @termsOfService https://github.com/blog-small-project
func main() {
	if herokuPort := os.Getenv("PORT"); herokuPort != "" {
		port = os.Getenv("PORT")
	}

	router := router.New()

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, //1MB
	}

	log.Fatal(s.ListenAndServe())
}
