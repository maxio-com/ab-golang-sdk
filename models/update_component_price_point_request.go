// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// UpdateComponentPricePointRequest represents a UpdateComponentPricePointRequest struct.
type UpdateComponentPricePointRequest struct {
	PricePoint           *UpdateComponentPricePoint `json:"price_point,omitempty"`
	AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateComponentPricePointRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateComponentPricePointRequest) String() string {
	return fmt.Sprintf(
		"UpdateComponentPricePointRequest[PricePoint=%v, AdditionalProperties=%v]",
		u.PricePoint, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateComponentPricePointRequest.
// It customizes the JSON marshaling process for UpdateComponentPricePointRequest objects.
func (u UpdateComponentPricePointRequest) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"price_point"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateComponentPricePointRequest object to a map representation for JSON marshaling.
func (u UpdateComponentPricePointRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.PricePoint != nil {
		structMap["price_point"] = u.PricePoint.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateComponentPricePointRequest.
// It customizes the JSON unmarshaling process for UpdateComponentPricePointRequest objects.
func (u *UpdateComponentPricePointRequest) UnmarshalJSON(input []byte) error {
	var temp tempUpdateComponentPricePointRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "price_point")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.PricePoint = temp.PricePoint
	return nil
}

// tempUpdateComponentPricePointRequest is a temporary struct used for validating the fields of UpdateComponentPricePointRequest.
type tempUpdateComponentPricePointRequest struct {
	PricePoint *UpdateComponentPricePoint `json:"price_point,omitempty"`
}
