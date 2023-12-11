package models

import (
	"encoding/json"
)

// InvoicePaymentApplication represents a InvoicePaymentApplication struct.
type InvoicePaymentApplication struct {
	// Unique identifier for the paid invoice. It has the prefix "inv_" followed by alphanumeric characters.
	InvoiceUid *string `json:"invoice_uid,omitempty"`
	// Unique identifier for the payment. It has the prefix "pmt_" followed by alphanumeric characters.
	ApplicationUid *string `json:"application_uid,omitempty"`
	// Dollar amount of the paid invoice.
	AppliedAmount *string `json:"applied_amount,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for InvoicePaymentApplication.
// It customizes the JSON marshaling process for InvoicePaymentApplication objects.
func (i *InvoicePaymentApplication) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InvoicePaymentApplication object to a map representation for JSON marshaling.
func (i *InvoicePaymentApplication) toMap() map[string]any {
	structMap := make(map[string]any)
	if i.InvoiceUid != nil {
		structMap["invoice_uid"] = i.InvoiceUid
	}
	if i.ApplicationUid != nil {
		structMap["application_uid"] = i.ApplicationUid
	}
	if i.AppliedAmount != nil {
		structMap["applied_amount"] = i.AppliedAmount
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoicePaymentApplication.
// It customizes the JSON unmarshaling process for InvoicePaymentApplication objects.
func (i *InvoicePaymentApplication) UnmarshalJSON(input []byte) error {
	temp := &struct {
		InvoiceUid     *string `json:"invoice_uid,omitempty"`
		ApplicationUid *string `json:"application_uid,omitempty"`
		AppliedAmount  *string `json:"applied_amount,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	i.InvoiceUid = temp.InvoiceUid
	i.ApplicationUid = temp.ApplicationUid
	i.AppliedAmount = temp.AppliedAmount
	return nil
}
