/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// RecordPaymentResponse represents a RecordPaymentResponse struct.
type RecordPaymentResponse struct {
    PaidInvoices         []PaidInvoice               `json:"paid_invoices,omitempty"`
    Prepayment           Optional[InvoicePrePayment] `json:"prepayment"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for RecordPaymentResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RecordPaymentResponse) String() string {
    return fmt.Sprintf(
    	"RecordPaymentResponse[PaidInvoices=%v, Prepayment=%v, AdditionalProperties=%v]",
    	r.PaidInvoices, r.Prepayment, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RecordPaymentResponse.
// It customizes the JSON marshaling process for RecordPaymentResponse objects.
func (r RecordPaymentResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "paid_invoices", "prepayment"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RecordPaymentResponse object to a map representation for JSON marshaling.
func (r RecordPaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.PaidInvoices != nil {
        structMap["paid_invoices"] = r.PaidInvoices
    }
    if r.Prepayment.IsValueSet() {
        if r.Prepayment.Value() != nil {
            structMap["prepayment"] = r.Prepayment.Value().toMap()
        } else {
            structMap["prepayment"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RecordPaymentResponse.
// It customizes the JSON unmarshaling process for RecordPaymentResponse objects.
func (r *RecordPaymentResponse) UnmarshalJSON(input []byte) error {
    var temp tempRecordPaymentResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "paid_invoices", "prepayment")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.PaidInvoices = temp.PaidInvoices
    r.Prepayment = temp.Prepayment
    return nil
}

// tempRecordPaymentResponse is a temporary struct used for validating the fields of RecordPaymentResponse.
type tempRecordPaymentResponse  struct {
    PaidInvoices []PaidInvoice               `json:"paid_invoices,omitempty"`
    Prepayment   Optional[InvoicePrePayment] `json:"prepayment"`
}
