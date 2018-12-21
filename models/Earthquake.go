package models

// JSONRequest tmp variable to handle Request info
type JSONRequest struct {
	Type     string          `json:"type,omitempty"`
	Metadata Metadata        `json:"metadata,omitempty"`
	Features []Earthquake    `json:"features,omitempty"`
	Bbox     map[int]float32 `json:"bbox,omitempty"`
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
	Felt     string   `json:"felt,omitempty"`
	Cdi      string   `json:"cdi,omitempty"`
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
	Nst      string   `json:"nst,omitempty"`
	Dmin     float32  `json:"dmin,omitempty"`
	Rms      float32  `json:"rms,omitempty"`
	Gap      int      `json:"gap,omitempty"`
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

// @TODO Retrieve info from provider
