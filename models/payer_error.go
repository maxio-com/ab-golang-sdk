package models

import (
    "encoding/json"
)

// PayerError represents a PayerError struct.
type PayerError struct {
    LastName             []string       `json:"last_name,omitempty"`
    FirstName            []string       `json:"first_name,omitempty"`
    Email                []string       `json:"email,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PayerError.
// It customizes the JSON marshaling process for PayerError objects.
func (p PayerError) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PayerError object to a map representation for JSON marshaling.
func (p PayerError) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.LastName != nil {
        structMap["last_name"] = p.LastName
    }
    if p.FirstName != nil {
        structMap["first_name"] = p.FirstName
    }
    if p.Email != nil {
        structMap["email"] = p.Email
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PayerError.
// It customizes the JSON unmarshaling process for PayerError objects.
func (p *PayerError) UnmarshalJSON(input []byte) error {
    var temp payerError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "last_name", "first_name", "email")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.LastName = temp.LastName
    p.FirstName = temp.FirstName
    p.Email = temp.Email
    return nil
}

// TODO
type payerError  struct {
    LastName  []string `json:"last_name,omitempty"`
    FirstName []string `json:"first_name,omitempty"`
    Email     []string `json:"email,omitempty"`
}
