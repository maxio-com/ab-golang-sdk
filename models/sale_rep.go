// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// SaleRep represents a SaleRep struct.
type SaleRep struct {
    Id                   *int                   `json:"id,omitempty"`
    FullName             *string                `json:"full_name,omitempty"`
    SubscriptionsCount   *int                   `json:"subscriptions_count,omitempty"`
    TestMode             *bool                  `json:"test_mode,omitempty"`
    Subscriptions        []SaleRepSubscription  `json:"subscriptions,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SaleRep,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SaleRep) String() string {
    return fmt.Sprintf(
    	"SaleRep[Id=%v, FullName=%v, SubscriptionsCount=%v, TestMode=%v, Subscriptions=%v, AdditionalProperties=%v]",
    	s.Id, s.FullName, s.SubscriptionsCount, s.TestMode, s.Subscriptions, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SaleRep.
// It customizes the JSON marshaling process for SaleRep objects.
func (s SaleRep) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "id", "full_name", "subscriptions_count", "test_mode", "subscriptions"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SaleRep object to a map representation for JSON marshaling.
func (s SaleRep) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.FullName != nil {
        structMap["full_name"] = s.FullName
    }
    if s.SubscriptionsCount != nil {
        structMap["subscriptions_count"] = s.SubscriptionsCount
    }
    if s.TestMode != nil {
        structMap["test_mode"] = s.TestMode
    }
    if s.Subscriptions != nil {
        structMap["subscriptions"] = s.Subscriptions
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SaleRep.
// It customizes the JSON unmarshaling process for SaleRep objects.
func (s *SaleRep) UnmarshalJSON(input []byte) error {
    var temp tempSaleRep
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "full_name", "subscriptions_count", "test_mode", "subscriptions")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Id = temp.Id
    s.FullName = temp.FullName
    s.SubscriptionsCount = temp.SubscriptionsCount
    s.TestMode = temp.TestMode
    s.Subscriptions = temp.Subscriptions
    return nil
}

// tempSaleRep is a temporary struct used for validating the fields of SaleRep.
type tempSaleRep  struct {
    Id                 *int                  `json:"id,omitempty"`
    FullName           *string               `json:"full_name,omitempty"`
    SubscriptionsCount *int                  `json:"subscriptions_count,omitempty"`
    TestMode           *bool                 `json:"test_mode,omitempty"`
    Subscriptions      []SaleRepSubscription `json:"subscriptions,omitempty"`
}
