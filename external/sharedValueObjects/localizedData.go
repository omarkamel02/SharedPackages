package sharedValueObjects

import (
	"encoding/json"

	"github.com/go-playground/validator"
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
		Lang string `json:"lang" validate:"required"`
		Data T      `json:"data" validate:"required"`
	}

	if err := json.Unmarshal(data, &_lData); err != nil {
		return err
	}
	validate := validator.New()
	err := validate.Struct(_lData)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}
	ld.Lang = _lData.Lang
	ld.LocalData = _lData.Data
	return nil
}
