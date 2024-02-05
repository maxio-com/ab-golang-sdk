package models

import (
    "encoding/json"
)

// UpdateSubscriptionComponent represents a UpdateSubscriptionComponent struct.
type UpdateSubscriptionComponent struct {
    ComponentId *int                  `json:"component_id,omitempty"`
    // Create or update custom pricing unique to the subscription. Used in place of `price_point_id`.
    CustomPrice *ComponentCustomPrice `json:"custom_price,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionComponent.
// It customizes the JSON marshaling process for UpdateSubscriptionComponent objects.
func (u *UpdateSubscriptionComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionComponent object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionComponent) toMap() map[string]any {
    structMap := make(map[string]any)
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
    temp := &struct {
        ComponentId *int                  `json:"component_id,omitempty"`
        CustomPrice *ComponentCustomPrice `json:"custom_price,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.ComponentId = temp.ComponentId
    u.CustomPrice = temp.CustomPrice
    return nil
}
