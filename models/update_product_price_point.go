// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// UpdateProductPricePoint represents a UpdateProductPricePoint struct.
type UpdateProductPricePoint struct {
	Handle               *string                `json:"handle,omitempty"`
	PriceInCents         *int64                 `json:"price_in_cents,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateProductPricePoint,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateProductPricePoint) String() string {
	return fmt.Sprintf(
		"UpdateProductPricePoint[Handle=%v, PriceInCents=%v, AdditionalProperties=%v]",
		u.Handle, u.PriceInCents, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateProductPricePoint.
// It customizes the JSON marshaling process for UpdateProductPricePoint objects.
func (u UpdateProductPricePoint) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"handle", "price_in_cents"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateProductPricePoint object to a map representation for JSON marshaling.
func (u UpdateProductPricePoint) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Handle != nil {
		structMap["handle"] = u.Handle
	}
	if u.PriceInCents != nil {
		structMap["price_in_cents"] = u.PriceInCents
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateProductPricePoint.
// It customizes the JSON unmarshaling process for UpdateProductPricePoint objects.
func (u *UpdateProductPricePoint) UnmarshalJSON(input []byte) error {
	var temp tempUpdateProductPricePoint
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "handle", "price_in_cents")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.Handle = temp.Handle
	u.PriceInCents = temp.PriceInCents
	return nil
}

// tempUpdateProductPricePoint is a temporary struct used for validating the fields of UpdateProductPricePoint.
type tempUpdateProductPricePoint struct {
	Handle       *string `json:"handle,omitempty"`
	PriceInCents *int64  `json:"price_in_cents,omitempty"`
}
