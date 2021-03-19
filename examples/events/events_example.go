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

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func createEvent() {

	l := log.New()
	l.Level = log.DebugLevel

	c := dynatrace.New(dynatrace.Config{
		APIKey:  os.Getenv("DT_API_KEY"),
		BaseURL: os.Getenv("DT_BASE_URL"),
		Log:     l,
	})

	event := dynatrace.EventCreation{
		EventType:      dynatrace.EventTypeErrorEvent,
		Start:          makeTimestamp(),
		TimeoutMinutes: 15,
		Source:         "Go",
		AttachRules: dynatrace.PushEventAttachRules{
			EntityIds: []string{"CUSTOM_DEVICE-D3692A3DBB1B6419"},
		},
		CustomProperties: map[string]string{
			"Origin": "Go",
		},
		Description: "Sample description",
		Title:       "Sample title",
	}
	e, _, err := c.Events.Create(event)
	if err != nil {
		panic(err)
	}

	fmt.Println(prettyPrint(e))

}

func main() {
	createEvent()

}
