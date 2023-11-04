package main

import (
	"fmt"
	"github.com/skripov-ds-ai/social_network/generated"
	"log"
	"net/http"
)

func main() {
	port := "8080"
	srv, err := generated.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(fmt.Sprint(":%s", port), srv); err != nil {
		log.Fatal(err)
	}
}
