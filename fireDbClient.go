package main

import (
	"log"

	"firebase.google.com/go/db"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// IFirebaseClient collects all available methods on firebase
type IFirebaseClient interface {
	GetBrewsystems() ([]Brewsystem, error)
	IsConnected() bool
	GetBrewsystemByID(brewsystemID string) (Brewsystem, error)
	UpdateServerState(srvKey string, state *serverState) error
	UpdateSensorValue(bs Brewsystem, s Sensor, sv SensorValue) error
}

type firebaseClient struct {
	app    *firebase.App
	client *db.Client
	ctx    context.Context
}

func (fbc firebaseClient) IsConnected() bool {
	if (fbc.app == nil) || (fbc.client == nil) {
		return false
	}

	return true
}

// ConnectToFirebase creates connection to firebase
func ConnectToFirebase() (IFirebaseClient, error) {
	var err error
	fbc := firebaseClient{ctx: context.Background()}
	conf := &firebase.Config{
		DatabaseURL: "https://embrewmaster-810608.firebaseio.com",
	}
	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("embrewmaster-810608-firebase-adminsdk.json")

	// Initialize the app with a service account, granting admin privileges
	fbc.app, err = firebase.NewApp(fbc.ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	fbc.client, err = fbc.app.Database(fbc.ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	return fbc, err
}

func (fbc firebaseClient) GetBrewsystems() ([]Brewsystem, error) {
	var bss []Brewsystem

	// check if firebase is connected
	if !fbc.IsConnected() {
		err := errorString{s: "firebase not connected"}
		return nil, err
	}

	// load brewsystems
	errFetch := fbc.client.NewRef("/brewsystemsGo").Get(fbc.ctx, &bss)

	if errFetch != nil {
		log.Fatalln("Error fetching brewsystems:", errFetch)
		return nil, errFetch
	}

	return bss, nil
}

func (fbc firebaseClient) GetBrewsystemByID(brewsystemID string) (Brewsystem, error) {
	bs := Brewsystem{}
	bss, errGetBS := fbc.GetBrewsystems()

	if errGetBS != nil {
		log.Fatalln("Error getting brewsystems: ", errGetBS)
		return bs, errGetBS
	}

	for _, bsIterate := range bss {
		if bsIterate.Key == brewsystemID {
			return bsIterate, nil
		}
	}

	return bs, errorString{s: "BrewsystemID '" + brewsystemID + "' not found"}
}

func (fbc firebaseClient) UpdateServerState(srvKey string, state *serverState) error {
	refPath := "/server/" + srvKey

	// generate key if empty
	if state.Key == "" {
		refKey, errGenKey := fbc.client.NewRef(refPath).Push(fbc.ctx, nil)

		if errGenKey != nil {
			return errGenKey
		}

		state.Key = refKey.Key
	}

	// write server state
	stateInterface := map[string]interface{}{"name": state.Name, "start": state.Start, "system": state.System, "tick": state.Tick, "state": state.State, "os": state.OS, "arch": state.Arch, "simulated": state.Simulated}
	errUpdate := fbc.client.NewRef(refPath+"/"+state.Key).Update(fbc.ctx, stateInterface)

	return errUpdate
}

func (fbc firebaseClient) UpdateSensorValue(bs Brewsystem, s Sensor, sv SensorValue) error {
	refPath := "/automation/" + bs.Key + "/sensors/" + s.Key + "/measurement"
	errUpdate := fbc.client.NewRef(refPath).Update(fbc.ctx, map[string]interface{}{"value": sv.Value, "time": sv.Time})
	return errUpdate
}

// As an admin, the app has access to read and write all data, regradless of Security Rules
// ref := client.NewRef("restricted_access/secret_document")
// var data map[string]interface{}
// if err := ref.Get(ctx, &data); err != nil {
//         log.Fatalln("Error reading from database:", err)
// }
// fmt.Println(data)
