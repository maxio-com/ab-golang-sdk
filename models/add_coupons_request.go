package models

import (
    "encoding/json"
)

// AddCouponsRequest represents a AddCouponsRequest struct.
type AddCouponsRequest struct {
    Codes []string `json:"codes,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AddCouponsRequest.
// It customizes the JSON marshaling process for AddCouponsRequest objects.
func (a *AddCouponsRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AddCouponsRequest object to a map representation for JSON marshaling.
func (a *AddCouponsRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    if a.Codes != nil {
        structMap["codes"] = a.Codes
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddCouponsRequest.
// It customizes the JSON unmarshaling process for AddCouponsRequest objects.
func (a *AddCouponsRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Codes []string `json:"codes,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    a.Codes = temp.Codes
    return nil
}
