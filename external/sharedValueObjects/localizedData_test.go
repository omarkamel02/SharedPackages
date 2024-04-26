package sharedValueObjects

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestLocalizedData_MarshalJSON(t *testing.T) {
	// Create a sample LocalizedData instance with a string as LocalData
	ldString := LocalizedData[string]{Lang: "en", LocalData: "Hello"}

	// Marshal LocalizedData to JSON
	expectedJSONString := `{"data":"Hello","lang":"en"}`
	actualJSONString, err := json.Marshal(&ldString)
	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}

	// Compare the marshaled JSON with the expected JSON
	if string(actualJSONString) != expectedJSONString {
		t.Errorf("Unexpected JSON result.\nExpected: %s\nActual: %s", expectedJSONString, string(actualJSONString))
	}

	// Create a sample LocalizedData instance with an integer as LocalData
	ldInt := LocalizedData[int]{Lang: "fr", LocalData: 42}

	// Marshal LocalizedData to JSON
	expectedJSONInt := `{"data":42,"lang":"fr"}`
	actualJSONInt, err := json.Marshal(&ldInt)
	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}

	// Compare the marshaled JSON with the expected JSON
	if string(actualJSONInt) != expectedJSONInt {
		t.Errorf("Unexpected JSON result.\nExpected: %s\nActual: %s", expectedJSONInt, string(actualJSONInt))
	}
}

func TestLocalizedData_UnmarshalJSON(t *testing.T) {
	// Define sample JSON payloads for string and integer types
	inputJSONString := []byte(`{"lang":"en","data":"Hello"}`)
	inputJSONInt := []byte(`{"lang":"fr","data":42}`)

	// Create empty LocalizedData instances for string and integer types
	var ldString LocalizedData[string]
	var ldInt LocalizedData[int]

	// Unmarshal JSON into LocalizedData for string type
	err := json.Unmarshal(inputJSONString, &ldString)
	if err != nil {
		t.Errorf("UnmarshalJSON (string) error: %v", err)
	}

	// Define the expected LocalizedData instance for string type
	expectedLDString := LocalizedData[string]{Lang: "en", LocalData: "Hello"}

	// Compare the unmarshaled LocalizedData with the expected LocalizedData for string type
	if !reflect.DeepEqual(ldString, expectedLDString) {
		t.Errorf("Unexpected LocalizedData result (string).\nExpected: %+v\nActual: %+v", expectedLDString, ldString)
	}

	// Unmarshal JSON into LocalizedData for integer type
	err = json.Unmarshal(inputJSONInt, &ldInt)
	if err != nil {
		t.Errorf("UnmarshalJSON (int) error: %v", err)
	}

	// Define the expected LocalizedData instance for integer type
	expectedLDInt := LocalizedData[int]{Lang: "fr", LocalData: 42}

	// Compare the unmarshaled LocalizedData with the expected LocalizedData for integer type
	if !reflect.DeepEqual(ldInt, expectedLDInt) {
		t.Errorf("Unexpected LocalizedData result (int).\nExpected: %+v\nActual: %+v", expectedLDInt, ldInt)
	}
}
