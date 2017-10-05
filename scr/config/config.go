package config

type Provider struct {
	Name   string                 `json:"name"`
	Params map[string]interface{} `json:"params"`
}

type Configuration struct {
	Providers []Provider `json:"providers"`
}