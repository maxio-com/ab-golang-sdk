// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// CreateMetafield represents a CreateMetafield struct.
type CreateMetafield struct {
	Name *string `json:"name,omitempty"`
	// Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings.
	Scope *MetafieldScope `json:"scope,omitempty"`
	// Indicates how data should be added to the metafield. For example, a text type is just a string, so a given metafield of this type can have any value attached. On the other hand, dropdown and radio have a set of allowed values that can be input, and appear differently on a Public Signup Page. Defaults to 'text'
	InputType *MetafieldInput `json:"input_type,omitempty"`
	// Only applicable when input_type is radio or dropdown. Empty strings will not be submitted.
	Enum                 []string               `json:"enum,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateMetafield,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateMetafield) String() string {
	return fmt.Sprintf(
		"CreateMetafield[Name=%v, Scope=%v, InputType=%v, Enum=%v, AdditionalProperties=%v]",
		c.Name, c.Scope, c.InputType, c.Enum, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateMetafield.
// It customizes the JSON marshaling process for CreateMetafield objects.
func (c CreateMetafield) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"name", "scope", "input_type", "enum"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreateMetafield object to a map representation for JSON marshaling.
func (c CreateMetafield) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	if c.Name != nil {
		structMap["name"] = c.Name
	}
	if c.Scope != nil {
		structMap["scope"] = c.Scope.toMap()
	}
	if c.InputType != nil {
		structMap["input_type"] = c.InputType
	}
	if c.Enum != nil {
		structMap["enum"] = c.Enum
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMetafield.
// It customizes the JSON unmarshaling process for CreateMetafield objects.
func (c *CreateMetafield) UnmarshalJSON(input []byte) error {
	var temp tempCreateMetafield
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "scope", "input_type", "enum")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.Name = temp.Name
	c.Scope = temp.Scope
	c.InputType = temp.InputType
	c.Enum = temp.Enum
	return nil
}

// tempCreateMetafield is a temporary struct used for validating the fields of CreateMetafield.
type tempCreateMetafield struct {
	Name      *string         `json:"name,omitempty"`
	Scope     *MetafieldScope `json:"scope,omitempty"`
	InputType *MetafieldInput `json:"input_type,omitempty"`
	Enum      []string        `json:"enum,omitempty"`
}
