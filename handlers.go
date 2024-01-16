package main 

import (
	"fmt"
	"net/http"
	"path"
	"regexp"
)


type NendoroidHandler struct {}

func (h *NendoroidHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // This is not a proper solution but is fine for now.
    isNumber, err := regexp.MatchString("\\d+$", path.Base(r.URL.Path))
    if err != nil {
        fmt.Println(err)
    }
    
    switch {
    case r.Method == http.MethodGet && isNumber:
        h.GetNendoroidById(w, r)
    case r.Method == http.MethodGet && !isNumber:
        h.GetAllNendoroids(w, r)
    }
}

func (h *NendoroidHandler) GetNendoroidById(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Nendoroid"))
}

func (h *NendoroidHandler) GetAllNendoroids(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("All Nendoroids"))
}

type HomeHandler struct{}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hi mom"))
}


// Generic 404 and 500 error handlers

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("404 Not Found"))
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 Internal Server Error"))
}
