package models

import (
    "encoding/json"
)

// InvoicePayer represents a InvoicePayer struct.
type InvoicePayer struct {
    ChargifyId           *int             `json:"chargify_id,omitempty"`
    FirstName            *string          `json:"first_name,omitempty"`
    LastName             *string          `json:"last_name,omitempty"`
    Organization         Optional[string] `json:"organization"`
    Email                *string          `json:"email,omitempty"`
    VatNumber            Optional[string] `json:"vat_number"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoicePayer.
// It customizes the JSON marshaling process for InvoicePayer objects.
func (i InvoicePayer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoicePayer object to a map representation for JSON marshaling.
func (i InvoicePayer) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.ChargifyId != nil {
        structMap["chargify_id"] = i.ChargifyId
    }
    if i.FirstName != nil {
        structMap["first_name"] = i.FirstName
    }
    if i.LastName != nil {
        structMap["last_name"] = i.LastName
    }
    if i.Organization.IsValueSet() {
        if i.Organization.Value() != nil {
            structMap["organization"] = i.Organization.Value()
        } else {
            structMap["organization"] = nil
        }
    }
    if i.Email != nil {
        structMap["email"] = i.Email
    }
    if i.VatNumber.IsValueSet() {
        if i.VatNumber.Value() != nil {
            structMap["vat_number"] = i.VatNumber.Value()
        } else {
            structMap["vat_number"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoicePayer.
// It customizes the JSON unmarshaling process for InvoicePayer objects.
func (i *InvoicePayer) UnmarshalJSON(input []byte) error {
    var temp invoicePayer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "chargify_id", "first_name", "last_name", "organization", "email", "vat_number")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.ChargifyId = temp.ChargifyId
    i.FirstName = temp.FirstName
    i.LastName = temp.LastName
    i.Organization = temp.Organization
    i.Email = temp.Email
    i.VatNumber = temp.VatNumber
    return nil
}

// TODO
type invoicePayer  struct {
    ChargifyId   *int             `json:"chargify_id,omitempty"`
    FirstName    *string          `json:"first_name,omitempty"`
    LastName     *string          `json:"last_name,omitempty"`
    Organization Optional[string] `json:"organization"`
    Email        *string          `json:"email,omitempty"`
    VatNumber    Optional[string] `json:"vat_number"`
}
