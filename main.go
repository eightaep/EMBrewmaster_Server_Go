package main

import (
	"fmt"
	"log"
)

var srv server

func main() {
	// open config
	conf, errConf := readConfig()
	if errConf != nil {
		log.Fatalln("Error reading config: ", errConf)
		return
	}

	// run server
	srv = server{}
	errRun := srv.run(conf)
	if errRun != nil {
		log.Fatalln("Error running server: ", errRun)
	}
	defer func() {
		srv.state.State = "stopped"
		srv.updateServerState()
	}()

	// listen to console
	fmt.Println("Type \"stop\" to stop server")
	scanIn := ""
	for scanIn != "stop" {
		fmt.Scanln(&scanIn)
	}
}
