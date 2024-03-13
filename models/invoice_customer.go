package models

import (
	"encoding/json"
)

// InvoiceCustomer represents a InvoiceCustomer struct.
// Information about the customer who is owner or recipient the invoiced subscription.
type InvoiceCustomer struct {
	ChargifyId   Optional[int]    `json:"chargify_id"`
	FirstName    *string          `json:"first_name,omitempty"`
	LastName     *string          `json:"last_name,omitempty"`
	Organization Optional[string] `json:"organization"`
	Email        *string          `json:"email,omitempty"`
	VatNumber    Optional[string] `json:"vat_number"`
	Reference    Optional[string] `json:"reference"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceCustomer.
// It customizes the JSON marshaling process for InvoiceCustomer objects.
func (i *InvoiceCustomer) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InvoiceCustomer object to a map representation for JSON marshaling.
func (i *InvoiceCustomer) toMap() map[string]any {
	structMap := make(map[string]any)
	if i.ChargifyId.IsValueSet() {
		if i.ChargifyId.Value() != nil {
			structMap["chargify_id"] = i.ChargifyId.Value()
		} else {
			structMap["chargify_id"] = nil
		}
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
	if i.Reference.IsValueSet() {
		if i.Reference.Value() != nil {
			structMap["reference"] = i.Reference.Value()
		} else {
			structMap["reference"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceCustomer.
// It customizes the JSON unmarshaling process for InvoiceCustomer objects.
func (i *InvoiceCustomer) UnmarshalJSON(input []byte) error {
	var temp invoiceCustomer
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	i.ChargifyId = temp.ChargifyId
	i.FirstName = temp.FirstName
	i.LastName = temp.LastName
	i.Organization = temp.Organization
	i.Email = temp.Email
	i.VatNumber = temp.VatNumber
	i.Reference = temp.Reference
	return nil
}

// TODO
type invoiceCustomer struct {
	ChargifyId   Optional[int]    `json:"chargify_id"`
	FirstName    *string          `json:"first_name,omitempty"`
	LastName     *string          `json:"last_name,omitempty"`
	Organization Optional[string] `json:"organization"`
	Email        *string          `json:"email,omitempty"`
	VatNumber    Optional[string] `json:"vat_number"`
	Reference    Optional[string] `json:"reference"`
}
