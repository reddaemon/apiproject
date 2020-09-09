package handlers

import (
	"fmt"
	"net/http"
)

// Default root handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello! You are in home!")
}
