package Mapping

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	key      = "AIzaSyD_HgYTZ0kHqXwStMx3JfLitzkrimNBl0E"
	base_url = "https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s"
	encoder  = json.NewEncoder(os.Stdout)
)

type GoogleMapsResponse struct {
	Results Results `json:"results"`
	Status  string  `json:"status"`
}

type Results []struct {
	AddressComponents []AddressComponents `json:"address_components"`
	FormattedAddress  string              `json:"formatted_address"`
	Geometry          Geometry            `json:"geometry"`
	PartialMatch      bool                `json:"partial_match"`
	PlaceID           string              `json:"place_id"`
	Types             []string            `json:"types"`
}

type AddressComponents struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type Geometry struct {
	Location     Location `json:"location"`
	LocationType string   `json:"location_type"`
	Viewport     Viewport `json:"viewport"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Viewport struct {
	Northeast `json:"northeast"`
	Southwest `json:"southwest"`
}

type Northeast struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Southwest struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type LocationPoint struct {
	X       float64 `json:"X"`
	Y       float64 `json:"Y"`
	Name    string  `json:"name"`
	Address string  `json:"address"`
}

type Route struct {
	Locations     []LocationPoint `json:"locations"`
	TotalDistance float64         `json:"totaldistance"`
}

func getLocations(body []byte) (*GoogleMapsResponse, error) {
	var s = new(GoogleMapsResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

func (r *Route) GetLocation() {
	resp, err := http.Get(fmt.Sprintf(base_url, r.Locations[0].Address, key))
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	s, err := getLocations([]byte(body))
	encoder.Encode(s)
}

func (r *Route) AppendStringLocations(s []string) {
	for index, value := range s {
		r.Locations[index] := LocationPoint{
			X:       0,
			Y:       0,
			Name:    "",
			Address: value,
		}
	}
}
