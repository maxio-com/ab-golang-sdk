// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// UpdateSubscriptionComponent represents a UpdateSubscriptionComponent struct.
type UpdateSubscriptionComponent struct {
	ComponentId *int `json:"component_id,omitempty"`
	// Create or update custom pricing unique to the subscription. Used in place of `price_point_id`.
	CustomPrice          *ComponentCustomPrice  `json:"custom_price,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateSubscriptionComponent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateSubscriptionComponent) String() string {
	return fmt.Sprintf(
		"UpdateSubscriptionComponent[ComponentId=%v, CustomPrice=%v, AdditionalProperties=%v]",
		u.ComponentId, u.CustomPrice, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionComponent.
// It customizes the JSON marshaling process for UpdateSubscriptionComponent objects.
func (u UpdateSubscriptionComponent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"component_id", "custom_price"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionComponent object to a map representation for JSON marshaling.
func (u UpdateSubscriptionComponent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.ComponentId != nil {
		structMap["component_id"] = u.ComponentId
	}
	if u.CustomPrice != nil {
		structMap["custom_price"] = u.CustomPrice.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionComponent.
// It customizes the JSON unmarshaling process for UpdateSubscriptionComponent objects.
func (u *UpdateSubscriptionComponent) UnmarshalJSON(input []byte) error {
	var temp tempUpdateSubscriptionComponent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "component_id", "custom_price")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.ComponentId = temp.ComponentId
	u.CustomPrice = temp.CustomPrice
	return nil
}

// tempUpdateSubscriptionComponent is a temporary struct used for validating the fields of UpdateSubscriptionComponent.
type tempUpdateSubscriptionComponent struct {
	ComponentId *int                  `json:"component_id,omitempty"`
	CustomPrice *ComponentCustomPrice `json:"custom_price,omitempty"`
}
