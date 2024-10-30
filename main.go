package main

import (
	"fmt"
	"log"
	"net/http"
)

// handlers
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Henlo wurlnt")
}

// url paths
func makeUrls() {
	http.HandleFunc("/", homepage)
}

func main() {
	log.Println("tryin spinnin")

	// registering the handlers, and making url paths
	makeUrls()

	log.Println("spinnin in port", 8008)
	fmt.Println("CTRL + C to stop spinnin n shit")

	// starting server
	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		log.Fatal(err)
	}
}
