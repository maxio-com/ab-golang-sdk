/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// UpdateCustomer represents a UpdateCustomer struct.
type UpdateCustomer struct {
    FirstName            *string        `json:"first_name,omitempty"`
    LastName             *string        `json:"last_name,omitempty"`
    Email                *string        `json:"email,omitempty"`
    CcEmails             *string        `json:"cc_emails,omitempty"`
    Organization         *string        `json:"organization,omitempty"`
    Reference            *string        `json:"reference,omitempty"`
    Address              *string        `json:"address,omitempty"`
    Address2             *string        `json:"address_2,omitempty"`
    City                 *string        `json:"city,omitempty"`
    State                *string        `json:"state,omitempty"`
    Zip                  *string        `json:"zip,omitempty"`
    Country              *string        `json:"country,omitempty"`
    Phone                *string        `json:"phone,omitempty"`
    // Set a specific language on a customer record.
    Locale               *string        `json:"locale,omitempty"`
    VatNumber            *string        `json:"vat_number,omitempty"`
    TaxExempt            *bool          `json:"tax_exempt,omitempty"`
    TaxExemptReason      *string        `json:"tax_exempt_reason,omitempty"`
    ParentId             Optional[int]  `json:"parent_id"`
    // Is the customer verified to use ACH as a payment method. Available only on Authorize.Net gateway
    Verified             Optional[bool] `json:"verified"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateCustomer.
// It customizes the JSON marshaling process for UpdateCustomer objects.
func (u UpdateCustomer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateCustomer object to a map representation for JSON marshaling.
func (u UpdateCustomer) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.FirstName != nil {
        structMap["first_name"] = u.FirstName
    }
    if u.LastName != nil {
        structMap["last_name"] = u.LastName
    }
    if u.Email != nil {
        structMap["email"] = u.Email
    }
    if u.CcEmails != nil {
        structMap["cc_emails"] = u.CcEmails
    }
    if u.Organization != nil {
        structMap["organization"] = u.Organization
    }
    if u.Reference != nil {
        structMap["reference"] = u.Reference
    }
    if u.Address != nil {
        structMap["address"] = u.Address
    }
    if u.Address2 != nil {
        structMap["address_2"] = u.Address2
    }
    if u.City != nil {
        structMap["city"] = u.City
    }
    if u.State != nil {
        structMap["state"] = u.State
    }
    if u.Zip != nil {
        structMap["zip"] = u.Zip
    }
    if u.Country != nil {
        structMap["country"] = u.Country
    }
    if u.Phone != nil {
        structMap["phone"] = u.Phone
    }
    if u.Locale != nil {
        structMap["locale"] = u.Locale
    }
    if u.VatNumber != nil {
        structMap["vat_number"] = u.VatNumber
    }
    if u.TaxExempt != nil {
        structMap["tax_exempt"] = u.TaxExempt
    }
    if u.TaxExemptReason != nil {
        structMap["tax_exempt_reason"] = u.TaxExemptReason
    }
    if u.ParentId.IsValueSet() {
        if u.ParentId.Value() != nil {
            structMap["parent_id"] = u.ParentId.Value()
        } else {
            structMap["parent_id"] = nil
        }
    }
    if u.Verified.IsValueSet() {
        if u.Verified.Value() != nil {
            structMap["verified"] = u.Verified.Value()
        } else {
            structMap["verified"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateCustomer.
// It customizes the JSON unmarshaling process for UpdateCustomer objects.
func (u *UpdateCustomer) UnmarshalJSON(input []byte) error {
    var temp tempUpdateCustomer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "first_name", "last_name", "email", "cc_emails", "organization", "reference", "address", "address_2", "city", "state", "zip", "country", "phone", "locale", "vat_number", "tax_exempt", "tax_exempt_reason", "parent_id", "verified")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.FirstName = temp.FirstName
    u.LastName = temp.LastName
    u.Email = temp.Email
    u.CcEmails = temp.CcEmails
    u.Organization = temp.Organization
    u.Reference = temp.Reference
    u.Address = temp.Address
    u.Address2 = temp.Address2
    u.City = temp.City
    u.State = temp.State
    u.Zip = temp.Zip
    u.Country = temp.Country
    u.Phone = temp.Phone
    u.Locale = temp.Locale
    u.VatNumber = temp.VatNumber
    u.TaxExempt = temp.TaxExempt
    u.TaxExemptReason = temp.TaxExemptReason
    u.ParentId = temp.ParentId
    u.Verified = temp.Verified
    return nil
}

// tempUpdateCustomer is a temporary struct used for validating the fields of UpdateCustomer.
type tempUpdateCustomer  struct {
    FirstName       *string        `json:"first_name,omitempty"`
    LastName        *string        `json:"last_name,omitempty"`
    Email           *string        `json:"email,omitempty"`
    CcEmails        *string        `json:"cc_emails,omitempty"`
    Organization    *string        `json:"organization,omitempty"`
    Reference       *string        `json:"reference,omitempty"`
    Address         *string        `json:"address,omitempty"`
    Address2        *string        `json:"address_2,omitempty"`
    City            *string        `json:"city,omitempty"`
    State           *string        `json:"state,omitempty"`
    Zip             *string        `json:"zip,omitempty"`
    Country         *string        `json:"country,omitempty"`
    Phone           *string        `json:"phone,omitempty"`
    Locale          *string        `json:"locale,omitempty"`
    VatNumber       *string        `json:"vat_number,omitempty"`
    TaxExempt       *bool          `json:"tax_exempt,omitempty"`
    TaxExemptReason *string        `json:"tax_exempt_reason,omitempty"`
    ParentId        Optional[int]  `json:"parent_id"`
    Verified        Optional[bool] `json:"verified"`
}
