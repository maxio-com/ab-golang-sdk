/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// SubscriptionGroupSignupError represents a SubscriptionGroupSignupError struct.
type SubscriptionGroupSignupError struct {
    // Object that as key have subscription position in request subscriptions array and as value subscription errors object.
    Subscriptions        map[string]SubscriptionGroupSubscriptionError `json:"subscriptions,omitempty"`
    PayerReference       *string                                       `json:"payer_reference,omitempty"`
    Payer                *PayerError                                   `json:"payer,omitempty"`
    SubscriptionGroup    []string                                      `json:"subscription_group,omitempty"`
    PaymentProfileId     *string                                       `json:"payment_profile_id,omitempty"`
    PayerId              *string                                       `json:"payer_id,omitempty"`
    AdditionalProperties map[string]interface{}                        `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionGroupSignupError,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroupSignupError) String() string {
    return fmt.Sprintf(
    	"SubscriptionGroupSignupError[Subscriptions=%v, PayerReference=%v, Payer=%v, SubscriptionGroup=%v, PaymentProfileId=%v, PayerId=%v, AdditionalProperties=%v]",
    	s.Subscriptions, s.PayerReference, s.Payer, s.SubscriptionGroup, s.PaymentProfileId, s.PayerId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupError.
// It customizes the JSON marshaling process for SubscriptionGroupSignupError objects.
func (s SubscriptionGroupSignupError) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "subscriptions", "payer_reference", "payer", "subscription_group", "payment_profile_id", "payer_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupError object to a map representation for JSON marshaling.
func (s SubscriptionGroupSignupError) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Subscriptions != nil {
        structMap["subscriptions"] = s.Subscriptions
    }
    if s.PayerReference != nil {
        structMap["payer_reference"] = s.PayerReference
    }
    if s.Payer != nil {
        structMap["payer"] = s.Payer.toMap()
    }
    if s.SubscriptionGroup != nil {
        structMap["subscription_group"] = s.SubscriptionGroup
    }
    if s.PaymentProfileId != nil {
        structMap["payment_profile_id"] = s.PaymentProfileId
    }
    if s.PayerId != nil {
        structMap["payer_id"] = s.PayerId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupError.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupError objects.
func (s *SubscriptionGroupSignupError) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionGroupSignupError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscriptions", "payer_reference", "payer", "subscription_group", "payment_profile_id", "payer_id")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Subscriptions = temp.Subscriptions
    s.PayerReference = temp.PayerReference
    s.Payer = temp.Payer
    s.SubscriptionGroup = temp.SubscriptionGroup
    s.PaymentProfileId = temp.PaymentProfileId
    s.PayerId = temp.PayerId
    return nil
}

// tempSubscriptionGroupSignupError is a temporary struct used for validating the fields of SubscriptionGroupSignupError.
type tempSubscriptionGroupSignupError  struct {
    Subscriptions     map[string]SubscriptionGroupSubscriptionError `json:"subscriptions,omitempty"`
    PayerReference    *string                                       `json:"payer_reference,omitempty"`
    Payer             *PayerError                                   `json:"payer,omitempty"`
    SubscriptionGroup []string                                      `json:"subscription_group,omitempty"`
    PaymentProfileId  *string                                       `json:"payment_profile_id,omitempty"`
    PayerId           *string                                       `json:"payer_id,omitempty"`
}
