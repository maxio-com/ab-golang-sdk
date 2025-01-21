/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// SubscriptionGroup represents a SubscriptionGroup struct.
type SubscriptionGroup struct {
    CustomerId              *int                             `json:"customer_id,omitempty"`
    PaymentProfile          *SubscriptionGroupPaymentProfile `json:"payment_profile,omitempty"`
    // The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`.
    PaymentCollectionMethod *CollectionMethod                `json:"payment_collection_method,omitempty"`
    SubscriptionIds         []int                            `json:"subscription_ids,omitempty"`
    CreatedAt               *time.Time                       `json:"created_at,omitempty"`
    AdditionalProperties    map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionGroup,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroup) String() string {
    return fmt.Sprintf(
    	"SubscriptionGroup[CustomerId=%v, PaymentProfile=%v, PaymentCollectionMethod=%v, SubscriptionIds=%v, CreatedAt=%v, AdditionalProperties=%v]",
    	s.CustomerId, s.PaymentProfile, s.PaymentCollectionMethod, s.SubscriptionIds, s.CreatedAt, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroup.
// It customizes the JSON marshaling process for SubscriptionGroup objects.
func (s SubscriptionGroup) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "customer_id", "payment_profile", "payment_collection_method", "subscription_ids", "created_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroup object to a map representation for JSON marshaling.
func (s SubscriptionGroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CustomerId != nil {
        structMap["customer_id"] = s.CustomerId
    }
    if s.PaymentProfile != nil {
        structMap["payment_profile"] = s.PaymentProfile.toMap()
    }
    if s.PaymentCollectionMethod != nil {
        structMap["payment_collection_method"] = s.PaymentCollectionMethod
    }
    if s.SubscriptionIds != nil {
        structMap["subscription_ids"] = s.SubscriptionIds
    }
    if s.CreatedAt != nil {
        structMap["created_at"] = s.CreatedAt.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroup.
// It customizes the JSON unmarshaling process for SubscriptionGroup objects.
func (s *SubscriptionGroup) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionGroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "customer_id", "payment_profile", "payment_collection_method", "subscription_ids", "created_at")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.CustomerId = temp.CustomerId
    s.PaymentProfile = temp.PaymentProfile
    s.PaymentCollectionMethod = temp.PaymentCollectionMethod
    s.SubscriptionIds = temp.SubscriptionIds
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        s.CreatedAt = &CreatedAtVal
    }
    return nil
}

// tempSubscriptionGroup is a temporary struct used for validating the fields of SubscriptionGroup.
type tempSubscriptionGroup  struct {
    CustomerId              *int                             `json:"customer_id,omitempty"`
    PaymentProfile          *SubscriptionGroupPaymentProfile `json:"payment_profile,omitempty"`
    PaymentCollectionMethod *CollectionMethod                `json:"payment_collection_method,omitempty"`
    SubscriptionIds         []int                            `json:"subscription_ids,omitempty"`
    CreatedAt               *string                          `json:"created_at,omitempty"`
}
