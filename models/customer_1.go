package models

import (
    "encoding/json"
)

// Customer1 represents a Customer1 struct.
type Customer1 struct {
    ChargifyId   Optional[int]    `json:"chargify_id"`
    FirstName    *string          `json:"first_name,omitempty"`
    LastName     *string          `json:"last_name,omitempty"`
    Organization Optional[string] `json:"organization"`
    Email        *string          `json:"email,omitempty"`
    VatNumber    Optional[string] `json:"vat_number"`
    Reference    Optional[string] `json:"reference"`
}

// MarshalJSON implements the json.Marshaler interface for Customer1.
// It customizes the JSON marshaling process for Customer1 objects.
func (c *Customer1) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the Customer1 object to a map representation for JSON marshaling.
func (c *Customer1) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.ChargifyId.IsValueSet() {
        structMap["chargify_id"] = c.ChargifyId.Value()
    }
    if c.FirstName != nil {
        structMap["first_name"] = c.FirstName
    }
    if c.LastName != nil {
        structMap["last_name"] = c.LastName
    }
    if c.Organization.IsValueSet() {
        structMap["organization"] = c.Organization.Value()
    }
    if c.Email != nil {
        structMap["email"] = c.Email
    }
    if c.VatNumber.IsValueSet() {
        structMap["vat_number"] = c.VatNumber.Value()
    }
    if c.Reference.IsValueSet() {
        structMap["reference"] = c.Reference.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Customer1.
// It customizes the JSON unmarshaling process for Customer1 objects.
func (c *Customer1) UnmarshalJSON(input []byte) error {
    temp := &struct {
        ChargifyId   Optional[int]    `json:"chargify_id"`
        FirstName    *string          `json:"first_name,omitempty"`
        LastName     *string          `json:"last_name,omitempty"`
        Organization Optional[string] `json:"organization"`
        Email        *string          `json:"email,omitempty"`
        VatNumber    Optional[string] `json:"vat_number"`
        Reference    Optional[string] `json:"reference"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.ChargifyId = temp.ChargifyId
    c.FirstName = temp.FirstName
    c.LastName = temp.LastName
    c.Organization = temp.Organization
    c.Email = temp.Email
    c.VatNumber = temp.VatNumber
    c.Reference = temp.Reference
    return nil
}
