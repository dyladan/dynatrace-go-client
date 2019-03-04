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

	fmt.Println("Creating a new AutoTag...")

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
			Type:        dynatrace.AutoTagRuleTypeCustomDevice,
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

	autoTag, resp, _ := c.AutoTags.Create(newAutoTag)
	fmt.Println("API Response", resp)
	fmt.Printf("Created new AutoTag: %+v\n", autoTag)
}

func getAutoTagDetails() {
	fmt.Println("\nGetting the details of every AutoTag...")

	autoTags, _, _ := c.AutoTags.GetAll()

	for i, autoTag := range autoTags {
		autoTagDetail, _, _ := c.AutoTags.Get(autoTag.ID, false)
		fmt.Printf("AutoTag Detail: %d: %+v\n", i, autoTagDetail)
	}
}

func updateAutoTag() {

	fmt.Println("\nUpdating a single AutoTag...")

	autoTags, _, _ := c.AutoTags.GetAll()
	autoTag, _, _ := c.AutoTags.Get(autoTags[0].ID, false)

	autoTag.Rules[0].Conditions[0].ComparisonInfo.Value = "New Comparison"

	_, resp, err := c.AutoTags.Update(autoTag.ID, *autoTag)
	fmt.Printf("%+v\n", resp.StatusCode())
	fmt.Println("Error", err)

}

func deleteAutoTag() {

	fmt.Println("\nDeleting an AutoTag...")
	autoTags, _, _ := c.AutoTags.GetAll()

	resp, err := c.AutoTags.Delete(autoTags[0].ID)
	fmt.Println(resp.StatusCode())
	fmt.Println("Error", err)

}

func validadeUpdate() {
	fmt.Println("\nValidating update...")
	autoTags, _, _ := c.AutoTags.GetAll()

	autoTagToValidate, _, _ := c.AutoTags.Get(autoTags[0].ID, false)
	autoTagToValidate.Rules[0].Conditions[0].ComparisonInfo.Negate = true

	validated, resp, err := c.AutoTags.ValidateUpdate(autoTagToValidate.ID, *autoTagToValidate)
	fmt.Println("Validated", validated)
	fmt.Println("resp", resp.StatusCode())
	fmt.Println("err", err)

}

func validateCreate() {

	fmt.Println("\nValidating create...")
	autoTags, _, _ := c.AutoTags.GetAll()

	autoTagToValidate, _, _ := c.AutoTags.Get(autoTags[0].ID, false)
	autoTagToValidate.Rules[0].Conditions[0].ComparisonInfo.Negate = true

	validated, resp, err := c.AutoTags.ValidateCreate(*autoTagToValidate)
	fmt.Println("Validated", validated)
	fmt.Println("resp", resp.StatusCode())
	fmt.Println("err", err)

}

func main() {
	createNewAutoTag()
	getAutoTagDetails()
	validateCreate()
	validadeUpdate()
	updateAutoTag()
	deleteAutoTag()

}
