package main

// Actor defines an actor in the brewsystem
type Actor struct {
	BrewsystemItem

	MinusPin   int    `json:"minusPin"`
	PlusPin    int    `json:"plusPin"`
	StandbyPin int    `json:"standbyPin"`
	Pin        int    `json:"Pin"`
	Type       string `json:"type"`
}
