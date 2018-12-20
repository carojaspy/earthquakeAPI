package controllers

import (
	"log"
	"net/http"

	"github.com/carojaspy/earthquakeAPI/utils"
)

// URLProvider URL of the provider
const URLProvider = "https://earthquake.usgs.gov/fdsnws/event/1/application.json?%s"

// URLCatalogsProvider URL of the provider
const URLCatalogsProvider = "https://earthquake.usgs.gov/fdsnws/event/1/catalogs"

// URLCountProvider URL of the provider
const URLCountProvider = "https://earthquake.usgs.gov/fdsnws/event/1/count?format=geojson&%s"

// URLQueryProvider URL of the provider
const URLQueryProvider = "https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&%s"

// GetEarthQuakes returns the list of EarthQuakes 2 days ago since now
func GetEarthQuakes(w http.ResponseWriter, r *http.Request) {
	// GetEarthQuakes
	log.Println("GetEarthQuakes")
	// Set content type to response
	utils.SetJSONResponse(&w)
}

// GetSingleEarthQuake try to get by ID the the earthquake info
func GetSingleEarthQuake(w http.ResponseWriter, r *http.Request) {
	// GetEarthQuakes
	log.Println("GetSingleEarthQuake")
	utils.SetJSONResponse(&w)
}

// DeleteEarthQuake .
func DeleteEarthQuake(w http.ResponseWriter, r *http.Request) {
	// GetEarthQuakes
	log.Println("DeleteEarthQuake")
	utils.SetJSONResponse(&w)
}

// UpdateEarthQuake .
func UpdateEarthQuake(w http.ResponseWriter, r *http.Request) {
	// GetEarthQuakes
	log.Println("UpdateEarthQuake")
	utils.SetJSONResponse(&w)
}
