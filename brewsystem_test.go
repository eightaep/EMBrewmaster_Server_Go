package main

import (
	"testing"
)

func getBrewsystem() Brewsystem {
	bss := InitBrewsystemsFromJSON(getJSONFileContent("testdata/Brewsystems.json"))
	return bss[0]
}
func TestInitFromJSON(t *testing.T) {
	bss := InitBrewsystemsFromJSON(getJSONFileContent("testdata/Brewsystems.json"))

	if bss == nil {
		t.Fail()
	}

	BrewsystemsTest(t, bss)
}

func TestInitFromJSONFile(t *testing.T) {
	bss := InitBrewsystemsFromJSONFile("testdata/Brewsystems.json")

	if bss == nil {
		t.Fail()
	}

	BrewsystemsTest(t, bss)
}

func BrewsystemsTest(t *testing.T, bss []Brewsystem) {
	for _, bs := range bss {
		BrewsystemTest(t, bs)
	}
}

func BrewsystemTest(t *testing.T, bs Brewsystem) {
	if len(bs.Name) == 0 {
		t.Fail()
	}

	if len(bs.Kettles) == 0 {
		t.Fail()
	}

	if len(bs.Kettles[0].Actors) == 0 {
		t.Fail()
	}

	if len(bs.Kettles[0].Sensors) == 0 {
		t.Fail()
	}
}

func TestGetJSONFileContent(t *testing.T) {
	jsonContent := getJSONFileContent("testdata/Brewsystems.json")

	if len(jsonContent) == 0 {
		t.Fail()
	}
}

func TestGetSensors(t *testing.T) {
	bs := getBrewsystem()
	sens := bs.GetSensors()
	if len(sens) <= 0 {
		t.Fail()
	}
}
func TestGetActors(t *testing.T) {
	bs := getBrewsystem()
	acts := bs.GetActors()
	if len(acts) <= 0 {
		t.Fail()
	}
}
