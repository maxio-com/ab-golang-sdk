package models

import (
	"encoding/json"
)

// InvoicePayer represents a InvoicePayer struct.
type InvoicePayer struct {
	ChargifyId   *int             `json:"chargify_id,omitempty"`
	FirstName    *string          `json:"first_name,omitempty"`
	LastName     *string          `json:"last_name,omitempty"`
	Organization Optional[string] `json:"organization"`
	Email        *string          `json:"email,omitempty"`
	VatNumber    Optional[string] `json:"vat_number"`
}

// MarshalJSON implements the json.Marshaler interface for InvoicePayer.
// It customizes the JSON marshaling process for InvoicePayer objects.
func (i *InvoicePayer) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InvoicePayer object to a map representation for JSON marshaling.
func (i *InvoicePayer) toMap() map[string]any {
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
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoicePayer.
// It customizes the JSON unmarshaling process for InvoicePayer objects.
func (i *InvoicePayer) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ChargifyId   *int             `json:"chargify_id,omitempty"`
		FirstName    *string          `json:"first_name,omitempty"`
		LastName     *string          `json:"last_name,omitempty"`
		Organization Optional[string] `json:"organization"`
		Email        *string          `json:"email,omitempty"`
		VatNumber    Optional[string] `json:"vat_number"`
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
	return nil
}
