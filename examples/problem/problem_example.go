package main

import (
	"encoding/json"
	"fmt"
	dynatrace "github.com/dyladan/dynatrace-go-client/api"
	log "github.com/sirupsen/logrus"
	"os"
)

func prettyPrint(i interface{}) string {
	s, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		return fmt.Sprintf("Could not parse the object: %+v", i)
	}
	return string(s)
}

func listProblems() {

	l := log.New()
	l.Level = log.DebugLevel

	c := dynatrace.New(dynatrace.Config{
		APIKey:  os.Getenv("DT_API_KEY"),
		BaseURL: os.Getenv("DT_BASE_URL"),
		Log:     l,
	})

	fields := []string{"evidenceDetails"}
	problemSelector := "status(\"open\")"

	problems, _, err := c.Problem.List(fields, problemSelector, "", "")
	if err != nil {
		panic(err)
	}
	for _, problem := range problems {
		for _, evidenceDetails := range problem.EvidenceDetails.Details {
			fmt.Println(evidenceDetails.DisplayName)
		}
	}

}

func main() {
	listProblems()

}
