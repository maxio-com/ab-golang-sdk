// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// CreateProductFamily represents a CreateProductFamily struct.
type CreateProductFamily struct {
	Name                 string                 `json:"name"`
	Handle               Optional[string]       `json:"handle"`
	Description          Optional[string]       `json:"description"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateProductFamily,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateProductFamily) String() string {
	return fmt.Sprintf(
		"CreateProductFamily[Name=%v, Handle=%v, Description=%v, AdditionalProperties=%v]",
		c.Name, c.Handle, c.Description, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateProductFamily.
// It customizes the JSON marshaling process for CreateProductFamily objects.
func (c CreateProductFamily) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"name", "handle", "description"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreateProductFamily object to a map representation for JSON marshaling.
func (c CreateProductFamily) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	structMap["name"] = c.Name
	if c.Handle.IsValueSet() {
		if c.Handle.Value() != nil {
			structMap["handle"] = c.Handle.Value()
		} else {
			structMap["handle"] = nil
		}
	}
	if c.Description.IsValueSet() {
		if c.Description.Value() != nil {
			structMap["description"] = c.Description.Value()
		} else {
			structMap["description"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateProductFamily.
// It customizes the JSON unmarshaling process for CreateProductFamily objects.
func (c *CreateProductFamily) UnmarshalJSON(input []byte) error {
	var temp tempCreateProductFamily
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "handle", "description")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.Name = *temp.Name
	c.Handle = temp.Handle
	c.Description = temp.Description
	return nil
}

// tempCreateProductFamily is a temporary struct used for validating the fields of CreateProductFamily.
type tempCreateProductFamily struct {
	Name        *string          `json:"name"`
	Handle      Optional[string] `json:"handle"`
	Description Optional[string] `json:"description"`
}

func (c *tempCreateProductFamily) validate() error {
	var errs []string
	if c.Name == nil {
		errs = append(errs, "required field `name` is missing for type `Create Product Family`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
