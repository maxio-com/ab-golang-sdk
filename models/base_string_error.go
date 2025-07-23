// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// BaseStringError represents a BaseStringError struct.
// The error is base if it is not directly associated with a single attribute.
type BaseStringError struct {
	Base                 []string               `json:"base,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for BaseStringError,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b BaseStringError) String() string {
	return fmt.Sprintf(
		"BaseStringError[Base=%v, AdditionalProperties=%v]",
		b.Base, b.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for BaseStringError.
// It customizes the JSON marshaling process for BaseStringError objects.
func (b BaseStringError) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(b.AdditionalProperties,
		"base"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(b.toMap())
}

// toMap converts the BaseStringError object to a map representation for JSON marshaling.
func (b BaseStringError) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, b.AdditionalProperties)
	if b.Base != nil {
		structMap["base"] = b.Base
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BaseStringError.
// It customizes the JSON unmarshaling process for BaseStringError objects.
func (b *BaseStringError) UnmarshalJSON(input []byte) error {
	var temp tempBaseStringError
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

// tempBaseStringError is a temporary struct used for validating the fields of BaseStringError.
type tempBaseStringError struct {
	Base []string `json:"base,omitempty"`
}
