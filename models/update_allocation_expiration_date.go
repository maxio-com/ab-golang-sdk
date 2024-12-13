/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// UpdateAllocationExpirationDate represents a UpdateAllocationExpirationDate struct.
type UpdateAllocationExpirationDate struct {
    Allocation           *AllocationExpirationDate `json:"allocation,omitempty"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateAllocationExpirationDate.
// It customizes the JSON marshaling process for UpdateAllocationExpirationDate objects.
func (u UpdateAllocationExpirationDate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "allocation"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateAllocationExpirationDate object to a map representation for JSON marshaling.
func (u UpdateAllocationExpirationDate) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Allocation != nil {
        structMap["allocation"] = u.Allocation.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateAllocationExpirationDate.
// It customizes the JSON unmarshaling process for UpdateAllocationExpirationDate objects.
func (u *UpdateAllocationExpirationDate) UnmarshalJSON(input []byte) error {
    var temp tempUpdateAllocationExpirationDate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "allocation")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Allocation = temp.Allocation
    return nil
}

// tempUpdateAllocationExpirationDate is a temporary struct used for validating the fields of UpdateAllocationExpirationDate.
type tempUpdateAllocationExpirationDate  struct {
    Allocation *AllocationExpirationDate `json:"allocation,omitempty"`
}
