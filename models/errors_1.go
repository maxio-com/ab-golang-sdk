// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// Errors1 represents a Errors1 struct.
type Errors1 struct {
    PerPage              []string               `json:"per_page,omitempty"`
    PricePoint           []string               `json:"price_point,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Errors1,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e Errors1) String() string {
    return fmt.Sprintf(
    	"Errors1[PerPage=%v, PricePoint=%v, AdditionalProperties=%v]",
    	e.PerPage, e.PricePoint, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Errors1.
// It customizes the JSON marshaling process for Errors1 objects.
func (e Errors1) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "per_page", "price_point"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the Errors1 object to a map representation for JSON marshaling.
func (e Errors1) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    if e.PerPage != nil {
        structMap["per_page"] = e.PerPage
    }
    if e.PricePoint != nil {
        structMap["price_point"] = e.PricePoint
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Errors1.
// It customizes the JSON unmarshaling process for Errors1 objects.
func (e *Errors1) UnmarshalJSON(input []byte) error {
    var temp tempErrors1
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "per_page", "price_point")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.PerPage = temp.PerPage
    e.PricePoint = temp.PricePoint
    return nil
}

// tempErrors1 is a temporary struct used for validating the fields of Errors1.
type tempErrors1  struct {
    PerPage    []string `json:"per_page,omitempty"`
    PricePoint []string `json:"price_point,omitempty"`
}
