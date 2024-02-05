package models

import (
    "encoding/json"
)

// AddressChange represents a AddressChange struct.
type AddressChange struct {
    Before InvoiceAddress `json:"before"`
    After  InvoiceAddress `json:"after"`
}

// MarshalJSON implements the json.Marshaler interface for AddressChange.
// It customizes the JSON marshaling process for AddressChange objects.
func (a *AddressChange) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AddressChange object to a map representation for JSON marshaling.
func (a *AddressChange) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["before"] = a.Before.toMap()
    structMap["after"] = a.After.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddressChange.
// It customizes the JSON unmarshaling process for AddressChange objects.
func (a *AddressChange) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Before InvoiceAddress `json:"before"`
        After  InvoiceAddress `json:"after"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    a.Before = temp.Before
    a.After = temp.After
    return nil
}
