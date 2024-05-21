package sharedValueObjects

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

// Define a generic struct MetaData
type MetaData struct {
	MetaTitle       string
	MetaDescription string
	Slug            string
	MetaKeywords    []string
}

// Implement MarshalJSON method for MetaData
func (md *MetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{

		"meta_title":       md.MetaTitle,
		"meta_description": md.MetaDescription,
		"meta_keywords":    md.MetaKeywords,
		"slug":             md.Slug,
	})
}

// Implement UnmarshalJSON method for MetaData
func (md *MetaData) UnmarshalJSON(data []byte) error {
	var mData struct {
		MetaTitle       string   `json:"meta_title"`
		MetaDescription string   `json:"meta_description"`
		MetaKeywords    []string `json:"meta_keywords"`
		Slug            string   `json:"slug"`
	}

	if err := json.Unmarshal(data, &mData); err != nil {
		return err
	}

	validate := validator.New()
	err := validate.Struct(mData)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}

	md.MetaTitle = mData.MetaTitle
	md.MetaDescription = mData.MetaDescription
	md.MetaKeywords = mData.MetaKeywords
	md.Slug = mData.Slug
	return nil
}
