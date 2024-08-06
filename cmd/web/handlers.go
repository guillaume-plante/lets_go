package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Snippet Create Post"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Snippet Create"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Snippet View with id %d", id)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Go")
	w.Write([]byte("Hello World"))
}
