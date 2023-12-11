package models

import (
	"encoding/json"
)

// BulkComponentSPricePointAssignment represents a BulkComponentSPricePointAssignment struct.
type BulkComponentSPricePointAssignment struct {
	Components []ComponentSPricePointAssignment `json:"components,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for BulkComponentSPricePointAssignment.
// It customizes the JSON marshaling process for BulkComponentSPricePointAssignment objects.
func (b *BulkComponentSPricePointAssignment) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BulkComponentSPricePointAssignment object to a map representation for JSON marshaling.
func (b *BulkComponentSPricePointAssignment) toMap() map[string]any {
	structMap := make(map[string]any)
	if b.Components != nil {
		structMap["components"] = b.Components
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkComponentSPricePointAssignment.
// It customizes the JSON unmarshaling process for BulkComponentSPricePointAssignment objects.
func (b *BulkComponentSPricePointAssignment) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Components []ComponentSPricePointAssignment `json:"components,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.Components = temp.Components
	return nil
}
