package models

import (
	"encoding/json"
)

// InvoiceCustomer represents a InvoiceCustomer struct.
// Information about the customer who is owner or recipient the invoiced subscription.
type InvoiceCustomer struct {
	ChargifyId   *int             `json:"chargify_id,omitempty"`
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
		structMap["organization"] = i.Organization.Value()
	}
	if i.Email != nil {
		structMap["email"] = i.Email
	}
	if i.VatNumber.IsValueSet() {
		structMap["vat_number"] = i.VatNumber.Value()
	}
	if i.Reference.IsValueSet() {
		structMap["reference"] = i.Reference.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceCustomer.
// It customizes the JSON unmarshaling process for InvoiceCustomer objects.
func (i *InvoiceCustomer) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ChargifyId   *int             `json:"chargify_id,omitempty"`
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

	i.ChargifyId = temp.ChargifyId
	i.FirstName = temp.FirstName
	i.LastName = temp.LastName
	i.Organization = temp.Organization
	i.Email = temp.Email
	i.VatNumber = temp.VatNumber
	i.Reference = temp.Reference
	return nil
}
