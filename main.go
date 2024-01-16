package main

import (
	"net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.Handle("/", &HomeHandler{})
    mux.Handle("/nendoroid", &NendoroidHandler{})
    mux.Handle("/nendoroid/", &NendoroidHandler{})
    http.ListenAndServe(":8080", mux)
}
