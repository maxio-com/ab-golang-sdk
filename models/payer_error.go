// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// PayerError represents a PayerError struct.
type PayerError struct {
    LastName             []string               `json:"last_name,omitempty"`
    FirstName            []string               `json:"first_name,omitempty"`
    Email                []string               `json:"email,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PayerError,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PayerError) String() string {
    return fmt.Sprintf(
    	"PayerError[LastName=%v, FirstName=%v, Email=%v, AdditionalProperties=%v]",
    	p.LastName, p.FirstName, p.Email, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PayerError.
// It customizes the JSON marshaling process for PayerError objects.
func (p PayerError) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "last_name", "first_name", "email"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PayerError object to a map representation for JSON marshaling.
func (p PayerError) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
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
    var temp tempPayerError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "last_name", "first_name", "email")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.LastName = temp.LastName
    p.FirstName = temp.FirstName
    p.Email = temp.Email
    return nil
}

// tempPayerError is a temporary struct used for validating the fields of PayerError.
type tempPayerError  struct {
    LastName  []string `json:"last_name,omitempty"`
    FirstName []string `json:"first_name,omitempty"`
    Email     []string `json:"email,omitempty"`
}
