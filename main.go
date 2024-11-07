package main

import (
	"go-ts-vite/ui"
	"io/fs"
	"log"
	"net/http"
	"text/template"
)

func loadApp(w http.ResponseWriter, r *http.Request) {
	app, err := fs.Sub(ui.Index, "frontend/dist")
	if err != nil {
		log.Panicln(err)
	}
	http.StripPrefix("/app", http.FileServerFS(app)).ServeHTTP(w, r)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("./ui/templates/index.html"))

	err := templ.Execute(w, nil)
	if err != nil {
		log.Panic(err)
	}
}

func muxInit() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/app/", loadApp)
	mux.HandleFunc("/", templateHandler)

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
