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

// AllocationExpirationDate represents a AllocationExpirationDate struct.
type AllocationExpirationDate struct {
    ExpiresAt            *time.Time             `json:"expires_at,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AllocationExpirationDate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AllocationExpirationDate) String() string {
    return fmt.Sprintf(
    	"AllocationExpirationDate[ExpiresAt=%v, AdditionalProperties=%v]",
    	a.ExpiresAt, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AllocationExpirationDate.
// It customizes the JSON marshaling process for AllocationExpirationDate objects.
func (a AllocationExpirationDate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "expires_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AllocationExpirationDate object to a map representation for JSON marshaling.
func (a AllocationExpirationDate) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "expires_at")
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
