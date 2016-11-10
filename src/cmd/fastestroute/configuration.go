package fastestroute

import (
	"github.com/BurntSushi/toml"
)

const (
	tomlFile = "~/configs/config.toml"
)

//Config holds the values for the google API
type Config struct {
	key     string
	baseurl string
}

func getConfiguration() {
	if _, err := toml.Decode(tomlFile, &config); err != nil {
		panic(err.Error())
	}
}
