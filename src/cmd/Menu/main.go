package Menu

import (
	"bufio"
	"fmt"
	"os"

	"cmd/Mapping"
)

type Menu struct{}

var (
	route Mapping.Route
)

func (m *Menu) Intro() {
	fmt.Println("Welcome to fastest route")
}

func (m *Menu) GetLocations() Mapping.Route {
	reader := bufio.NewReader(os.Stdin)

	var stringLocations []string
	var length int

	fmt.Print("Enter starting location address: ")
	start, _ := reader.ReadString('\n')
	stringLocations = append(stringLocations, start)
	fmt.Print("Enter ending location address: ")
	end, _ := reader.ReadString('\n')
	stringLocations = append(stringLocations, end)
	fmt.Print("Enter the number of stops you are making in between: ")
	fmt.Scan(&length)

	for i := 0; i < length; i++ {
		fmt.Print("Enter your next stop: ")
		s, _ := reader.ReadString('\n')
		stringLocations = append(stringLocations, s)
	}

	route.AppendStringLocations(stringLocations)

	return route
}

func (m *Menu) EnterCommand(s string) {

}
