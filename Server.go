package main

import (
	"log"
	"net/http"

	"github.com/carojaspy/earthquakeAPI/controllers"
	"github.com/gorilla/mux"
)

func main() {
	log.Print("Running Server on localhost:8000...")

	router := mux.NewRouter()

	// Listing all earthquakes
	router.HandleFunc("/earthquakes", controllers.GetEarthQuakes).Methods("GET")

	//	Operations over a earthquake
	// get earthquake by id
	router.HandleFunc("/earthquakes/{id}", controllers.GetSingleEarthQuake).Methods("GET")
	// delete earthquake
	router.HandleFunc("/earthquakes/{id}", controllers.DeleteEarthQuake).Methods("DELETE")
	// update earthquake
	router.HandleFunc("/earthquakes/{id}", controllers.UpdateEarthQuake).Methods("PUT")

	// Running server
	log.Fatal(http.ListenAndServe(":8000", router))
}
