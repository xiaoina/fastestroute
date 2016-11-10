package fastestroute

import (
	"bufio"
	"fmt"
	"os"
)

//Menu ...
type Menu struct{}

//Intro prints the intro to the fastest route generator
func (m *Menu) Intro() {
	fmt.Println("Welcome to fastest route")
}

//GetLocations ...
func (m *Menu) GetLocations() Route {
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

	route.GetLocation(stringLocations)

	return route
}
