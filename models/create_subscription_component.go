// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// CreateSubscriptionComponent represents a CreateSubscriptionComponent struct.
type CreateSubscriptionComponent struct {
	ComponentId *CreateSubscriptionComponentComponentId `json:"component_id,omitempty"`
	// Used for on/off components only.
	Enabled *bool `json:"enabled,omitempty"`
	// Used for metered and events based components.
	UnitBalance *int `json:"unit_balance,omitempty"`
	// Used for quantity based components.
	AllocatedQuantity *CreateSubscriptionComponentAllocatedQuantity `json:"allocated_quantity,omitempty"`
	// Deprecated. Use `allocated_quantity` instead.
	Quantity     *int                                     `json:"quantity,omitempty"` // Deprecated
	PricePointId *CreateSubscriptionComponentPricePointId `json:"price_point_id,omitempty"`
	// Create or update custom pricing unique to the subscription. Used in place of `price_point_id`.
	CustomPrice          *ComponentCustomPrice  `json:"custom_price,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateSubscriptionComponent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateSubscriptionComponent) String() string {
	return fmt.Sprintf(
		"CreateSubscriptionComponent[ComponentId=%v, Enabled=%v, UnitBalance=%v, AllocatedQuantity=%v, Quantity=%v, PricePointId=%v, CustomPrice=%v, AdditionalProperties=%v]",
		c.ComponentId, c.Enabled, c.UnitBalance, c.AllocatedQuantity, c.Quantity, c.PricePointId, c.CustomPrice, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionComponent.
// It customizes the JSON marshaling process for CreateSubscriptionComponent objects.
func (c CreateSubscriptionComponent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"component_id", "enabled", "unit_balance", "allocated_quantity", "quantity", "price_point_id", "custom_price"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionComponent object to a map representation for JSON marshaling.
func (c CreateSubscriptionComponent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	if c.ComponentId != nil {
		structMap["component_id"] = c.ComponentId.toMap()
	}
	if c.Enabled != nil {
		structMap["enabled"] = c.Enabled
	}
	if c.UnitBalance != nil {
		structMap["unit_balance"] = c.UnitBalance
	}
	if c.AllocatedQuantity != nil {
		structMap["allocated_quantity"] = c.AllocatedQuantity.toMap()
	}
	if c.Quantity != nil {
		structMap["quantity"] = c.Quantity
	}
	if c.PricePointId != nil {
		structMap["price_point_id"] = c.PricePointId.toMap()
	}
	if c.CustomPrice != nil {
		structMap["custom_price"] = c.CustomPrice.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionComponent.
// It customizes the JSON unmarshaling process for CreateSubscriptionComponent objects.
func (c *CreateSubscriptionComponent) UnmarshalJSON(input []byte) error {
	var temp tempCreateSubscriptionComponent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "component_id", "enabled", "unit_balance", "allocated_quantity", "quantity", "price_point_id", "custom_price")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.ComponentId = temp.ComponentId
	c.Enabled = temp.Enabled
	c.UnitBalance = temp.UnitBalance
	c.AllocatedQuantity = temp.AllocatedQuantity
	c.Quantity = temp.Quantity
	c.PricePointId = temp.PricePointId
	c.CustomPrice = temp.CustomPrice
	return nil
}

// tempCreateSubscriptionComponent is a temporary struct used for validating the fields of CreateSubscriptionComponent.
type tempCreateSubscriptionComponent struct {
	ComponentId       *CreateSubscriptionComponentComponentId       `json:"component_id,omitempty"`
	Enabled           *bool                                         `json:"enabled,omitempty"`
	UnitBalance       *int                                          `json:"unit_balance,omitempty"`
	AllocatedQuantity *CreateSubscriptionComponentAllocatedQuantity `json:"allocated_quantity,omitempty"`
	Quantity          *int                                          `json:"quantity,omitempty"`
	PricePointId      *CreateSubscriptionComponentPricePointId      `json:"price_point_id,omitempty"`
	CustomPrice       *ComponentCustomPrice                         `json:"custom_price,omitempty"`
}
