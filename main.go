package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	log.Println("Listening at http://localhost:8080")
	err := http.ListenAndServe("0.0.0.0:8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
