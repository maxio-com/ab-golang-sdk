/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// MultiInvoicePayment represents a MultiInvoicePayment struct.
type MultiInvoicePayment struct {
    // The numeric ID of the transaction.
    TransactionId        *int                        `json:"transaction_id,omitempty"`
    // Dollar amount of the sum of the paid invoices.
    TotalAmount          *string                     `json:"total_amount,omitempty"`
    // The ISO 4217 currency code (3 character string) representing the currency of invoice transaction.
    CurrencyCode         *string                     `json:"currency_code,omitempty"`
    Applications         []InvoicePaymentApplication `json:"applications,omitempty"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MultiInvoicePayment.
// It customizes the JSON marshaling process for MultiInvoicePayment objects.
func (m MultiInvoicePayment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "transaction_id", "total_amount", "currency_code", "applications"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MultiInvoicePayment object to a map representation for JSON marshaling.
func (m MultiInvoicePayment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
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
    var temp tempMultiInvoicePayment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "transaction_id", "total_amount", "currency_code", "applications")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.TransactionId = temp.TransactionId
    m.TotalAmount = temp.TotalAmount
    m.CurrencyCode = temp.CurrencyCode
    m.Applications = temp.Applications
    return nil
}

// tempMultiInvoicePayment is a temporary struct used for validating the fields of MultiInvoicePayment.
type tempMultiInvoicePayment  struct {
    TransactionId *int                        `json:"transaction_id,omitempty"`
    TotalAmount   *string                     `json:"total_amount,omitempty"`
    CurrencyCode  *string                     `json:"currency_code,omitempty"`
    Applications  []InvoicePaymentApplication `json:"applications,omitempty"`
}
