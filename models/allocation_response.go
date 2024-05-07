package models

import (
    "encoding/json"
)

// AllocationResponse represents a AllocationResponse struct.
type AllocationResponse struct {
    Allocation           *Allocation    `json:"allocation,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AllocationResponse.
// It customizes the JSON marshaling process for AllocationResponse objects.
func (a AllocationResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AllocationResponse object to a map representation for JSON marshaling.
func (a AllocationResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Allocation != nil {
        structMap["allocation"] = a.Allocation.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationResponse.
// It customizes the JSON unmarshaling process for AllocationResponse objects.
func (a *AllocationResponse) UnmarshalJSON(input []byte) error {
    var temp allocationResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "allocation")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Allocation = temp.Allocation
    return nil
}

// allocationResponse is a temporary struct used for validating the fields of AllocationResponse.
type allocationResponse  struct {
    Allocation *Allocation `json:"allocation,omitempty"`
}
