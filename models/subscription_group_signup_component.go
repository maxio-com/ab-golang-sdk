package models

import (
	"encoding/json"
)

// SubscriptionGroupSignupComponent represents a SubscriptionGroupSignupComponent struct.
type SubscriptionGroupSignupComponent struct {
	// Required if passing any component to `components` attribute.
	ComponentId       *SubscriptionGroupSignupComponentComponentId       `json:"component_id,omitempty"`
	AllocatedQuantity *SubscriptionGroupSignupComponentAllocatedQuantity `json:"allocated_quantity,omitempty"`
	UnitBalance       *SubscriptionGroupSignupComponentUnitBalance       `json:"unit_balance,omitempty"`
	PricePointId      *SubscriptionGroupSignupComponentPricePointId      `json:"price_point_id,omitempty"`
	// Used in place of `price_point_id` to define a custom price point unique to the subscription. You still need to provide `component_id`.
	CustomPrice *SubscriptionGroupComponentCustomPrice `json:"custom_price,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupComponent.
// It customizes the JSON marshaling process for SubscriptionGroupSignupComponent objects.
func (s *SubscriptionGroupSignupComponent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupComponent object to a map representation for JSON marshaling.
func (s *SubscriptionGroupSignupComponent) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.ComponentId != nil {
		structMap["component_id"] = s.ComponentId.toMap()
	}
	if s.AllocatedQuantity != nil {
		structMap["allocated_quantity"] = s.AllocatedQuantity.toMap()
	}
	if s.UnitBalance != nil {
		structMap["unit_balance"] = s.UnitBalance.toMap()
	}
	if s.PricePointId != nil {
		structMap["price_point_id"] = s.PricePointId.toMap()
	}
	if s.CustomPrice != nil {
		structMap["custom_price"] = s.CustomPrice.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupComponent.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupComponent objects.
func (s *SubscriptionGroupSignupComponent) UnmarshalJSON(input []byte) error {
	var temp subscriptionGroupSignupComponent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	s.ComponentId = temp.ComponentId
	s.AllocatedQuantity = temp.AllocatedQuantity
	s.UnitBalance = temp.UnitBalance
	s.PricePointId = temp.PricePointId
	s.CustomPrice = temp.CustomPrice
	return nil
}

// TODO
type subscriptionGroupSignupComponent struct {
	ComponentId       *SubscriptionGroupSignupComponentComponentId       `json:"component_id,omitempty"`
	AllocatedQuantity *SubscriptionGroupSignupComponentAllocatedQuantity `json:"allocated_quantity,omitempty"`
	UnitBalance       *SubscriptionGroupSignupComponentUnitBalance       `json:"unit_balance,omitempty"`
	PricePointId      *SubscriptionGroupSignupComponentPricePointId      `json:"price_point_id,omitempty"`
	CustomPrice       *SubscriptionGroupComponentCustomPrice             `json:"custom_price,omitempty"`
}
