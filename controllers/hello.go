package controllers

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    // Set the content type to application/json
    w.Header().Set("Content-Type", "application/json")
    // Write the response
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `{"message": "Hello, World!"}`)
}