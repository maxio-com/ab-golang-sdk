/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// PayerAttributes represents a PayerAttributes struct.
type PayerAttributes struct {
    FirstName            *string                `json:"first_name,omitempty"`
    LastName             *string                `json:"last_name,omitempty"`
    Email                *string                `json:"email,omitempty"`
    CcEmails             *string                `json:"cc_emails,omitempty"`
    Organization         *string                `json:"organization,omitempty"`
    Reference            *string                `json:"reference,omitempty"`
    Address              *string                `json:"address,omitempty"`
    Address2             *string                `json:"address_2,omitempty"`
    City                 *string                `json:"city,omitempty"`
    State                *string                `json:"state,omitempty"`
    Zip                  *string                `json:"zip,omitempty"`
    Country              *string                `json:"country,omitempty"`
    Phone                *string                `json:"phone,omitempty"`
    Locale               *string                `json:"locale,omitempty"`
    VatNumber            *string                `json:"vat_number,omitempty"`
    TaxExempt            *bool                  `json:"tax_exempt,omitempty"`
    TaxExemptReason      *string                `json:"tax_exempt_reason,omitempty"`
    // (Optional) A set of key/value pairs representing custom fields and their values. Metafields will be created “on-the-fly” in your site for a given key, if they have not been created yet.
    Metafields           map[string]string      `json:"metafields,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PayerAttributes,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PayerAttributes) String() string {
    return fmt.Sprintf(
    	"PayerAttributes[FirstName=%v, LastName=%v, Email=%v, CcEmails=%v, Organization=%v, Reference=%v, Address=%v, Address2=%v, City=%v, State=%v, Zip=%v, Country=%v, Phone=%v, Locale=%v, VatNumber=%v, TaxExempt=%v, TaxExemptReason=%v, Metafields=%v, AdditionalProperties=%v]",
    	p.FirstName, p.LastName, p.Email, p.CcEmails, p.Organization, p.Reference, p.Address, p.Address2, p.City, p.State, p.Zip, p.Country, p.Phone, p.Locale, p.VatNumber, p.TaxExempt, p.TaxExemptReason, p.Metafields, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PayerAttributes.
// It customizes the JSON marshaling process for PayerAttributes objects.
func (p PayerAttributes) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "first_name", "last_name", "email", "cc_emails", "organization", "reference", "address", "address_2", "city", "state", "zip", "country", "phone", "locale", "vat_number", "tax_exempt", "tax_exempt_reason", "metafields"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PayerAttributes object to a map representation for JSON marshaling.
func (p PayerAttributes) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.FirstName != nil {
        structMap["first_name"] = p.FirstName
    }
    if p.LastName != nil {
        structMap["last_name"] = p.LastName
    }
    if p.Email != nil {
        structMap["email"] = p.Email
    }
    if p.CcEmails != nil {
        structMap["cc_emails"] = p.CcEmails
    }
    if p.Organization != nil {
        structMap["organization"] = p.Organization
    }
    if p.Reference != nil {
        structMap["reference"] = p.Reference
    }
    if p.Address != nil {
        structMap["address"] = p.Address
    }
    if p.Address2 != nil {
        structMap["address_2"] = p.Address2
    }
    if p.City != nil {
        structMap["city"] = p.City
    }
    if p.State != nil {
        structMap["state"] = p.State
    }
    if p.Zip != nil {
        structMap["zip"] = p.Zip
    }
    if p.Country != nil {
        structMap["country"] = p.Country
    }
    if p.Phone != nil {
        structMap["phone"] = p.Phone
    }
    if p.Locale != nil {
        structMap["locale"] = p.Locale
    }
    if p.VatNumber != nil {
        structMap["vat_number"] = p.VatNumber
    }
    if p.TaxExempt != nil {
        structMap["tax_exempt"] = p.TaxExempt
    }
    if p.TaxExemptReason != nil {
        structMap["tax_exempt_reason"] = p.TaxExemptReason
    }
    if p.Metafields != nil {
        structMap["metafields"] = p.Metafields
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PayerAttributes.
// It customizes the JSON unmarshaling process for PayerAttributes objects.
func (p *PayerAttributes) UnmarshalJSON(input []byte) error {
    var temp tempPayerAttributes
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "first_name", "last_name", "email", "cc_emails", "organization", "reference", "address", "address_2", "city", "state", "zip", "country", "phone", "locale", "vat_number", "tax_exempt", "tax_exempt_reason", "metafields")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.FirstName = temp.FirstName
    p.LastName = temp.LastName
    p.Email = temp.Email
    p.CcEmails = temp.CcEmails
    p.Organization = temp.Organization
    p.Reference = temp.Reference
    p.Address = temp.Address
    p.Address2 = temp.Address2
    p.City = temp.City
    p.State = temp.State
    p.Zip = temp.Zip
    p.Country = temp.Country
    p.Phone = temp.Phone
    p.Locale = temp.Locale
    p.VatNumber = temp.VatNumber
    p.TaxExempt = temp.TaxExempt
    p.TaxExemptReason = temp.TaxExemptReason
    p.Metafields = temp.Metafields
    return nil
}

// tempPayerAttributes is a temporary struct used for validating the fields of PayerAttributes.
type tempPayerAttributes  struct {
    FirstName       *string           `json:"first_name,omitempty"`
    LastName        *string           `json:"last_name,omitempty"`
    Email           *string           `json:"email,omitempty"`
    CcEmails        *string           `json:"cc_emails,omitempty"`
    Organization    *string           `json:"organization,omitempty"`
    Reference       *string           `json:"reference,omitempty"`
    Address         *string           `json:"address,omitempty"`
    Address2        *string           `json:"address_2,omitempty"`
    City            *string           `json:"city,omitempty"`
    State           *string           `json:"state,omitempty"`
    Zip             *string           `json:"zip,omitempty"`
    Country         *string           `json:"country,omitempty"`
    Phone           *string           `json:"phone,omitempty"`
    Locale          *string           `json:"locale,omitempty"`
    VatNumber       *string           `json:"vat_number,omitempty"`
    TaxExempt       *bool             `json:"tax_exempt,omitempty"`
    TaxExemptReason *string           `json:"tax_exempt_reason,omitempty"`
    Metafields      map[string]string `json:"metafields,omitempty"`
}
