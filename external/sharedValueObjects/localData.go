package sharedValueObjects

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

// Define a generic struct LocalData
type LocalData struct {
	Name           string
	Description    string
	VariantMessage string
}

// Implement MarshalJSON method for LocalData
func (ld *LocalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"name":            ld.Name,
		"description":     ld.Description,
		"variant_message": ld.VariantMessage,
	})
}

// Implement UnmarshalJSON method for LocalData
func (ld *LocalData) UnmarshalJSON(data []byte) error {
	var lData struct {
		Name           string `json:"name" validate:"required"`
		Description    string `json:"description"`
		VariantMessage string `json:"variant_message"`
	}

	if err := json.Unmarshal(data, &lData); err != nil {
		return err
	}

	validate := validator.New()
	err := validate.Struct(lData)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}

	ld.Name = lData.Name
	ld.Description = lData.Description
	ld.VariantMessage = lData.VariantMessage
	return nil
}
