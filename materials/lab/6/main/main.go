package main

// main.go HAS FOUR TODOS - TODO_1 - TODO_4

import (
	"log"
	"net/http"
	"scrape/scrape"

	"github.com/gorilla/mux"
)

// I can't get scrapeapi to import this package, so I put LOG_LEVEL in there
// SEE SCRAPEAPI.GO FOR LOG_LEVEL

//TODO_1: Logging right now just happens, create a global constant integer LOG_LEVEL
//TODO_1: When LOG_LEVEL = 0 DO NOT LOG anything
//TODO_1: When LOG_LEVEL = 1 LOG API details only
//TODO_1: When LOG_LEVEL = 2 LOG API details and file matches (e.g., everything)

func main() {

	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/", scrape.MainPage).Methods("GET")

	router.HandleFunc("/api-status", scrape.APISTATUS).Methods("GET")

	router.HandleFunc("/indexer", scrape.IndexFiles).Methods("GET")
	router.HandleFunc("/search", scrape.FindFile).Methods("GET")
	router.HandleFunc("/addsearch/{regex}", scrape.AddRE).Methods("GET")
	router.HandleFunc("/clear", scrape.Clear).Methods("GET")
	router.HandleFunc("/reset", scrape.Reset).Methods("GET")

	http.Handle("/", router)

	//start and listen to requestsS
	http.ListenAndServe(":3000", router)

}
