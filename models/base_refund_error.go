// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// BaseRefundError represents a BaseRefundError struct.
type BaseRefundError struct {
    Base                 []interface{}          `json:"base,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for BaseRefundError,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b BaseRefundError) String() string {
    return fmt.Sprintf(
    	"BaseRefundError[Base=%v, AdditionalProperties=%v]",
    	b.Base, b.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for BaseRefundError.
// It customizes the JSON marshaling process for BaseRefundError objects.
func (b BaseRefundError) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "base"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BaseRefundError object to a map representation for JSON marshaling.
func (b BaseRefundError) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "base")
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
