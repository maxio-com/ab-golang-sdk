/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "log"
    "time"
)

// AllocationExpirationDate represents a AllocationExpirationDate struct.
type AllocationExpirationDate struct {
    ExpiresAt            *time.Time     `json:"expires_at,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AllocationExpirationDate.
// It customizes the JSON marshaling process for AllocationExpirationDate objects.
func (a AllocationExpirationDate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AllocationExpirationDate object to a map representation for JSON marshaling.
func (a AllocationExpirationDate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.ExpiresAt != nil {
        structMap["expires_at"] = a.ExpiresAt.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationExpirationDate.
// It customizes the JSON unmarshaling process for AllocationExpirationDate objects.
func (a *AllocationExpirationDate) UnmarshalJSON(input []byte) error {
    var temp tempAllocationExpirationDate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "expires_at")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    if temp.ExpiresAt != nil {
        ExpiresAtVal, err := time.Parse(time.RFC3339, *temp.ExpiresAt)
        if err != nil {
            log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
        }
        a.ExpiresAt = &ExpiresAtVal
    }
    return nil
}

// tempAllocationExpirationDate is a temporary struct used for validating the fields of AllocationExpirationDate.
type tempAllocationExpirationDate  struct {
    ExpiresAt *string `json:"expires_at,omitempty"`
}
