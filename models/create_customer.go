package models

import (
	"encoding/json"
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
}

// MarshalJSON implements the json.Marshaler interface for CreateCustomer.
// It customizes the JSON marshaling process for CreateCustomer objects.
func (c *CreateCustomer) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCustomer object to a map representation for JSON marshaling.
func (c *CreateCustomer) toMap() map[string]any {
	structMap := make(map[string]any)
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
		structMap["parent_id"] = c.ParentId.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCustomer.
// It customizes the JSON unmarshaling process for CreateCustomer objects.
func (c *CreateCustomer) UnmarshalJSON(input []byte) error {
	temp := &struct {
		FirstName       string        `json:"first_name"`
		LastName        string        `json:"last_name"`
		Email           string        `json:"email"`
		CcEmails        *string       `json:"cc_emails,omitempty"`
		Organization    *string       `json:"organization,omitempty"`
		Reference       *string       `json:"reference,omitempty"`
		Address         *string       `json:"address,omitempty"`
		Address2        *string       `json:"address_2,omitempty"`
		City            *string       `json:"city,omitempty"`
		State           *string       `json:"state,omitempty"`
		Zip             *string       `json:"zip,omitempty"`
		Country         *string       `json:"country,omitempty"`
		Phone           *string       `json:"phone,omitempty"`
		Locale          *string       `json:"locale,omitempty"`
		VatNumber       *string       `json:"vat_number,omitempty"`
		TaxExempt       *bool         `json:"tax_exempt,omitempty"`
		TaxExemptReason *string       `json:"tax_exempt_reason,omitempty"`
		ParentId        Optional[int] `json:"parent_id"`
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
	return nil
}
