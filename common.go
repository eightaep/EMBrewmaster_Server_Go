package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type config struct {
	BrewsystemID         string `json:"brewsystemID"`
	SensorUpdateInterval int    `json:"sensorUpdateInterval"`
}

func readConfig() (config, error) {
	var conf config

	err := json.Unmarshal(getJSONFileContent("config.json"), &conf)

	return conf, err
}

func getJSONFileContent(jsonPath string) []byte {
	// Open test jsonFile
	jsonFile, err := os.Open(jsonPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	jsonContent, err := ioutil.ReadAll(jsonFile)

	return jsonContent
}
