package http

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func StartServer() {
	fmt.Println("Starting server on port", port)

	mux := http.NewServeMux()

	AddRoutes(mux)

	log.Fatal(http.ListenAndServe(port, mux))
}
