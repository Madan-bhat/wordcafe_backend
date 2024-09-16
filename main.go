package main

import (
	"fmt"
	"log"
	"net/http"

	"wordcafe_backend/config"
	"wordcafe_backend/controllers"
)

func main() {
	config.ConnectMongo()

	http.HandleFunc("/register", controllers.RegisterHandler)
	http.HandleFunc("/login", controllers.LoginHandler)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
