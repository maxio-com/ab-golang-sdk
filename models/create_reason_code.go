// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// CreateReasonCode represents a CreateReasonCode struct.
type CreateReasonCode struct {
	// The unique identifier for the ReasonCode
	Code string `json:"code"`
	// The friendly summary of what the code signifies
	Description string `json:"description"`
	// The order that code appears in lists
	Position             *int                   `json:"position,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateReasonCode,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateReasonCode) String() string {
	return fmt.Sprintf(
		"CreateReasonCode[Code=%v, Description=%v, Position=%v, AdditionalProperties=%v]",
		c.Code, c.Description, c.Position, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateReasonCode.
// It customizes the JSON marshaling process for CreateReasonCode objects.
func (c CreateReasonCode) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"code", "description", "position"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreateReasonCode object to a map representation for JSON marshaling.
func (c CreateReasonCode) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	structMap["code"] = c.Code
	structMap["description"] = c.Description
	if c.Position != nil {
		structMap["position"] = c.Position
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateReasonCode.
// It customizes the JSON unmarshaling process for CreateReasonCode objects.
func (c *CreateReasonCode) UnmarshalJSON(input []byte) error {
	var temp tempCreateReasonCode
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "code", "description", "position")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.Code = *temp.Code
	c.Description = *temp.Description
	c.Position = temp.Position
	return nil
}

// tempCreateReasonCode is a temporary struct used for validating the fields of CreateReasonCode.
type tempCreateReasonCode struct {
	Code        *string `json:"code"`
	Description *string `json:"description"`
	Position    *int    `json:"position,omitempty"`
}

func (c *tempCreateReasonCode) validate() error {
	var errs []string
	if c.Code == nil {
		errs = append(errs, "required field `code` is missing for type `Create Reason Code`")
	}
	if c.Description == nil {
		errs = append(errs, "required field `description` is missing for type `Create Reason Code`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
