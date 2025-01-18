package api

import (
	"fmt"
	"net/http"
)

// EntryPoint handles all requests
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from Vercel!")
}
