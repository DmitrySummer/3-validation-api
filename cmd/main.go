package main

import (
	"3-validation-api/config"
	"3-validation-api/internal/verifi"
	"fmt"
	"net/http"
)

func main() {
	conf := config.LoadConfig()
	router := http.NewServeMux()
	verifi.NewAuthHandler(router, verifi.VerifiHandler{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Server listen on port 8081")
	server.ListenAndServe()
}
