package main

import (
	"example/go-tasks/pkg"
	"fmt"
	"log"
	"net/http"
)

func main() {
	pkg.SetRouter()

	// Start the server
	fmt.Printf("Listening on port 8080\n")
	if err := http.ListenAndServe(":8080", pkg.R); err != nil {
		log.Fatal(err)
	}
}
