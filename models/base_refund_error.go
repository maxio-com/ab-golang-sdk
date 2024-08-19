/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// BaseRefundError represents a BaseRefundError struct.
type BaseRefundError struct {
    Base                 []interface{}  `json:"base,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BaseRefundError.
// It customizes the JSON marshaling process for BaseRefundError objects.
func (b BaseRefundError) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BaseRefundError object to a map representation for JSON marshaling.
func (b BaseRefundError) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    if b.Base != nil {
        structMap["base"] = b.Base
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BaseRefundError.
// It customizes the JSON unmarshaling process for BaseRefundError objects.
func (b *BaseRefundError) UnmarshalJSON(input []byte) error {
    var temp tempBaseRefundError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "base")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.Base = temp.Base
    return nil
}

// tempBaseRefundError is a temporary struct used for validating the fields of BaseRefundError.
type tempBaseRefundError  struct {
    Base []interface{} `json:"base,omitempty"`
}
