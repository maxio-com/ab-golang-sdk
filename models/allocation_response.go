// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// AllocationResponse represents a AllocationResponse struct.
type AllocationResponse struct {
	Allocation           *Allocation            `json:"allocation,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AllocationResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AllocationResponse) String() string {
	return fmt.Sprintf(
		"AllocationResponse[Allocation=%v, AdditionalProperties=%v]",
		a.Allocation, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AllocationResponse.
// It customizes the JSON marshaling process for AllocationResponse objects.
func (a AllocationResponse) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"allocation"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AllocationResponse object to a map representation for JSON marshaling.
func (a AllocationResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Allocation != nil {
		structMap["allocation"] = a.Allocation.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationResponse.
// It customizes the JSON unmarshaling process for AllocationResponse objects.
func (a *AllocationResponse) UnmarshalJSON(input []byte) error {
	var temp tempAllocationResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "allocation")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.Allocation = temp.Allocation
	return nil
}

// tempAllocationResponse is a temporary struct used for validating the fields of AllocationResponse.
type tempAllocationResponse struct {
	Allocation *Allocation `json:"allocation,omitempty"`
}
