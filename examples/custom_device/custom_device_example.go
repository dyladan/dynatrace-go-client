package main

import (
	"encoding/json"
	"fmt"
	dynatrace "github.com/dyladan/dynatrace-go-client/api"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func prettyPrint(i interface{}) string {
	s, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		return fmt.Sprintf("Could not parse the object: %+v", i)
	}
	return string(s)
}

func createEvent() {

	l := log.New()
	l.Level = log.DebugLevel

	c := dynatrace.New(dynatrace.Config{
		APIKey:    os.Getenv("DT_API_TOKEN"),
		BaseURL:   os.Getenv("DT_API_URL"),
		Log:       l,
		RetryTime: 2 * time.Second,
		Retries:   5,
	})

	cd := dynatrace.CustomDevicePushMessage{
		DisplayName: "Alertmanager Events",
		Group:       "Alertmanager",
	}
	e, _, err := c.CustomDevice.Create("Alertmanager Events", cd)
	if err != nil {
		panic(err)
	}

	fmt.Println(prettyPrint(e))

}

func main() {
	createEvent()

}
