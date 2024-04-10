package models

import (
    "encoding/json"
)

// SubscriptionComponentAllocationErrorItem represents a SubscriptionComponentAllocationErrorItem struct.
type SubscriptionComponentAllocationErrorItem struct {
    Kind                 *string        `json:"kind,omitempty"`
    Message              *string        `json:"message,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentAllocationErrorItem.
// It customizes the JSON marshaling process for SubscriptionComponentAllocationErrorItem objects.
func (s SubscriptionComponentAllocationErrorItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentAllocationErrorItem object to a map representation for JSON marshaling.
func (s SubscriptionComponentAllocationErrorItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Kind != nil {
        structMap["kind"] = s.Kind
    }
    if s.Message != nil {
        structMap["message"] = s.Message
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentAllocationErrorItem.
// It customizes the JSON unmarshaling process for SubscriptionComponentAllocationErrorItem objects.
func (s *SubscriptionComponentAllocationErrorItem) UnmarshalJSON(input []byte) error {
    var temp subscriptionComponentAllocationErrorItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "kind", "message")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Kind = temp.Kind
    s.Message = temp.Message
    return nil
}

// TODO
type subscriptionComponentAllocationErrorItem  struct {
    Kind    *string `json:"kind,omitempty"`
    Message *string `json:"message,omitempty"`
}
