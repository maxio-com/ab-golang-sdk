package models

import (
	"encoding/json"
)

// UpdateAllocationExpirationDate represents a UpdateAllocationExpirationDate struct.
type UpdateAllocationExpirationDate struct {
	Allocation *AllocationExpirationDate `json:"allocation,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateAllocationExpirationDate.
// It customizes the JSON marshaling process for UpdateAllocationExpirationDate objects.
func (u *UpdateAllocationExpirationDate) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateAllocationExpirationDate object to a map representation for JSON marshaling.
func (u *UpdateAllocationExpirationDate) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.Allocation != nil {
		structMap["allocation"] = u.Allocation.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateAllocationExpirationDate.
// It customizes the JSON unmarshaling process for UpdateAllocationExpirationDate objects.
func (u *UpdateAllocationExpirationDate) UnmarshalJSON(input []byte) error {
	var temp updateAllocationExpirationDate
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	u.Allocation = temp.Allocation
	return nil
}

// TODO
type updateAllocationExpirationDate struct {
	Allocation *AllocationExpirationDate `json:"allocation,omitempty"`
}
