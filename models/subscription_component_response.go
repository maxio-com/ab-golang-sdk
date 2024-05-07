package models

import (
    "encoding/json"
)

// SubscriptionComponentResponse represents a SubscriptionComponentResponse struct.
type SubscriptionComponentResponse struct {
    Component            *SubscriptionComponent `json:"component,omitempty"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentResponse.
// It customizes the JSON marshaling process for SubscriptionComponentResponse objects.
func (s SubscriptionComponentResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentResponse object to a map representation for JSON marshaling.
func (s SubscriptionComponentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Component != nil {
        structMap["component"] = s.Component.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentResponse.
// It customizes the JSON unmarshaling process for SubscriptionComponentResponse objects.
func (s *SubscriptionComponentResponse) UnmarshalJSON(input []byte) error {
    var temp subscriptionComponentResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "component")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Component = temp.Component
    return nil
}

// subscriptionComponentResponse is a temporary struct used for validating the fields of SubscriptionComponentResponse.
type subscriptionComponentResponse  struct {
    Component *SubscriptionComponent `json:"component,omitempty"`
}
