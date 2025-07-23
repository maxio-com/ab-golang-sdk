// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// CreateMetafieldsRequest represents a CreateMetafieldsRequest struct.
type CreateMetafieldsRequest struct {
	Metafields           CreateMetafieldsRequestMetafields `json:"metafields"`
	AdditionalProperties map[string]interface{}            `json:"_"`
}

// String implements the fmt.Stringer interface for CreateMetafieldsRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateMetafieldsRequest) String() string {
	return fmt.Sprintf(
		"CreateMetafieldsRequest[Metafields=%v, AdditionalProperties=%v]",
		c.Metafields, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateMetafieldsRequest.
// It customizes the JSON marshaling process for CreateMetafieldsRequest objects.
func (c CreateMetafieldsRequest) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"metafields"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreateMetafieldsRequest object to a map representation for JSON marshaling.
func (c CreateMetafieldsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	structMap["metafields"] = c.Metafields.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMetafieldsRequest.
// It customizes the JSON unmarshaling process for CreateMetafieldsRequest objects.
func (c *CreateMetafieldsRequest) UnmarshalJSON(input []byte) error {
	var temp tempCreateMetafieldsRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "metafields")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.Metafields = *temp.Metafields
	return nil
}

// tempCreateMetafieldsRequest is a temporary struct used for validating the fields of CreateMetafieldsRequest.
type tempCreateMetafieldsRequest struct {
	Metafields *CreateMetafieldsRequestMetafields `json:"metafields"`
}

func (c *tempCreateMetafieldsRequest) validate() error {
	var errs []string
	if c.Metafields == nil {
		errs = append(errs, "required field `metafields` is missing for type `Create Metafields Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
