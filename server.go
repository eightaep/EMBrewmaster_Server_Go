// main package of EMBrewmaster_Server
package main

import (
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

// represents the state of the server
type serverState struct {
	Start     time.Time `json:"start"`
	Tick      time.Time `json:"tick"`
	Name      string    `json:"name"`
	System    string    `json:"system"`
	State     string    `json:"state"`
	OS        string    `json:"os"`
	Arch      string    `json:"arch"`
	Simulated bool      `json:"simulated"`
	Key       string
}

type server struct {
	state serverState
	bs    Brewsystem
	conf  config
	fbc   IFirebaseClient
	Key   string
}

func (srv *server) updateServerState() {
	srv.state.Tick = time.Now()
	srv.fbc.UpdateServerState(srv.Key, &srv.state)
}

func (srv *server) init(conf config) error {
	srv.Key = "EMBrewserver_Go"
	srv.conf = conf

	hostname, errHost := os.Hostname()
	if errHost != nil {
		log.Println("Getting hostname failed: ", errHost)
		hostname = "<failed>"
		return errHost
	}
	srv.state = serverState{Start: time.Now(), Tick: time.Now(), Name: "GoServer-Test", System: hostname, State: "running", OS: runtime.GOOS, Arch: runtime.GOARCH}
	if strings.ToUpper(srv.state.OS) == strings.ToUpper("Windows") {
		srv.state.Simulated = true
	}

	return srv.initFirebase()
}

func (srv *server) initFirebase() error {
	var errConnFB error
	srv.fbc, errConnFB = ConnectToFirebase()
	if errConnFB != nil {
		log.Fatalln("firebase cannot be connected: ", errConnFB)
		return errConnFB
	}

	//  get brewsystem
	var errGetBS error
	srv.bs, errGetBS = srv.fbc.GetBrewsystemByID(srv.conf.BrewsystemID)

	if errGetBS != nil {
		log.Fatalln("Error getting brewsystem: ", errGetBS)
		return errGetBS
	}

	return nil
}

func (srv *server) actorServer() {

}

// sensorServer refreshes the sensor values
func (srv *server) sensorServer() {
	sens := srv.bs.GetSensors()

	for {
		time.Sleep(time.Duration(srv.conf.SensorUpdateInterval) * time.Millisecond)

		for _, s := range sens {
			srv.fbc.UpdateSensorValue(srv.bs, s, s.GetValue())
		}
	}
}

func (srv *server) updateServer() {
	for {
		time.Sleep(1 * time.Second)
		srv.updateServerState()
	}
}

func (srv *server) run(conf config) error {
	srv.init(conf)

	// start update server
	go srv.updateServer()

	// start actor server
	go srv.actorServer()

	// start sensor server
	go srv.sensorServer()

	return nil
}
