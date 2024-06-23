package main

import (
	"fmt"
	"log"
	"net/http"

	"sportnation/routes"
)

func main() {
    router := routes.SetupRoutes()
    fmt.Println("Server is listening on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", router))
}
