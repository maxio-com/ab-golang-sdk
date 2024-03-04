package models

import (
    "encoding/json"
)

// BulkComponentsPricePointAssignment represents a BulkComponentsPricePointAssignment struct.
type BulkComponentsPricePointAssignment struct {
    Components []ComponentPricePointAssignment `json:"components,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for BulkComponentsPricePointAssignment.
// It customizes the JSON marshaling process for BulkComponentsPricePointAssignment objects.
func (b *BulkComponentsPricePointAssignment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BulkComponentsPricePointAssignment object to a map representation for JSON marshaling.
func (b *BulkComponentsPricePointAssignment) toMap() map[string]any {
    structMap := make(map[string]any)
    if b.Components != nil {
        structMap["components"] = b.Components
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkComponentsPricePointAssignment.
// It customizes the JSON unmarshaling process for BulkComponentsPricePointAssignment objects.
func (b *BulkComponentsPricePointAssignment) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Components []ComponentPricePointAssignment `json:"components,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    b.Components = temp.Components
    return nil
}
