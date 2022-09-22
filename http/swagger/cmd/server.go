package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed dist
var dist embed.FS

func main() {
	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.FS(dist))))

	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		log.Fatal(err)
	}
}
