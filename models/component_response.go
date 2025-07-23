// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ComponentResponse represents a ComponentResponse struct.
type ComponentResponse struct {
	Component            Component              `json:"component"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ComponentResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ComponentResponse) String() string {
	return fmt.Sprintf(
		"ComponentResponse[Component=%v, AdditionalProperties=%v]",
		c.Component, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ComponentResponse.
// It customizes the JSON marshaling process for ComponentResponse objects.
func (c ComponentResponse) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"component"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the ComponentResponse object to a map representation for JSON marshaling.
func (c ComponentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	structMap["component"] = c.Component.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentResponse.
// It customizes the JSON unmarshaling process for ComponentResponse objects.
func (c *ComponentResponse) UnmarshalJSON(input []byte) error {
	var temp tempComponentResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "component")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.Component = *temp.Component
	return nil
}

// tempComponentResponse is a temporary struct used for validating the fields of ComponentResponse.
type tempComponentResponse struct {
	Component *Component `json:"component"`
}

func (c *tempComponentResponse) validate() error {
	var errs []string
	if c.Component == nil {
		errs = append(errs, "required field `component` is missing for type `Component Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
