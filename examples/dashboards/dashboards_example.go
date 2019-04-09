package main

import (
	"encoding/json"
	"fmt"
	dynatrace "github.com/dyladan/dynatrace-go-client/api"
	"log"
	"os"
)

var c = dynatrace.New(dynatrace.Config{
	APIKey:  os.Getenv("DT_API_KEY"),
	BaseURL: os.Getenv("DT_BASE_URL"),
})

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func getAllDashboards() {
	dashboards, rawResp, err := c.Dashboards.GetAll()
	if err != nil {
		log.Fatalf("Error getting dashboards : %s", err.Error())
	}
	fmt.Printf("Response: %d\n", rawResp.StatusCode())
	for _, dashboard := range dashboards {
		detailed, _, _ := c.Dashboards.Get(dashboard.ID)
		fmt.Println(prettyPrint(detailed))

	}

}

func main() {
	getAllDashboards()

}
