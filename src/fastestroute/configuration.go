package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

//Config holds the values for the google API
type Config struct {
	Title   string
	Key     string
	Baseurl string
}

func (c *Config) getConfiguration() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Title: %s\n", c.Title)
	fmt.Printf("Key: %s\n", c.Key)
	fmt.Printf("Baseurl: %s\n", c.Baseurl)
}
