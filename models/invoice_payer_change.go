package models

import (
    "encoding/json"
)

// InvoicePayerChange represents a InvoicePayerChange struct.
type InvoicePayerChange struct {
    FirstName            *string        `json:"first_name,omitempty"`
    LastName             *string        `json:"last_name,omitempty"`
    Organization         *string        `json:"organization,omitempty"`
    Email                *string        `json:"email,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoicePayerChange.
// It customizes the JSON marshaling process for InvoicePayerChange objects.
func (i InvoicePayerChange) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoicePayerChange object to a map representation for JSON marshaling.
func (i InvoicePayerChange) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.FirstName != nil {
        structMap["first_name"] = i.FirstName
    }
    if i.LastName != nil {
        structMap["last_name"] = i.LastName
    }
    if i.Organization != nil {
        structMap["organization"] = i.Organization
    }
    if i.Email != nil {
        structMap["email"] = i.Email
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoicePayerChange.
// It customizes the JSON unmarshaling process for InvoicePayerChange objects.
func (i *InvoicePayerChange) UnmarshalJSON(input []byte) error {
    var temp invoicePayerChange
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "first_name", "last_name", "organization", "email")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.FirstName = temp.FirstName
    i.LastName = temp.LastName
    i.Organization = temp.Organization
    i.Email = temp.Email
    return nil
}

// invoicePayerChange is a temporary struct used for validating the fields of InvoicePayerChange.
type invoicePayerChange  struct {
    FirstName    *string `json:"first_name,omitempty"`
    LastName     *string `json:"last_name,omitempty"`
    Organization *string `json:"organization,omitempty"`
    Email        *string `json:"email,omitempty"`
}
