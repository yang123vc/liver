package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var cfg *config

func initConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err.Error())
	}
	cfgByte, _ := ioutil.ReadAll(file)
	json.Unmarshal(cfgByte, &cfg)
}
