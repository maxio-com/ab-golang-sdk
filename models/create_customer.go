// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// CreateCustomer represents a CreateCustomer struct.
type CreateCustomer struct {
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	CcEmails     *string `json:"cc_emails,omitempty"`
	Organization *string `json:"organization,omitempty"`
	Reference    *string `json:"reference,omitempty"`
	Address      *string `json:"address,omitempty"`
	Address2     *string `json:"address_2,omitempty"`
	City         *string `json:"city,omitempty"`
	State        *string `json:"state,omitempty"`
	Zip          *string `json:"zip,omitempty"`
	Country      *string `json:"country,omitempty"`
	Phone        *string `json:"phone,omitempty"`
	// Set a specific language on a customer record.
	Locale          *string `json:"locale,omitempty"`
	VatNumber       *string `json:"vat_number,omitempty"`
	TaxExempt       *bool   `json:"tax_exempt,omitempty"`
	TaxExemptReason *string `json:"tax_exempt_reason,omitempty"`
	// The parent ID in Chargify if applicable. Parent is another Customer object.
	ParentId Optional[int] `json:"parent_id"`
	// The Salesforce ID of the customer
	SalesforceId         Optional[string]       `json:"salesforce_id"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateCustomer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateCustomer) String() string {
	return fmt.Sprintf(
		"CreateCustomer[FirstName=%v, LastName=%v, Email=%v, CcEmails=%v, Organization=%v, Reference=%v, Address=%v, Address2=%v, City=%v, State=%v, Zip=%v, Country=%v, Phone=%v, Locale=%v, VatNumber=%v, TaxExempt=%v, TaxExemptReason=%v, ParentId=%v, SalesforceId=%v, AdditionalProperties=%v]",
		c.FirstName, c.LastName, c.Email, c.CcEmails, c.Organization, c.Reference, c.Address, c.Address2, c.City, c.State, c.Zip, c.Country, c.Phone, c.Locale, c.VatNumber, c.TaxExempt, c.TaxExemptReason, c.ParentId, c.SalesforceId, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateCustomer.
// It customizes the JSON marshaling process for CreateCustomer objects.
func (c CreateCustomer) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"first_name", "last_name", "email", "cc_emails", "organization", "reference", "address", "address_2", "city", "state", "zip", "country", "phone", "locale", "vat_number", "tax_exempt", "tax_exempt_reason", "parent_id", "salesforce_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCustomer object to a map representation for JSON marshaling.
func (c CreateCustomer) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	structMap["first_name"] = c.FirstName
	structMap["last_name"] = c.LastName
	structMap["email"] = c.Email
	if c.CcEmails != nil {
		structMap["cc_emails"] = c.CcEmails
	}
	if c.Organization != nil {
		structMap["organization"] = c.Organization
	}
	if c.Reference != nil {
		structMap["reference"] = c.Reference
	}
	if c.Address != nil {
		structMap["address"] = c.Address
	}
	if c.Address2 != nil {
		structMap["address_2"] = c.Address2
	}
	if c.City != nil {
		structMap["city"] = c.City
	}
	if c.State != nil {
		structMap["state"] = c.State
	}
	if c.Zip != nil {
		structMap["zip"] = c.Zip
	}
	if c.Country != nil {
		structMap["country"] = c.Country
	}
	if c.Phone != nil {
		structMap["phone"] = c.Phone
	}
	if c.Locale != nil {
		structMap["locale"] = c.Locale
	}
	if c.VatNumber != nil {
		structMap["vat_number"] = c.VatNumber
	}
	if c.TaxExempt != nil {
		structMap["tax_exempt"] = c.TaxExempt
	}
	if c.TaxExemptReason != nil {
		structMap["tax_exempt_reason"] = c.TaxExemptReason
	}
	if c.ParentId.IsValueSet() {
		if c.ParentId.Value() != nil {
			structMap["parent_id"] = c.ParentId.Value()
		} else {
			structMap["parent_id"] = nil
		}
	}
	if c.SalesforceId.IsValueSet() {
		if c.SalesforceId.Value() != nil {
			structMap["salesforce_id"] = c.SalesforceId.Value()
		} else {
			structMap["salesforce_id"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCustomer.
// It customizes the JSON unmarshaling process for CreateCustomer objects.
func (c *CreateCustomer) UnmarshalJSON(input []byte) error {
	var temp tempCreateCustomer
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "first_name", "last_name", "email", "cc_emails", "organization", "reference", "address", "address_2", "city", "state", "zip", "country", "phone", "locale", "vat_number", "tax_exempt", "tax_exempt_reason", "parent_id", "salesforce_id")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.FirstName = *temp.FirstName
	c.LastName = *temp.LastName
	c.Email = *temp.Email
	c.CcEmails = temp.CcEmails
	c.Organization = temp.Organization
	c.Reference = temp.Reference
	c.Address = temp.Address
	c.Address2 = temp.Address2
	c.City = temp.City
	c.State = temp.State
	c.Zip = temp.Zip
	c.Country = temp.Country
	c.Phone = temp.Phone
	c.Locale = temp.Locale
	c.VatNumber = temp.VatNumber
	c.TaxExempt = temp.TaxExempt
	c.TaxExemptReason = temp.TaxExemptReason
	c.ParentId = temp.ParentId
	c.SalesforceId = temp.SalesforceId
	return nil
}

// tempCreateCustomer is a temporary struct used for validating the fields of CreateCustomer.
type tempCreateCustomer struct {
	FirstName       *string          `json:"first_name"`
	LastName        *string          `json:"last_name"`
	Email           *string          `json:"email"`
	CcEmails        *string          `json:"cc_emails,omitempty"`
	Organization    *string          `json:"organization,omitempty"`
	Reference       *string          `json:"reference,omitempty"`
	Address         *string          `json:"address,omitempty"`
	Address2        *string          `json:"address_2,omitempty"`
	City            *string          `json:"city,omitempty"`
	State           *string          `json:"state,omitempty"`
	Zip             *string          `json:"zip,omitempty"`
	Country         *string          `json:"country,omitempty"`
	Phone           *string          `json:"phone,omitempty"`
	Locale          *string          `json:"locale,omitempty"`
	VatNumber       *string          `json:"vat_number,omitempty"`
	TaxExempt       *bool            `json:"tax_exempt,omitempty"`
	TaxExemptReason *string          `json:"tax_exempt_reason,omitempty"`
	ParentId        Optional[int]    `json:"parent_id"`
	SalesforceId    Optional[string] `json:"salesforce_id"`
}

func (c *tempCreateCustomer) validate() error {
	var errs []string
	if c.FirstName == nil {
		errs = append(errs, "required field `first_name` is missing for type `Create Customer`")
	}
	if c.LastName == nil {
		errs = append(errs, "required field `last_name` is missing for type `Create Customer`")
	}
	if c.Email == nil {
		errs = append(errs, "required field `email` is missing for type `Create Customer`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
