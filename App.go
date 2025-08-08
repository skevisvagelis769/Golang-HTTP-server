package main

import (
	"cpy/funcs"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	mux := http.NewServeMux()
	mux.Handle("/", fs)
	mux.HandleFunc("/search_id", funcs.SrchId)
	mux.HandleFunc("/search_name", funcs.SrchName)

	mux.HandleFunc("POST /users", funcs.SetUser)
	mux.HandleFunc("GET /users/{id}", funcs.GetUser)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
