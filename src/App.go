package main

import (
	"Project/functions"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", functions.HandleRoot)
	mux.HandleFunc("POST /users", functions.SetUser)
	mux.HandleFunc("GET /users/{id}", functions.GetUser)
	mux.HandleFunc("DELETE /users/{id}", functions.DeleteUser)
	fmt.Printf("Server listening to 8080")
	http.ListenAndServe(":8080", mux)
}
