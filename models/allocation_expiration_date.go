package models

import (
	"encoding/json"
)

// AllocationExpirationDate represents a AllocationExpirationDate struct.
type AllocationExpirationDate struct {
	ExpiresAt *string `json:"expires_at,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AllocationExpirationDate.
// It customizes the JSON marshaling process for AllocationExpirationDate objects.
func (a *AllocationExpirationDate) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AllocationExpirationDate object to a map representation for JSON marshaling.
func (a *AllocationExpirationDate) toMap() map[string]any {
	structMap := make(map[string]any)
	if a.ExpiresAt != nil {
		structMap["expires_at"] = a.ExpiresAt
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationExpirationDate.
// It customizes the JSON unmarshaling process for AllocationExpirationDate objects.
func (a *AllocationExpirationDate) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ExpiresAt *string `json:"expires_at,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.ExpiresAt = temp.ExpiresAt
	return nil
}
