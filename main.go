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

var tsHandler = frontendAppHandler{}

// handlers
func (h *frontendAppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//file, err := fs.Sub(ui.Index, ".")
	//if err != nil {
	//	panic(err)
	//}
	http.FileServer(http.FS(ui.Index)).ServeHTTP(w, r)
}

// url paths
func makeUrls() {
	http.HandleFunc("/", tsHandler.ServeHTTP)
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
