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
		detailed, _, err := c.Dashboards.Get(dashboard.ID)
		// fmt.Printf("Raw: %+v", raw)
		if err != nil {
			panic(err)
		}
		if detailed.ID == "8e8aebac-ade0-4134-850e-cd564f2fa125" {
			fmt.Println(prettyPrint(detailed))
		}
		// fmt.Println(detailed.ID, detailed.DashboardMetadata.Name)
	}

}

func getDashboard(id string) {
	_, raw, err := c.Dashboards.Get(id)
	if err != nil {
		panic(err)
	}
	fmt.Println(raw)

}

func createDashboard() {

	tiles := &[]dynatrace.Tile{
		{
			Name:     "CPU Usage",
			TileType: dynatrace.TileTypeCustomCharting,
			Bounds: &dynatrace.TileBounds{
				Height: 380,
				Width:  380,
				Top:    0,
			},
			Configured: true,
			CustomFilterConfig: &dynatrace.CustomFilterConfig{
				CustomName:  "HOSTS - CPU Usage",
				DefaultName: "Custom chart",
				Type:        "MIXED",
				ChartConfig: &dynatrace.CustomFilterChartConfig{
					Type: dynatrace.TIMESERIES,
					Series: &[]dynatrace.CustomFilterChartSeriesConfig{
						{
							Aggregation: dynatrace.AVG,
							Type:        dynatrace.LINE,
							EntityType:  "HOST",
							Metric:      "CTS_PLUGIN_ruxit.python.heartbeat:cpu_system",
							Dimensions:  []dynatrace.CustomFilterChartSeriesDimensionConfig{},
						},
					},
				},
				FiltersPerEntityType: map[string]map[string][]string{
					"REMOTE_PLUGIN": {
						"SPECIFIC_ENTITIES": {"CUSTOM_DEVICE-6385CC5F374FE680"},
					},
				},
			},
			TileFilter: &dynatrace.TileFilter{},
		},
	}

	dash := &dynatrace.Dashboard{
		Tiles: *tiles,
		DashboardMetadata: &dynatrace.DashboardMetadata{
			Name: "Dashboard from Go!",
			DashboardFilter: &dynatrace.DashboardFilter{
				Timeframe: "l_2_HOURS",
			},
		},
	}

	fmt.Println(prettyPrint(dash))

	dash, raw, err := c.Dashboards.Create(*dash)
	fmt.Println(raw)
	if err != nil {
		panic(err)
	}

}

func main() {
	createDashboard()

}
