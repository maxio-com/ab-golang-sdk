package models

import (
    "encoding/json"
)

// SubscriptionResponse represents a SubscriptionResponse struct.
type SubscriptionResponse struct {
    Subscription         *Subscription  `json:"subscription,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionResponse.
// It customizes the JSON marshaling process for SubscriptionResponse objects.
func (s SubscriptionResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionResponse object to a map representation for JSON marshaling.
func (s SubscriptionResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Subscription != nil {
        structMap["subscription"] = s.Subscription.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionResponse.
// It customizes the JSON unmarshaling process for SubscriptionResponse objects.
func (s *SubscriptionResponse) UnmarshalJSON(input []byte) error {
    var temp subscriptionResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "subscription")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Subscription = temp.Subscription
    return nil
}

// subscriptionResponse is a temporary struct used for validating the fields of SubscriptionResponse.
type subscriptionResponse  struct {
    Subscription *Subscription `json:"subscription,omitempty"`
}
