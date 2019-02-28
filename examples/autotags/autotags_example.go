package main

import (
	"fmt"
	dynatrace "github.com/dyladan/dynatrace-go-client/api"
)

var c = dynatrace.New(dynatrace.Config{
	APIKey:  "MyAPIKey",
	BaseURL: "https://my.tenant.url/",
})

func createNewAutoTag() {

	condition := dynatrace.AutoTagRuleCondition{
		Key: dynatrace.AutoTagRuleConditionKey{
			Attribute: "CUSTOM_DEVICE_NAME",
		},

		ComparisonInfo: dynatrace.AutoTagRuleConditionComparisonInfo{
			Type:          "STRING",
			Value:         "Juniper",
			CaseSensitive: false,
			Negate:        false,
			Operator:      "BEGINS_WITH",
		},
	}

	rules := []dynatrace.AutoTagRule{
		{
			Type:        dynatrace.CustomDevice,
			Enabled:     true,
			ValueFormat: "Test {CustomDevice:DetectedName}",
			Conditions:  []dynatrace.AutoTagRuleCondition{condition},
		},
	}

	newAutoTag := dynatrace.AutoTag{
		Description: "Sample AutoTag",
		Name:        "autotag_test",
		Rules:       rules,
	}

	autoTag, _, err := c.AutoTags.Create(newAutoTag)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New AutoTag: %+v\n", autoTag)
}

func getAutoTagDetails() {

	autoTags, _, _ := c.AutoTags.GetAll()

	for i, autoTag := range autoTags {
		autoTagDetail, _, _ := c.AutoTags.Get(autoTag.ID, false)
		fmt.Printf("%d: %+v\n", i, autoTagDetail)
	}
}

func main() {
	createNewAutoTag()
	getAutoTagDetails()
}
