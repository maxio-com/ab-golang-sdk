/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// BulkComponentsPricePointAssignment represents a BulkComponentsPricePointAssignment struct.
type BulkComponentsPricePointAssignment struct {
    Components           []ComponentPricePointAssignment `json:"components,omitempty"`
    AdditionalProperties map[string]interface{}          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BulkComponentsPricePointAssignment.
// It customizes the JSON marshaling process for BulkComponentsPricePointAssignment objects.
func (b BulkComponentsPricePointAssignment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "components"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BulkComponentsPricePointAssignment object to a map representation for JSON marshaling.
func (b BulkComponentsPricePointAssignment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
    if b.Components != nil {
        structMap["components"] = b.Components
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkComponentsPricePointAssignment.
// It customizes the JSON unmarshaling process for BulkComponentsPricePointAssignment objects.
func (b *BulkComponentsPricePointAssignment) UnmarshalJSON(input []byte) error {
    var temp tempBulkComponentsPricePointAssignment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "components")
    if err != nil {
    	return err
    }
    b.AdditionalProperties = additionalProperties
    
    b.Components = temp.Components
    return nil
}

// tempBulkComponentsPricePointAssignment is a temporary struct used for validating the fields of BulkComponentsPricePointAssignment.
type tempBulkComponentsPricePointAssignment  struct {
    Components []ComponentPricePointAssignment `json:"components,omitempty"`
}
