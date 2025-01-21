/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// SubscriptionResponse represents a SubscriptionResponse struct.
type SubscriptionResponse struct {
    Subscription         *Subscription          `json:"subscription,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionResponse) String() string {
    return fmt.Sprintf(
    	"SubscriptionResponse[Subscription=%v, AdditionalProperties=%v]",
    	s.Subscription, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionResponse.
// It customizes the JSON marshaling process for SubscriptionResponse objects.
func (s SubscriptionResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "subscription"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionResponse object to a map representation for JSON marshaling.
func (s SubscriptionResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Subscription != nil {
        structMap["subscription"] = s.Subscription.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionResponse.
// It customizes the JSON unmarshaling process for SubscriptionResponse objects.
func (s *SubscriptionResponse) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscription")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Subscription = temp.Subscription
    return nil
}

// tempSubscriptionResponse is a temporary struct used for validating the fields of SubscriptionResponse.
type tempSubscriptionResponse  struct {
    Subscription *Subscription `json:"subscription,omitempty"`
}
