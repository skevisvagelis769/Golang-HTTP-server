package main

import (
	"cpy/funcs"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", funcs.SetUser)
	mux.HandleFunc("GET /users/search/{id}", funcs.GetUser)
	mux.Handle("/", fs)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}
