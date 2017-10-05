package config

import (
	"os"
	"fmt"
	"encoding/json"
)

type Provider struct {
	Name   string            `json:"name"`
	Params map[string]string `json:"params"`
}

type Configuration struct {
	Providers []Provider `json:"providers"`
}

func (c *Configuration) GetProvider(name string) *Provider {
	for _, prov := range c.Providers {
		if prov.Name == name {
			return &prov
		}
	}
	return nil
}

func (c *Configuration) LoadConfiguration(file string) {
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(c)
}
