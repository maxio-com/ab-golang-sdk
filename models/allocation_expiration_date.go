package models

import (
    "encoding/json"
    "log"
    "time"
)

// AllocationExpirationDate represents a AllocationExpirationDate struct.
type AllocationExpirationDate struct {
    ExpiresAt *time.Time `json:"expires_at,omitempty"`
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
        structMap["expires_at"] = a.ExpiresAt.Format(time.RFC3339)
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
    
    if temp.ExpiresAt != nil {
        ExpiresAtVal, err := time.Parse(time.RFC3339, *temp.ExpiresAt)
        if err != nil {
            log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
        }
        a.ExpiresAt = &ExpiresAtVal
    }
    return nil
}
