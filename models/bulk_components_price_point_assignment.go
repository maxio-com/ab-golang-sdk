package models

import (
    "encoding/json"
)

// BulkComponentsPricePointAssignment represents a BulkComponentsPricePointAssignment struct.
type BulkComponentsPricePointAssignment struct {
    Components           []ComponentPricePointAssignment `json:"components,omitempty"`
    AdditionalProperties map[string]any                  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BulkComponentsPricePointAssignment.
// It customizes the JSON marshaling process for BulkComponentsPricePointAssignment objects.
func (b BulkComponentsPricePointAssignment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BulkComponentsPricePointAssignment object to a map representation for JSON marshaling.
func (b BulkComponentsPricePointAssignment) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    if b.Components != nil {
        structMap["components"] = b.Components
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkComponentsPricePointAssignment.
// It customizes the JSON unmarshaling process for BulkComponentsPricePointAssignment objects.
func (b *BulkComponentsPricePointAssignment) UnmarshalJSON(input []byte) error {
    var temp bulkComponentsPricePointAssignment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "components")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.Components = temp.Components
    return nil
}

// bulkComponentsPricePointAssignment is a temporary struct used for validating the fields of BulkComponentsPricePointAssignment.
type bulkComponentsPricePointAssignment  struct {
    Components []ComponentPricePointAssignment `json:"components,omitempty"`
}
