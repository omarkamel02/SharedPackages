package sharedValueObjects

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestLocalData_MarshalJSON(t *testing.T) {
	// Create a sample LocalData instance
	ld := LocalData{
		Name:        "Example",
		Slug:        "example",
		Description: "This is an example",
	}

	// Marshal LocalData to JSON
	expectedJSON := `{"description":"This is an example","name":"Example","slug":"example"}`
	actualJSON, err := json.Marshal(&ld)
	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}

	// Compare the marshaled JSON with the expected JSON
	if !reflect.DeepEqual(string(actualJSON), expectedJSON) {
		t.Errorf("Unexpected JSON result.\nExpected: %s\nActual: %s", expectedJSON, string(actualJSON))
	}
}

func TestLocalData_UnmarshalJSON(t *testing.T) {
	// Define a sample JSON payload
	inputJSON := []byte(`{"name":"Test","slug":"test","description":"This is a test"}`)

	// Create an empty LocalData instance
	var ld LocalData

	// Unmarshal JSON into LocalData
	err := json.Unmarshal(inputJSON, &ld)
	if err != nil {
		t.Errorf("UnmarshalJSON error: %v", err)
	}

	// Define the expected LocalData instance after unmarshaling
	expectedLD := LocalData{
		Name:        "Test",
		Slug:        "test",
		Description: "This is a test",
	}

	// Compare the unmarshaled LocalData with the expected LocalData
	if !reflect.DeepEqual(ld, expectedLD) {
		t.Errorf("Unexpected LocalData result.\nExpected: %+v\nActual: %+v", expectedLD, ld)
	}
}
