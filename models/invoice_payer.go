/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// InvoicePayer represents a InvoicePayer struct.
type InvoicePayer struct {
    ChargifyId           *int                   `json:"chargify_id,omitempty"`
    FirstName            *string                `json:"first_name,omitempty"`
    LastName             *string                `json:"last_name,omitempty"`
    Organization         Optional[string]       `json:"organization"`
    Email                *string                `json:"email,omitempty"`
    VatNumber            Optional[string]       `json:"vat_number"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InvoicePayer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoicePayer) String() string {
    return fmt.Sprintf(
    	"InvoicePayer[ChargifyId=%v, FirstName=%v, LastName=%v, Organization=%v, Email=%v, VatNumber=%v, AdditionalProperties=%v]",
    	i.ChargifyId, i.FirstName, i.LastName, i.Organization, i.Email, i.VatNumber, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoicePayer.
// It customizes the JSON marshaling process for InvoicePayer objects.
func (i InvoicePayer) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "chargify_id", "first_name", "last_name", "organization", "email", "vat_number"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoicePayer object to a map representation for JSON marshaling.
func (i InvoicePayer) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
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
    var temp tempInvoicePayer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "chargify_id", "first_name", "last_name", "organization", "email", "vat_number")
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

// tempInvoicePayer is a temporary struct used for validating the fields of InvoicePayer.
type tempInvoicePayer  struct {
    ChargifyId   *int             `json:"chargify_id,omitempty"`
    FirstName    *string          `json:"first_name,omitempty"`
    LastName     *string          `json:"last_name,omitempty"`
    Organization Optional[string] `json:"organization"`
    Email        *string          `json:"email,omitempty"`
    VatNumber    Optional[string] `json:"vat_number"`
}
