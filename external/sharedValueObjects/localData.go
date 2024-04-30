package sharedValueObjects

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

// Define a generic struct LocalData
type LocalData struct {
	Name        string
	Slug        string
	Description string
}

// Implement MarshalJSON method for LocalData
func (ld *LocalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"name":        ld.Name,
		"slug":        ld.Slug,
		"description": ld.Description,
	})
}

// Implement UnmarshalJSON method for LocalData
func (ld *LocalData) UnmarshalJSON(data []byte) error {
	var lData struct {
		Name        string `json:"name" validate:"required"`
		Slug        string `json:"slug" validate:"required"`
		Description string `json:"description" validate:"required"`
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
	ld.Slug = lData.Slug
	ld.Description = lData.Description
	return nil
}
