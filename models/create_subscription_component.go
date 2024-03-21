package models

import (
	"encoding/json"
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
	CustomPrice *ComponentCustomPrice `json:"custom_price,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionComponent.
// It customizes the JSON marshaling process for CreateSubscriptionComponent objects.
func (c *CreateSubscriptionComponent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionComponent object to a map representation for JSON marshaling.
func (c *CreateSubscriptionComponent) toMap() map[string]any {
	structMap := make(map[string]any)
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
	var temp createSubscriptionComponent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	c.ComponentId = temp.ComponentId
	c.Enabled = temp.Enabled
	c.UnitBalance = temp.UnitBalance
	c.AllocatedQuantity = temp.AllocatedQuantity
	c.Quantity = temp.Quantity
	c.PricePointId = temp.PricePointId
	c.CustomPrice = temp.CustomPrice
	return nil
}

// TODO
type createSubscriptionComponent struct {
	ComponentId       *CreateSubscriptionComponentComponentId       `json:"component_id,omitempty"`
	Enabled           *bool                                         `json:"enabled,omitempty"`
	UnitBalance       *int                                          `json:"unit_balance,omitempty"`
	AllocatedQuantity *CreateSubscriptionComponentAllocatedQuantity `json:"allocated_quantity,omitempty"`
	Quantity          *int                                          `json:"quantity,omitempty"`
	PricePointId      *CreateSubscriptionComponentPricePointId      `json:"price_point_id,omitempty"`
	CustomPrice       *ComponentCustomPrice                         `json:"custom_price,omitempty"`
}
