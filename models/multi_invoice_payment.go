package models

import (
	"encoding/json"
)

// MultiInvoicePayment represents a MultiInvoicePayment struct.
type MultiInvoicePayment struct {
	// The numeric ID of the transaction.
	TransactionId *int `json:"transaction_id,omitempty"`
	// Dollar amount of the sum of the paid invoices.
	TotalAmount *string `json:"total_amount,omitempty"`
	// The ISO 4217 currency code (3 character string) representing the currency of invoice transaction.
	CurrencyCode *string                     `json:"currency_code,omitempty"`
	Applications []InvoicePaymentApplication `json:"applications,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for MultiInvoicePayment.
// It customizes the JSON marshaling process for MultiInvoicePayment objects.
func (m *MultiInvoicePayment) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the MultiInvoicePayment object to a map representation for JSON marshaling.
func (m *MultiInvoicePayment) toMap() map[string]any {
	structMap := make(map[string]any)
	if m.TransactionId != nil {
		structMap["transaction_id"] = m.TransactionId
	}
	if m.TotalAmount != nil {
		structMap["total_amount"] = m.TotalAmount
	}
	if m.CurrencyCode != nil {
		structMap["currency_code"] = m.CurrencyCode
	}
	if m.Applications != nil {
		structMap["applications"] = m.Applications
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MultiInvoicePayment.
// It customizes the JSON unmarshaling process for MultiInvoicePayment objects.
func (m *MultiInvoicePayment) UnmarshalJSON(input []byte) error {
	temp := &struct {
		TransactionId *int                        `json:"transaction_id,omitempty"`
		TotalAmount   *string                     `json:"total_amount,omitempty"`
		CurrencyCode  *string                     `json:"currency_code,omitempty"`
		Applications  []InvoicePaymentApplication `json:"applications,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.TransactionId = temp.TransactionId
	m.TotalAmount = temp.TotalAmount
	m.CurrencyCode = temp.CurrencyCode
	m.Applications = temp.Applications
	return nil
}
