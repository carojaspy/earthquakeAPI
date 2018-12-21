package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// SetJSONResponse .
func SetJSONResponse(w *http.ResponseWriter) {
	//Set Header
	// Set content type to response
	(*w).Header().Set("Content-Type", "application/json")
}

// GetPeriod by the moment, returns by default a interval of 1 day
func GetPeriod() string {
	// @TODO returns the dates since 2 months ago until now
	return "starttime=2018-12-19&endtime=2018-12-20"
}

// SerializeObject .
func SerializeObject(elm interface{}) []byte {
	jData, err := json.Marshal(elm)
	if err != nil {
		// handle error
		log.Println("Error sending JSON response: ", err)
	}
	return jData
}
