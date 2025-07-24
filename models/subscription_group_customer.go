// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// SubscriptionGroupCustomer represents a SubscriptionGroupCustomer struct.
type SubscriptionGroupCustomer struct {
    FirstName            *string                `json:"first_name,omitempty"`
    LastName             *string                `json:"last_name,omitempty"`
    Organization         *string                `json:"organization,omitempty"`
    Email                *string                `json:"email,omitempty"`
    Reference            *string                `json:"reference,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionGroupCustomer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroupCustomer) String() string {
    return fmt.Sprintf(
    	"SubscriptionGroupCustomer[FirstName=%v, LastName=%v, Organization=%v, Email=%v, Reference=%v, AdditionalProperties=%v]",
    	s.FirstName, s.LastName, s.Organization, s.Email, s.Reference, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupCustomer.
// It customizes the JSON marshaling process for SubscriptionGroupCustomer objects.
func (s SubscriptionGroupCustomer) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "first_name", "last_name", "organization", "email", "reference"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupCustomer object to a map representation for JSON marshaling.
func (s SubscriptionGroupCustomer) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.FirstName != nil {
        structMap["first_name"] = s.FirstName
    }
    if s.LastName != nil {
        structMap["last_name"] = s.LastName
    }
    if s.Organization != nil {
        structMap["organization"] = s.Organization
    }
    if s.Email != nil {
        structMap["email"] = s.Email
    }
    if s.Reference != nil {
        structMap["reference"] = s.Reference
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupCustomer.
// It customizes the JSON unmarshaling process for SubscriptionGroupCustomer objects.
func (s *SubscriptionGroupCustomer) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionGroupCustomer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "first_name", "last_name", "organization", "email", "reference")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.FirstName = temp.FirstName
    s.LastName = temp.LastName
    s.Organization = temp.Organization
    s.Email = temp.Email
    s.Reference = temp.Reference
    return nil
}

// tempSubscriptionGroupCustomer is a temporary struct used for validating the fields of SubscriptionGroupCustomer.
type tempSubscriptionGroupCustomer  struct {
    FirstName    *string `json:"first_name,omitempty"`
    LastName     *string `json:"last_name,omitempty"`
    Organization *string `json:"organization,omitempty"`
    Email        *string `json:"email,omitempty"`
    Reference    *string `json:"reference,omitempty"`
}
