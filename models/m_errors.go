// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// Errors represents a Errors struct.
type Errors struct {
    PerPage              []string               `json:"per_page,omitempty"`
    PricePoint           []string               `json:"price_point,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Errors,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m Errors) String() string {
    return fmt.Sprintf(
    	"Errors[PerPage=%v, PricePoint=%v, AdditionalProperties=%v]",
    	m.PerPage, m.PricePoint, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Errors.
// It customizes the JSON marshaling process for Errors objects.
func (m Errors) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "per_page", "price_point"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the Errors object to a map representation for JSON marshaling.
func (m Errors) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.PerPage != nil {
        structMap["per_page"] = m.PerPage
    }
    if m.PricePoint != nil {
        structMap["price_point"] = m.PricePoint
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Errors.
// It customizes the JSON unmarshaling process for Errors objects.
func (m *Errors) UnmarshalJSON(input []byte) error {
    var temp tempErrors
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "per_page", "price_point")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.PerPage = temp.PerPage
    m.PricePoint = temp.PricePoint
    return nil
}

// tempErrors is a temporary struct used for validating the fields of Errors.
type tempErrors  struct {
    PerPage    []string `json:"per_page,omitempty"`
    PricePoint []string `json:"price_point,omitempty"`
}
