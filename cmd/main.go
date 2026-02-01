package main

import (
	"example/tasks/pkg"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	r, dbConnection := pkg.SetRouter()
	defer dbConnection.Close()
	// Start the server
	fmt.Printf("Listening on port 8080\n")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
