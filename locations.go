package main

import "net/http"

func (h Handler) LocationHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}
