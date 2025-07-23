// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Customer represents a Customer struct.
type Customer struct {
	// The first name of the customer
	FirstName *string `json:"first_name,omitempty"`
	// The last name of the customer
	LastName *string `json:"last_name,omitempty"`
	// The email address of the customer
	Email *string `json:"email,omitempty"`
	// A comma-separated list of emails that should be cc’d on all customer communications (i.e. “joe@example.com, sue@example.com”)
	CcEmails Optional[string] `json:"cc_emails"`
	// The organization of the customer. If no value, `null` or empty string is provided, `organization` will be populated with the customer's first and last name, separated with a space.
	Organization Optional[string] `json:"organization"`
	// The unique identifier used within your own application for this customer
	Reference Optional[string] `json:"reference"`
	// The customer ID in Chargify
	Id *int `json:"id,omitempty"`
	// The timestamp in which the customer object was created in Chargify
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The timestamp in which the customer object was last edited
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The customer’s shipping street address (i.e. “123 Main St.”)
	Address Optional[string] `json:"address"`
	// Second line of the customer’s shipping address i.e. “Apt. 100”
	Address2 Optional[string] `json:"address_2"`
	// The customer’s shipping address city (i.e. “Boston”)
	City Optional[string] `json:"city"`
	// The customer’s shipping address state (i.e. “MA”)
	State Optional[string] `json:"state"`
	// The customer's full name of state
	StateName Optional[string] `json:"state_name"`
	// The customer’s shipping address zip code (i.e. “12345”)
	Zip Optional[string] `json:"zip"`
	// The customer shipping address country
	Country Optional[string] `json:"country"`
	// The customer's full name of country
	CountryName Optional[string] `json:"country_name"`
	// The phone number of the customer
	Phone Optional[string] `json:"phone"`
	// Is the customer verified to use ACH as a payment method.
	Verified Optional[bool] `json:"verified"`
	// The timestamp of when the Billing Portal entry was created at for the customer
	PortalCustomerCreatedAt Optional[time.Time] `json:"portal_customer_created_at"`
	// The timestamp of when the Billing Portal invite was last sent at
	PortalInviteLastSentAt Optional[time.Time] `json:"portal_invite_last_sent_at"`
	// The timestamp of when the Billing Portal invite was last accepted
	PortalInviteLastAcceptedAt Optional[time.Time] `json:"portal_invite_last_accepted_at"`
	// The tax exempt status for the customer. Acceptable values are true or 1 for true and false or 0 for false.
	TaxExempt *bool `json:"tax_exempt,omitempty"`
	// The VAT business identification number for the customer. This number is used to determine VAT tax opt out rules. It is not validated when added or updated on a customer record. Instead, it is validated via VIES before calculating taxes. Only valid business identification numbers will allow for VAT opt out.
	VatNumber Optional[string] `json:"vat_number"`
	// The parent ID in Chargify if applicable. Parent is another Customer object.
	ParentId Optional[int] `json:"parent_id"`
	// The locale for the customer to identify language-region
	Locale                      Optional[string] `json:"locale"`
	DefaultSubscriptionGroupUid Optional[string] `json:"default_subscription_group_uid"`
	// The Salesforce ID for the customer
	SalesforceId Optional[string] `json:"salesforce_id"`
	// The Tax Exemption Reason Code for the customer
	TaxExemptReason Optional[string] `json:"tax_exempt_reason"`
	// The default auto-renewal profile ID for the customer
	DefaultAutoRenewalProfileId Optional[int]          `json:"default_auto_renewal_profile_id"`
	AdditionalProperties        map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Customer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c Customer) String() string {
	return fmt.Sprintf(
		"Customer[FirstName=%v, LastName=%v, Email=%v, CcEmails=%v, Organization=%v, Reference=%v, Id=%v, CreatedAt=%v, UpdatedAt=%v, Address=%v, Address2=%v, City=%v, State=%v, StateName=%v, Zip=%v, Country=%v, CountryName=%v, Phone=%v, Verified=%v, PortalCustomerCreatedAt=%v, PortalInviteLastSentAt=%v, PortalInviteLastAcceptedAt=%v, TaxExempt=%v, VatNumber=%v, ParentId=%v, Locale=%v, DefaultSubscriptionGroupUid=%v, SalesforceId=%v, TaxExemptReason=%v, DefaultAutoRenewalProfileId=%v, AdditionalProperties=%v]",
		c.FirstName, c.LastName, c.Email, c.CcEmails, c.Organization, c.Reference, c.Id, c.CreatedAt, c.UpdatedAt, c.Address, c.Address2, c.City, c.State, c.StateName, c.Zip, c.Country, c.CountryName, c.Phone, c.Verified, c.PortalCustomerCreatedAt, c.PortalInviteLastSentAt, c.PortalInviteLastAcceptedAt, c.TaxExempt, c.VatNumber, c.ParentId, c.Locale, c.DefaultSubscriptionGroupUid, c.SalesforceId, c.TaxExemptReason, c.DefaultAutoRenewalProfileId, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Customer.
// It customizes the JSON marshaling process for Customer objects.
func (c Customer) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(c.AdditionalProperties,
		"first_name", "last_name", "email", "cc_emails", "organization", "reference", "id", "created_at", "updated_at", "address", "address_2", "city", "state", "state_name", "zip", "country", "country_name", "phone", "verified", "portal_customer_created_at", "portal_invite_last_sent_at", "portal_invite_last_accepted_at", "tax_exempt", "vat_number", "parent_id", "locale", "default_subscription_group_uid", "salesforce_id", "tax_exempt_reason", "default_auto_renewal_profile_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c.toMap())
}

// toMap converts the Customer object to a map representation for JSON marshaling.
func (c Customer) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, c.AdditionalProperties)
	if c.FirstName != nil {
		structMap["first_name"] = c.FirstName
	}
	if c.LastName != nil {
		structMap["last_name"] = c.LastName
	}
	if c.Email != nil {
		structMap["email"] = c.Email
	}
	if c.CcEmails.IsValueSet() {
		if c.CcEmails.Value() != nil {
			structMap["cc_emails"] = c.CcEmails.Value()
		} else {
			structMap["cc_emails"] = nil
		}
	}
	if c.Organization.IsValueSet() {
		if c.Organization.Value() != nil {
			structMap["organization"] = c.Organization.Value()
		} else {
			structMap["organization"] = nil
		}
	}
	if c.Reference.IsValueSet() {
		if c.Reference.Value() != nil {
			structMap["reference"] = c.Reference.Value()
		} else {
			structMap["reference"] = nil
		}
	}
	if c.Id != nil {
		structMap["id"] = c.Id
	}
	if c.CreatedAt != nil {
		structMap["created_at"] = c.CreatedAt.Format(time.RFC3339)
	}
	if c.UpdatedAt != nil {
		structMap["updated_at"] = c.UpdatedAt.Format(time.RFC3339)
	}
	if c.Address.IsValueSet() {
		if c.Address.Value() != nil {
			structMap["address"] = c.Address.Value()
		} else {
			structMap["address"] = nil
		}
	}
	if c.Address2.IsValueSet() {
		if c.Address2.Value() != nil {
			structMap["address_2"] = c.Address2.Value()
		} else {
			structMap["address_2"] = nil
		}
	}
	if c.City.IsValueSet() {
		if c.City.Value() != nil {
			structMap["city"] = c.City.Value()
		} else {
			structMap["city"] = nil
		}
	}
	if c.State.IsValueSet() {
		if c.State.Value() != nil {
			structMap["state"] = c.State.Value()
		} else {
			structMap["state"] = nil
		}
	}
	if c.StateName.IsValueSet() {
		if c.StateName.Value() != nil {
			structMap["state_name"] = c.StateName.Value()
		} else {
			structMap["state_name"] = nil
		}
	}
	if c.Zip.IsValueSet() {
		if c.Zip.Value() != nil {
			structMap["zip"] = c.Zip.Value()
		} else {
			structMap["zip"] = nil
		}
	}
	if c.Country.IsValueSet() {
		if c.Country.Value() != nil {
			structMap["country"] = c.Country.Value()
		} else {
			structMap["country"] = nil
		}
	}
	if c.CountryName.IsValueSet() {
		if c.CountryName.Value() != nil {
			structMap["country_name"] = c.CountryName.Value()
		} else {
			structMap["country_name"] = nil
		}
	}
	if c.Phone.IsValueSet() {
		if c.Phone.Value() != nil {
			structMap["phone"] = c.Phone.Value()
		} else {
			structMap["phone"] = nil
		}
	}
	if c.Verified.IsValueSet() {
		if c.Verified.Value() != nil {
			structMap["verified"] = c.Verified.Value()
		} else {
			structMap["verified"] = nil
		}
	}
	if c.PortalCustomerCreatedAt.IsValueSet() {
		var PortalCustomerCreatedAtVal *string = nil
		if c.PortalCustomerCreatedAt.Value() != nil {
			val := c.PortalCustomerCreatedAt.Value().Format(time.RFC3339)
			PortalCustomerCreatedAtVal = &val
		}
		if c.PortalCustomerCreatedAt.Value() != nil {
			structMap["portal_customer_created_at"] = PortalCustomerCreatedAtVal
		} else {
			structMap["portal_customer_created_at"] = nil
		}
	}
	if c.PortalInviteLastSentAt.IsValueSet() {
		var PortalInviteLastSentAtVal *string = nil
		if c.PortalInviteLastSentAt.Value() != nil {
			val := c.PortalInviteLastSentAt.Value().Format(time.RFC3339)
			PortalInviteLastSentAtVal = &val
		}
		if c.PortalInviteLastSentAt.Value() != nil {
			structMap["portal_invite_last_sent_at"] = PortalInviteLastSentAtVal
		} else {
			structMap["portal_invite_last_sent_at"] = nil
		}
	}
	if c.PortalInviteLastAcceptedAt.IsValueSet() {
		var PortalInviteLastAcceptedAtVal *string = nil
		if c.PortalInviteLastAcceptedAt.Value() != nil {
			val := c.PortalInviteLastAcceptedAt.Value().Format(time.RFC3339)
			PortalInviteLastAcceptedAtVal = &val
		}
		if c.PortalInviteLastAcceptedAt.Value() != nil {
			structMap["portal_invite_last_accepted_at"] = PortalInviteLastAcceptedAtVal
		} else {
			structMap["portal_invite_last_accepted_at"] = nil
		}
	}
	if c.TaxExempt != nil {
		structMap["tax_exempt"] = c.TaxExempt
	}
	if c.VatNumber.IsValueSet() {
		if c.VatNumber.Value() != nil {
			structMap["vat_number"] = c.VatNumber.Value()
		} else {
			structMap["vat_number"] = nil
		}
	}
	if c.ParentId.IsValueSet() {
		if c.ParentId.Value() != nil {
			structMap["parent_id"] = c.ParentId.Value()
		} else {
			structMap["parent_id"] = nil
		}
	}
	if c.Locale.IsValueSet() {
		if c.Locale.Value() != nil {
			structMap["locale"] = c.Locale.Value()
		} else {
			structMap["locale"] = nil
		}
	}
	if c.DefaultSubscriptionGroupUid.IsValueSet() {
		if c.DefaultSubscriptionGroupUid.Value() != nil {
			structMap["default_subscription_group_uid"] = c.DefaultSubscriptionGroupUid.Value()
		} else {
			structMap["default_subscription_group_uid"] = nil
		}
	}
	if c.SalesforceId.IsValueSet() {
		if c.SalesforceId.Value() != nil {
			structMap["salesforce_id"] = c.SalesforceId.Value()
		} else {
			structMap["salesforce_id"] = nil
		}
	}
	if c.TaxExemptReason.IsValueSet() {
		if c.TaxExemptReason.Value() != nil {
			structMap["tax_exempt_reason"] = c.TaxExemptReason.Value()
		} else {
			structMap["tax_exempt_reason"] = nil
		}
	}
	if c.DefaultAutoRenewalProfileId.IsValueSet() {
		if c.DefaultAutoRenewalProfileId.Value() != nil {
			structMap["default_auto_renewal_profile_id"] = c.DefaultAutoRenewalProfileId.Value()
		} else {
			structMap["default_auto_renewal_profile_id"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Customer.
// It customizes the JSON unmarshaling process for Customer objects.
func (c *Customer) UnmarshalJSON(input []byte) error {
	var temp tempCustomer
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "first_name", "last_name", "email", "cc_emails", "organization", "reference", "id", "created_at", "updated_at", "address", "address_2", "city", "state", "state_name", "zip", "country", "country_name", "phone", "verified", "portal_customer_created_at", "portal_invite_last_sent_at", "portal_invite_last_accepted_at", "tax_exempt", "vat_number", "parent_id", "locale", "default_subscription_group_uid", "salesforce_id", "tax_exempt_reason", "default_auto_renewal_profile_id")
	if err != nil {
		return err
	}
	c.AdditionalProperties = additionalProperties

	c.FirstName = temp.FirstName
	c.LastName = temp.LastName
	c.Email = temp.Email
	c.CcEmails = temp.CcEmails
	c.Organization = temp.Organization
	c.Reference = temp.Reference
	c.Id = temp.Id
	if temp.CreatedAt != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		c.CreatedAt = &CreatedAtVal
	}
	if temp.UpdatedAt != nil {
		UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
		if err != nil {
			log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
		}
		c.UpdatedAt = &UpdatedAtVal
	}
	c.Address = temp.Address
	c.Address2 = temp.Address2
	c.City = temp.City
	c.State = temp.State
	c.StateName = temp.StateName
	c.Zip = temp.Zip
	c.Country = temp.Country
	c.CountryName = temp.CountryName
	c.Phone = temp.Phone
	c.Verified = temp.Verified
	c.PortalCustomerCreatedAt.ShouldSetValue(temp.PortalCustomerCreatedAt.IsValueSet())
	if temp.PortalCustomerCreatedAt.Value() != nil {
		PortalCustomerCreatedAtVal, err := time.Parse(time.RFC3339, (*temp.PortalCustomerCreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse portal_customer_created_at as % s format.", time.RFC3339)
		}
		c.PortalCustomerCreatedAt.SetValue(&PortalCustomerCreatedAtVal)
	}
	c.PortalInviteLastSentAt.ShouldSetValue(temp.PortalInviteLastSentAt.IsValueSet())
	if temp.PortalInviteLastSentAt.Value() != nil {
		PortalInviteLastSentAtVal, err := time.Parse(time.RFC3339, (*temp.PortalInviteLastSentAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse portal_invite_last_sent_at as % s format.", time.RFC3339)
		}
		c.PortalInviteLastSentAt.SetValue(&PortalInviteLastSentAtVal)
	}
	c.PortalInviteLastAcceptedAt.ShouldSetValue(temp.PortalInviteLastAcceptedAt.IsValueSet())
	if temp.PortalInviteLastAcceptedAt.Value() != nil {
		PortalInviteLastAcceptedAtVal, err := time.Parse(time.RFC3339, (*temp.PortalInviteLastAcceptedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse portal_invite_last_accepted_at as % s format.", time.RFC3339)
		}
		c.PortalInviteLastAcceptedAt.SetValue(&PortalInviteLastAcceptedAtVal)
	}
	c.TaxExempt = temp.TaxExempt
	c.VatNumber = temp.VatNumber
	c.ParentId = temp.ParentId
	c.Locale = temp.Locale
	c.DefaultSubscriptionGroupUid = temp.DefaultSubscriptionGroupUid
	c.SalesforceId = temp.SalesforceId
	c.TaxExemptReason = temp.TaxExemptReason
	c.DefaultAutoRenewalProfileId = temp.DefaultAutoRenewalProfileId
	return nil
}

// tempCustomer is a temporary struct used for validating the fields of Customer.
type tempCustomer struct {
	FirstName                   *string          `json:"first_name,omitempty"`
	LastName                    *string          `json:"last_name,omitempty"`
	Email                       *string          `json:"email,omitempty"`
	CcEmails                    Optional[string] `json:"cc_emails"`
	Organization                Optional[string] `json:"organization"`
	Reference                   Optional[string] `json:"reference"`
	Id                          *int             `json:"id,omitempty"`
	CreatedAt                   *string          `json:"created_at,omitempty"`
	UpdatedAt                   *string          `json:"updated_at,omitempty"`
	Address                     Optional[string] `json:"address"`
	Address2                    Optional[string] `json:"address_2"`
	City                        Optional[string] `json:"city"`
	State                       Optional[string] `json:"state"`
	StateName                   Optional[string] `json:"state_name"`
	Zip                         Optional[string] `json:"zip"`
	Country                     Optional[string] `json:"country"`
	CountryName                 Optional[string] `json:"country_name"`
	Phone                       Optional[string] `json:"phone"`
	Verified                    Optional[bool]   `json:"verified"`
	PortalCustomerCreatedAt     Optional[string] `json:"portal_customer_created_at"`
	PortalInviteLastSentAt      Optional[string] `json:"portal_invite_last_sent_at"`
	PortalInviteLastAcceptedAt  Optional[string] `json:"portal_invite_last_accepted_at"`
	TaxExempt                   *bool            `json:"tax_exempt,omitempty"`
	VatNumber                   Optional[string] `json:"vat_number"`
	ParentId                    Optional[int]    `json:"parent_id"`
	Locale                      Optional[string] `json:"locale"`
	DefaultSubscriptionGroupUid Optional[string] `json:"default_subscription_group_uid"`
	SalesforceId                Optional[string] `json:"salesforce_id"`
	TaxExemptReason             Optional[string] `json:"tax_exempt_reason"`
	DefaultAutoRenewalProfileId Optional[int]    `json:"default_auto_renewal_profile_id"`
}
