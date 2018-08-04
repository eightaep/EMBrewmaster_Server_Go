package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"golang.org/x/net/context"
)

func TestFirebaseRead(t *testing.T) {
	fbc := connectFirebaseTest(t)

	fetchDataFromFirebaseTest(fbc, t)
	getBrewsystemsTest(fbc, t)
	getBrewsystemByIDTest(fbc, t)
}

func TestFirebaseWrite(t *testing.T) {
	fbc := connectFirebaseTest(t)
	updateServerStateTest(fbc, t)
}

func connectFirebaseTest(t *testing.T) firebaseClient {
	fbc, err := ConnectToFirebase()

	if (err != nil) || (fbc == nil) {
		t.Fail()
	}

	return fbc.(firebaseClient)
}

func fetchDataFromFirebaseTest(fbc firebaseClient, t *testing.T) {
	var bss []Brewsystem
	err := fbc.client.NewRef("/brewsystemsGo").Get(context.Background(), &bss)

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	BrewsystemsTest(t, bss)
}

func getBrewsystemsTest(fbc firebaseClient, t *testing.T) {
	bss, err := fbc.GetBrewsystems()

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	BrewsystemsTest(t, bss)
}

func updateServerStateTest(fbc firebaseClient, t *testing.T) {
	hostname, _ := os.Hostname()
	state := serverState{Start: time.Now(), Tick: time.Now(), Name: "GoServer-Test", System: hostname}

	err := fbc.UpdateServerState("Test_Server_Go", &state)

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func getBrewsystemByIDTest(fbc firebaseClient, t *testing.T) {
	bs, err := fbc.GetBrewsystemByID("RecipeTestSystem")

	if err != nil {
		t.Fail()
	}

	var bss []Brewsystem
	BrewsystemsTest(t, append(bss, bs))
}
