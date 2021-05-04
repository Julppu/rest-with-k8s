package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server operational!")
}

func main() {
	fmt.Println("Started server, listening on port 3000.")
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
