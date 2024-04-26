package sharedValueObjects

import "encoding/json"

// Define a generic struct MetaData
type MetaData struct {
	MetaTitle       string
	MetaDescription string
	MetaKeywords    []string
}

// Implement MarshalJSON method for MetaData
func (md *MetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{

		"meta_title":       md.MetaTitle,
		"meta_description": md.MetaDescription,
		"meta_keywords":    md.MetaKeywords,
	})
}

// Implement UnmarshalJSON method for MetaData
func (md *MetaData) UnmarshalJSON(data []byte) error {
	var mData struct {
		MetaTitle       string   `json:"meta_title"`
		MetaDescription string   `json:"meta_description"`
		MetaKeywords    []string `json:"meta_keywords"`
	}

	if err := json.Unmarshal(data, &mData); err != nil {
		return err
	}
	md.MetaTitle = mData.MetaTitle
	md.MetaDescription = mData.MetaDescription
	md.MetaKeywords = mData.MetaKeywords
	return nil
}
