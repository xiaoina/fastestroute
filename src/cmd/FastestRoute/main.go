package main

import (
	"encoding/json"
	"fmt"
	"os"

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

	fmt.Println("Locations are:")

	for _, value := range route.Locations {
		_ = json.NewEncoder(os.Stdout).Encode(
			value,
		)
	}
}
