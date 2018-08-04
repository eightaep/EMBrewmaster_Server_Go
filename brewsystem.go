package main

import (
	"encoding/json"
)

// BrewsystemItem represents the generic attributes of every part of the brewsystem
type BrewsystemItem struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// Brewsystem represents the brewsystem
type Brewsystem struct {
	BrewsystemItem

	Kettles []Kettle `json:"kettles"`
}

// Kettle represents a kettle in a brewsystem
type Kettle struct {
	BrewsystemItem

	Function string   `json:"function"`
	No       int      `json:"no"`
	Size     int      `json:"size"`
	Actors   []Actor  `json:"actors"`
	Sensors  []Sensor `json:"sensors"`
}

// InitBrewsystemsFromJSON initializes a brewsystem from a JSON byte array
func InitBrewsystemsFromJSON(bsJSON []byte) []Brewsystem {
	var bs []Brewsystem

	json.Unmarshal(bsJSON, &bs)

	return bs
}

// InitBrewsystemsFromJSONFile initializes a brewsystem from a JSON file
func InitBrewsystemsFromJSONFile(bsJSONPath string) []Brewsystem {
	return InitBrewsystemsFromJSON(getJSONFileContent(bsJSONPath))
}

// GetActors returns all actors from the brewsystem
func (bs Brewsystem) GetActors() []Actor {
	var acts []Actor

	for _, k := range bs.Kettles {
		acts = append(acts, k.Actors...)
	}

	return acts
}

// GetSensors returns all sensors from the brewsystem
func (bs Brewsystem) GetSensors() []Sensor {
	var sens []Sensor

	for _, k := range bs.Kettles {
		sens = append(sens, k.Sensors...)
	}

	return sens
}
