package models

import (
    "encoding/json"
)

// Customer represents a Customer struct.
type Customer struct {
    // The first name of the customer
    FirstName                   *string          `json:"first_name,omitempty"`
    // The last name of the customer
    LastName                    *string          `json:"last_name,omitempty"`
    // The email address of the customer
    Email                       *string          `json:"email,omitempty"`
    // A comma-separated list of emails that should be cc’d on all customer communications (i.e. “joe@example.com, sue@example.com”)
    CcEmails                    Optional[string] `json:"cc_emails"`
    // The organization of the customer
    Organization                Optional[string] `json:"organization"`
    // The unique identifier used within your own application for this customer
    Reference                   Optional[string] `json:"reference"`
    // The customer ID in Chargify
    Id                          *int             `json:"id,omitempty"`
    // The timestamp in which the customer object was created in Chargify
    CreatedAt                   *string          `json:"created_at,omitempty"`
    // The timestamp in which the customer object was last edited
    UpdatedAt                   *string          `json:"updated_at,omitempty"`
    // The customer’s shipping street address (i.e. “123 Main St.”)
    Address                     Optional[string] `json:"address"`
    // Second line of the customer’s shipping address i.e. “Apt. 100”
    Address2                    Optional[string] `json:"address_2"`
    // The customer’s shipping address city (i.e. “Boston”)
    City                        Optional[string] `json:"city"`
    // The customer’s shipping address state (i.e. “MA”)
    State                       Optional[string] `json:"state"`
    // The customer's full name of state
    StateName                   Optional[string] `json:"state_name"`
    // The customer’s shipping address zip code (i.e. “12345”)
    Zip                         Optional[string] `json:"zip"`
    // The customer shipping address country
    Country                     Optional[string] `json:"country"`
    // The customer's full name of country
    CountryName                 Optional[string] `json:"country_name"`
    // The phone number of the customer
    Phone                       Optional[string] `json:"phone"`
    // Is the customer verified to use ACH as a payment method. Available only on Authorize.Net gateway
    Verified                    Optional[bool]   `json:"verified"`
    // The timestamp of when the Billing Portal entry was created at for the customer
    PortalCustomerCreatedAt     Optional[string] `json:"portal_customer_created_at"`
    // The timestamp of when the Billing Portal invite was last sent at
    PortalInviteLastSentAt      Optional[string] `json:"portal_invite_last_sent_at"`
    // The timestamp of when the Billing Portal invite was last accepted
    PortalInviteLastAcceptedAt  Optional[string] `json:"portal_invite_last_accepted_at"`
    // The tax exempt status for the customer. Acceptable values are true or 1 for true and false or 0 for false.
    TaxExempt                   *bool            `json:"tax_exempt,omitempty"`
    // The VAT business identification number for the customer. This number is used to determine VAT tax opt out rules. It is not validated when added or updated on a customer record. Instead, it is validated via VIES before calculating taxes. Only valid business identification numbers will allow for VAT opt out.
    VatNumber                   Optional[string] `json:"vat_number"`
    // The parent ID in Chargify if applicable. Parent is another Customer object.
    ParentId                    Optional[int]    `json:"parent_id"`
    // The locale for the customer to identify language-region
    Locale                      Optional[string] `json:"locale"`
    DefaultSubscriptionGroupUid Optional[string] `json:"default_subscription_group_uid"`
}

// MarshalJSON implements the json.Marshaler interface for Customer.
// It customizes the JSON marshaling process for Customer objects.
func (c *Customer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the Customer object to a map representation for JSON marshaling.
func (c *Customer) toMap() map[string]any {
    structMap := make(map[string]any)
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
        structMap["cc_emails"] = c.CcEmails.Value()
    }
    if c.Organization.IsValueSet() {
        structMap["organization"] = c.Organization.Value()
    }
    if c.Reference.IsValueSet() {
        structMap["reference"] = c.Reference.Value()
    }
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.CreatedAt != nil {
        structMap["created_at"] = c.CreatedAt
    }
    if c.UpdatedAt != nil {
        structMap["updated_at"] = c.UpdatedAt
    }
    if c.Address.IsValueSet() {
        structMap["address"] = c.Address.Value()
    }
    if c.Address2.IsValueSet() {
        structMap["address_2"] = c.Address2.Value()
    }
    if c.City.IsValueSet() {
        structMap["city"] = c.City.Value()
    }
    if c.State.IsValueSet() {
        structMap["state"] = c.State.Value()
    }
    if c.StateName.IsValueSet() {
        structMap["state_name"] = c.StateName.Value()
    }
    if c.Zip.IsValueSet() {
        structMap["zip"] = c.Zip.Value()
    }
    if c.Country.IsValueSet() {
        structMap["country"] = c.Country.Value()
    }
    if c.CountryName.IsValueSet() {
        structMap["country_name"] = c.CountryName.Value()
    }
    if c.Phone.IsValueSet() {
        structMap["phone"] = c.Phone.Value()
    }
    if c.Verified.IsValueSet() {
        structMap["verified"] = c.Verified.Value()
    }
    if c.PortalCustomerCreatedAt.IsValueSet() {
        structMap["portal_customer_created_at"] = c.PortalCustomerCreatedAt.Value()
    }
    if c.PortalInviteLastSentAt.IsValueSet() {
        structMap["portal_invite_last_sent_at"] = c.PortalInviteLastSentAt.Value()
    }
    if c.PortalInviteLastAcceptedAt.IsValueSet() {
        structMap["portal_invite_last_accepted_at"] = c.PortalInviteLastAcceptedAt.Value()
    }
    if c.TaxExempt != nil {
        structMap["tax_exempt"] = c.TaxExempt
    }
    if c.VatNumber.IsValueSet() {
        structMap["vat_number"] = c.VatNumber.Value()
    }
    if c.ParentId.IsValueSet() {
        structMap["parent_id"] = c.ParentId.Value()
    }
    if c.Locale.IsValueSet() {
        structMap["locale"] = c.Locale.Value()
    }
    if c.DefaultSubscriptionGroupUid.IsValueSet() {
        structMap["default_subscription_group_uid"] = c.DefaultSubscriptionGroupUid.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Customer.
// It customizes the JSON unmarshaling process for Customer objects.
func (c *Customer) UnmarshalJSON(input []byte) error {
    temp := &struct {
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
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.FirstName = temp.FirstName
    c.LastName = temp.LastName
    c.Email = temp.Email
    c.CcEmails = temp.CcEmails
    c.Organization = temp.Organization
    c.Reference = temp.Reference
    c.Id = temp.Id
    c.CreatedAt = temp.CreatedAt
    c.UpdatedAt = temp.UpdatedAt
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
    c.PortalCustomerCreatedAt = temp.PortalCustomerCreatedAt
    c.PortalInviteLastSentAt = temp.PortalInviteLastSentAt
    c.PortalInviteLastAcceptedAt = temp.PortalInviteLastAcceptedAt
    c.TaxExempt = temp.TaxExempt
    c.VatNumber = temp.VatNumber
    c.ParentId = temp.ParentId
    c.Locale = temp.Locale
    c.DefaultSubscriptionGroupUid = temp.DefaultSubscriptionGroupUid
    return nil
}
