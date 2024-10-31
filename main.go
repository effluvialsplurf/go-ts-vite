package main

import (
	"fmt"
	"go-ts-vite/ui"
	"io/fs"
	"log"
	"net/http"
)

// frontend app struct
type frontendAppHandler struct {
	// global entrypoint for frontend app
	index fs.FS
}

var tsHandler = frontendAppHandler{
	index: nil,
}

// handlers
func (h *frontendAppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.FS(tsHandler.index)).ServeHTTP(w, r)
}

// url paths
func makeUrls() {
	http.HandleFunc("/", tsHandler.ServeHTTP)
}

func main() {
	var err error = nil

	// get filesystem for frontend app
	tsHandler.index, err = fs.Sub(ui.Index, "frontend/templates")
	if err != nil {
		panic(err)
	}

	log.Println("tryin spinnin")

	// registering the handlers, and making url paths
	makeUrls()

	log.Println("spinnin in port", 8008)
	fmt.Println("CTRL + C to stop spinnin n shit")

	// starting server
	err = http.ListenAndServe(":8008", nil)
	if err != nil {
		log.Fatal(err)
	}
}
