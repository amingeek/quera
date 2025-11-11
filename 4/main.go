package main

import (
	"fmt"
	"gravatar/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/avatar", handlers.HandleGravatarRequest)
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
