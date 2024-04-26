package sharedValueObjects

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMetaData_MarshalJSON(t *testing.T) {
	// Create a sample MetaData instance
	md := MetaData{
		MetaTitle:       "Page Title",
		MetaDescription: "This is a page description",
		MetaKeywords:    []string{"keyword1", "keyword2", "keyword3"},
	}

	// Marshal MetaData to JSON
	expectedJSON := `{"meta_description":"This is a page description","meta_keywords":["keyword1","keyword2","keyword3"],"meta_title":"Page Title"}`
	actualJSON, err := json.Marshal(&md)
	if err != nil {
		t.Errorf("MarshalJSON error: %v", err)
	}

	// Compare the marshaled JSON with the expected JSON
	if string(actualJSON) != expectedJSON {
		t.Errorf("Unexpected JSON result.\nExpected: %s\nActual: %s", expectedJSON, string(actualJSON))
	}
}

func TestMetaData_UnmarshalJSON(t *testing.T) {
	// Define a sample JSON payload
	inputJSON := []byte(`{"meta_title":"Page Title","meta_description":"This is a page description","meta_keywords":["keyword1","keyword2","keyword3"]}`)

	// Create an empty MetaData instance
	var md MetaData

	// Unmarshal JSON into MetaData
	err := json.Unmarshal(inputJSON, &md)
	if err != nil {
		t.Errorf("UnmarshalJSON error: %v", err)
	}

	// Define the expected MetaData instance after unmarshaling
	expectedMD := MetaData{
		MetaTitle:       "Page Title",
		MetaDescription: "This is a page description",
		MetaKeywords:    []string{"keyword1", "keyword2", "keyword3"},
	}

	// Compare the unmarshaled MetaData with the expected MetaData
	if !reflect.DeepEqual(md, expectedMD) {
		t.Errorf("Unexpected MetaData result.\nExpected: %+v\nActual: %+v", expectedMD, md)
	}
}
