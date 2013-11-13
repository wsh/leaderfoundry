package main

import (
	"encoding/json"
	"os"
)

type TwitterAPIConfig struct {
	Key    string
	Secret string
}

func readconfig() *TwitterAPIConfig {
	file, err := os.Open("auth.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	config := &TwitterAPIConfig{}
	decoder.Decode(&config)
	return config
}
