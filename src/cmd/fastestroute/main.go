package fastestroute

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	menu   Menu
	route  Route
	config Config
)

func main() {

	getConfiguration()

	menu.Intro()
	route = menu.GetLocations()

	fmt.Println("Locations are:")

	for _, value := range route.Locations {
		_ = json.NewEncoder(os.Stdout).Encode(
			value,
		)
	}
}
