package models

import (
	"encoding/json"
)

// AllocationResponse represents a AllocationResponse struct.
type AllocationResponse struct {
	Allocation *Allocation `json:"allocation,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AllocationResponse.
// It customizes the JSON marshaling process for AllocationResponse objects.
func (a *AllocationResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AllocationResponse object to a map representation for JSON marshaling.
func (a *AllocationResponse) toMap() map[string]any {
	structMap := make(map[string]any)
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
	a.Allocation = temp.Allocation
	return nil
}

// TODO
type allocationResponse struct {
	Allocation *Allocation `json:"allocation,omitempty"`
}
