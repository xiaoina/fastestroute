package main

import (
	//"encoding/json"
	"fmt"
	//"os"

	"cmd/Mapping"
	"cmd/Menu"
	//"github.com/mitchellh/mapstructure"
)

var (
	menu  Menu.Menu
	route Mapping.Route
)

func main() {
	menu.Intro()
	route = menu.GetLocations()

	fmt.Println("Locations are: \n")

	for _, value := range route.Locations {
		fmt.Println(value)
		fmt.Println("\n")
	}

	/*	route := Mapping.Route{
		Locations:     []Mapping.LocationPoint{location1, location2},
		TotalDistance: 69,
	}*/

}
