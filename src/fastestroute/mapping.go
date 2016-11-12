package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//GoogleMapsResponse ...
type GoogleMapsResponse struct {
	Results Results `json:"results"`
	Status  string  `json:"status"`
}

//Results ...
type Results []struct {
	AddressComponents []AddressComponents `json:"address_components"`
	FormattedAddress  string              `json:"formatted_address"`
	Geometry          Geometry            `json:"geometry"`
	PartialMatch      bool                `json:"partial_match"`
	PlaceID           string              `json:"place_id"`
	Types             []string            `json:"types"`
}

//AddressComponents ...
type AddressComponents struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

//Geometry ...
type Geometry struct {
	Location     Location `json:"location"`
	LocationType string   `json:"location_type"`
	Viewport     Viewport `json:"viewport"`
}

//Location ...
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

//Viewport ...
type Viewport struct {
	Northeast `json:"northeast"`
	Southwest `json:"southwest"`
}

//Northeast ...
type Northeast struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

//Southwest ...
type Southwest struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

//Route ...
type Route struct {
	Locations     []GoogleMapsResponse `json:"response"`
	TotalDistance float64              `json:"totaldistance"`
}

//...
func getLocations(body []byte) (*GoogleMapsResponse, error) {
	var s = new(GoogleMapsResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

//GetLocation ...
func (r *Route) GetLocation(s []string) {
	responses := make([]GoogleMapsResponse, len(s))
	for i := 0; i < len(s); i++ {
		resp, err := http.Get(fmt.Sprintf(config.Baseurl, TrimString(s[i]), config.Key))
		if err != nil {
			panic(err.Error())
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		response, err := getLocations([]byte(body))
		if err != nil {
			panic(err.Error())
		}
		responses[i] = *response
	}
	r.Locations = responses
}
