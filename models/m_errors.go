package models

import (
    "encoding/json"
)

// Errors represents a Errors struct.
type Errors struct {
    PerPage              []string       `json:"per_page,omitempty"`
    PricePoint           []string       `json:"price_point,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Errors.
// It customizes the JSON marshaling process for Errors objects.
func (e Errors) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the Errors object to a map representation for JSON marshaling.
func (e Errors) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    if e.PerPage != nil {
        structMap["per_page"] = e.PerPage
    }
    if e.PricePoint != nil {
        structMap["price_point"] = e.PricePoint
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Errors.
// It customizes the JSON unmarshaling process for Errors objects.
func (e *Errors) UnmarshalJSON(input []byte) error {
    var temp mErrors
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "per_page", "price_point")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.PerPage = temp.PerPage
    e.PricePoint = temp.PricePoint
    return nil
}

// mErrors is a temporary struct used for validating the fields of Errors.
type mErrors  struct {
    PerPage    []string `json:"per_page,omitempty"`
    PricePoint []string `json:"price_point,omitempty"`
}
