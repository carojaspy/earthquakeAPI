package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/carojaspy/earthquakeAPI/utils"
)

// JSONRequest tmp variable to handle Request info
type JSONRequest struct {
	Type     string       `json:"type,omitempty"`
	Metadata Metadata     `json:"metadata,omitempty"`
	Features []Earthquake `json:"features,omitempty"`
	Bbox     []float32    `json:"bbox,omitempty"`
}

// Metadata .
type Metadata struct {
	Generated float64 `json:"generated,omitempty"`
	Url       string  `json:"url,omitempty"`
	Title     string  `json:"title,omitempty"`
	Status    int     `json:"status,omitempty"`
	Api       string  `json:"api,omitempty"`
	Count     int     `json:"count,omitempty"`
}

// Geometry .
type Geometry struct {
	//
	Type        string    `json:"base,omitempty"`
	Coordinates [3]string `json:"base,omitempty"`
}

// Propertie .
type Propertie struct {
	Mag      float32  `json:"mag,omitempty"`
	Place    string   `json:"place,omitempty"`
	Time     float64  `json:"time,omitempty"`
	Updated  float64  `json:"updated,omitempty"`
	Tz       int      `json:"tz,omitempty"`
	Url      string   `json:"url,omitempty"`
	Detail   string   `json:"detail,omitempty"`
	Felt     int      `json:"felt,omitempty"`
	Cdi      float32  `json:"cdi,omitempty"`
	Mmi      float32  `json:"mmi,omitempty"`
	Alert    string   `json:"alert,omitempty"`
	Status   string   `json:"status,omitempty"`
	Tsunami  int      `json:"tsunami,omitempty"`
	Sig      int      `json:"sig,omitempty"`
	Net      string   `json:"net,omitempty"`
	Code     string   `json:"code,omitempty"`
	Ids      string   `json:"ids,omitempty"`
	Sources  string   `json:"sources,omitempty"`
	Types    string   `json:"types,omitempty"`
	Nst      int      `json:"nst,omitempty"`
	Dmin     float32  `json:"dmin,omitempty"`
	Rms      float32  `json:"rms,omitempty"`
	Gap      float32  `json:"gap,omitempty"`
	MagType  string   `json:"magType,omitempty"`
	Type     string   `json:"type,omitempty"`
	Title    string   `json:"title,omitempty"`
	Geometry Geometry `json:"geometry,omitempty"`
	Id       string   `json:"id,omitempty"`
}

// Earthquake defining model to store Earthquake info
type Earthquake struct {
	//
	Id         string    `json:"id,omitempty"`
	Type       string    `json:"type,omitempty"`
	Properties Propertie `json:"properties,omitempty"`
	Geometry   Geometry  `json:"geometry,omitempty"`
}

// CRUD Interface with methods to persist info
type CRUD interface {
	Get() error
	Save() error
	Update() error
	Delete() error
}

// Get .
func (e Earthquake) Get() error {
	return nil
}

// Save .
func (e Earthquake) Save() error {
	return nil
}

// Update .
func (e Earthquake) Update() error {
	return nil
}

// Delete .
func (e Earthquake) Delete() error {
	return nil
}

// GetEarthQuakeFromAPI Returns the Weather info from API Provider
func GetEarthQuakeFromAPI(ejson *JSONRequest) error {
	// Example URL to get Weather

	querystring := fmt.Sprintf("https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&%s", utils.GetPeriod())
	// log.Println(querystring)
	resp, err := http.Get(querystring)
	if err != nil {
		// handle error
		log.Println("Error getting Weather from api.openweathermap.org")
		return err
	}
	if resp.StatusCode != 200 {
		log.Println("Code Status invalid: ", resp.StatusCode)
		return errors.New("Code Status invalid")
	}
	// Closing conection
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Error reading result from api.openweathermap.org")
		return errors.New("Error reading result from api.openweathermap.org")
	}

	// Building response
	err = json.Unmarshal(body, &ejson)
	if err != nil {
		fmt.Println("Error in unmarshall: ", err)
		return errors.New("Error Unpacking WeatherAPI info")
	}
	// log.Println(ejson)

	// Building response
	return nil
} //end GetWeatherFromProvider

// ToMapEarthquakes Converts the array of Earthquakes giving by the provider, to a Map using the ID as key
func ToMapEarthquakes(arr []Earthquake) map[string]Earthquake {
	edict := make(map[string]Earthquake)
	for _, element := range arr {
		edict[element.Id] = element
	}
	return edict
	// return edict, errors.New("")
}
