package main

import (
	"go-ts-vite/ui"
	"io/fs"
	"log"
	"net/http"
)

func makeHandler(str string, pre string) http.HandlerFunc {
	sub, err := fs.Sub(ui.Index, str)
	if err != nil {
		log.Panicln(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix(pre, http.FileServerFS(sub)).ServeHTTP(w, r)
	}
}

func muxInit() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", makeHandler("staticPages", "/"))
	mux.HandleFunc("/app/", makeHandler("frontend/dist", "/app"))

	return mux
}

func main() {
	log.Println("tryin spinnin")
	log.Println("CTRL + C to stop spinnin n shit")

	mux := muxInit()

	log.Println("spinnin in port", 8008)
	// starting server
	err := http.ListenAndServe(":8008", mux)
	if err != nil {
		log.Fatal(err)
	}
}
