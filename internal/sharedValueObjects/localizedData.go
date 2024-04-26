package sharedValueObjects

import (
	"encoding/json"
)

// Define a generic struct LocalizedData
type LocalizedData[T any] struct {
	Lang      string // Language
	LocalData T      // Localized data of any type
}

// Implement MarshalJSON method for LocalizedData
func (ld *LocalizedData[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{

		"lang": ld.Lang,
		"data": ld.LocalData,
	})
}

// Implement UnmarshalJSON method for LocalizedData
func (ld *LocalizedData[T]) UnmarshalJSON(data []byte) error {
	var _lData struct {
		Lang string `json:"lang"`
		Data T      `json:"data"`
	}

	if err := json.Unmarshal(data, &_lData); err != nil {
		return err
	}

	ld.Lang = _lData.Lang
	ld.LocalData = _lData.Data
	return nil
}
