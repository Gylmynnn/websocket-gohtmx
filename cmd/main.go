package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
	if r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
}







