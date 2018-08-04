package main

import (
	"math/rand"
	"time"
)

// Sensor defines a sensor
type Sensor struct {
	BrewsystemItem

	OneWireID string `json:"1-wire-id"`
	Type      string `json:"type"`
	Unit      string `json:"unit"`
}

// DS18B20 temperature sensor
type DS18B20 struct {
	Sensor
}

// SensorValue returns value and time the value was taken
type SensorValue struct {
	Value float32   `json:"value"`
	Time  time.Time `json:"time"`
}

// GetValue returns the sensor value
func (s Sensor) GetValue() SensorValue {
	sv := SensorValue{Value: 0.0, Time: time.Now()}

	if srv.state.Simulated {
		sv.Value = rand.Float32() * 100
	}

	return sv
}
