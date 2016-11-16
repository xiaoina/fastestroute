package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RouteService provides operations on strings.
type RouteService interface {
	Route([]string) ([]GoogleMapsResponse, error)
}

type routeService struct{}

func (routeService) Route(s []string) ([]GoogleMapsResponse, error) {
	unorderedResponses := make([]GoogleMapsResponse, len(s))
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
		response, err := unmarshallLocations([]byte(body))
		if err != nil {
			panic(err.Error())
		}
		unorderedResponses[i] = *response
	}
	responses, err := calculateRoute(g)
	if err != nil {
		panic(err.Error())
	}
	return responses, nil
}

func calculateRoute(g []GoogleMapsResponse) ([]GoogleMapsResponse, error) {
	orderedResponses := make([]GoogleMapsResponse, len(g))

	return orderedResponses, nil
}

//...
func unmarshallLocations(body []byte) (*GoogleMapsResponse, error) {
	var s = new(GoogleMapsResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// ServiceMiddleware is a chainable behavior modifier for StringService.
type ServiceMiddleware func(RouteService) RouteService
