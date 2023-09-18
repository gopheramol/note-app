// main.go

package main

import (
	"log"
	"net/http"

	"github.com/gopheramol/notesapp/config"
	"github.com/gopheramol/notesapp/db"
	"github.com/gopheramol/notesapp/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// Load config
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	client, err := db.ConnectDB(&config)
	if err != nil {
		log.Println("failed to connect DB")
		panic(err)
	}

	if client == nil {
		log.Println("erroer in connectDB: ")
		return
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetNotes(w, r, client)
	}).Methods("GET")

	r.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddNote(w, r, client)
	}).Methods("POST")

	r.HandleFunc("/delete/{title}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteNote(w, r, client)
	}).Methods("GET")

	http.Handle("/", r)

	log.Println("Server running on port: 8081")
	http.ListenAndServe(config.ServerAddress, nil)
}
