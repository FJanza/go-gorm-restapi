package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/posts", GetPosts).Methods("GET")
	r.HandleFunc("/post", GetPost).Methods("GET")
	r.HandleFunc("/posts", AddPost).Methods("POST")

	http.ListenAndServe(":3000", r)
}
