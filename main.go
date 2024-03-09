package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("John")
	http.Handle("/welcome", templ.Handler(component))
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
