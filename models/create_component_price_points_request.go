// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// CreateComponentPricePointsRequest represents a CreateComponentPricePointsRequest struct.
type CreateComponentPricePointsRequest struct {
	PricePoints          []CreateComponentPricePointsRequestPricePoints `json:"price_points"`
	AdditionalProperties map[string]interface{}                         `json:"_"`
}

// String implements the fmt.Stringer interface for CreateComponentPricePointsRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateComponentPricePointsRequest) String() string {
	return fmt.Sprintf(
		"CreateComponentPricePointsRequest[PricePoints=%v, AdditionalProperties=%v]",
		c.PricePoints, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateComponentPricePointsRequest.
// It customizes the JSON marshaling process for CreateComponentPricePointsRequest objects.
func (c CreateComponentPricePointsRequest) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"price_points"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreateComponentPricePointsRequest object to a map representation for JSON marshaling.
func (c CreateComponentPricePointsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	structMap["price_points"] = c.PricePoints
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateComponentPricePointsRequest.
// It customizes the JSON unmarshaling process for CreateComponentPricePointsRequest objects.
func (c *CreateComponentPricePointsRequest) UnmarshalJSON(input []byte) error {
	var temp tempCreateComponentPricePointsRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "price_points")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.PricePoints = *temp.PricePoints
	return nil
}

// tempCreateComponentPricePointsRequest is a temporary struct used for validating the fields of CreateComponentPricePointsRequest.
type tempCreateComponentPricePointsRequest struct {
	PricePoints *[]CreateComponentPricePointsRequestPricePoints `json:"price_points"`
}

func (c *tempCreateComponentPricePointsRequest) validate() error {
	var errs []string
	if c.PricePoints == nil {
		errs = append(errs, "required field `price_points` is missing for type `Create Component Price Points Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
