package models

import (
    "encoding/json"
    "errors"
    "log"
    "strings"
    "time"
)

// ApplyPaymentEventData represents a ApplyPaymentEventData struct.
// Example schema for an `apply_payment` event
type ApplyPaymentEventData struct {
    ConsolidationLevel        InvoiceConsolidationLevel `json:"consolidation_level"`
    // The payment memo
    Memo                      string                    `json:"memo"`
    // The full, original amount of the payment transaction as a string in full units. Incoming payments can be split amongst several invoices, which will result in a `applied_amount` less than the `original_amount`. Example: A $100.99 payment, of which $40.11 is applied to this invoice, will have an `original_amount` of `"100.99"`.
    OriginalAmount            string                    `json:"original_amount"`
    // The amount of the payment applied to this invoice. Incoming payments can be split amongst several invoices, which will result in a `applied_amount` less than the `original_amount`. Example: A $100.99 payment, of which $40.11 is applied to this invoice, will have an `applied_amount` of `"40.11"`.
    AppliedAmount             string                    `json:"applied_amount"`
    // The time the payment was applied, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z"
    TransactionTime           time.Time                 `json:"transaction_time"`
    // A nested data structure detailing the method of payment
    PaymentMethod             InvoiceEventPayment       `json:"payment_method"`
    // The Chargify id of the original payment
    TransactionId             *int                      `json:"transaction_id,omitempty"`
    ParentInvoiceNumber       Optional[int]             `json:"parent_invoice_number"`
    RemainingPrepaymentAmount Optional[string]          `json:"remaining_prepayment_amount"`
    Prepayment                *bool                     `json:"prepayment,omitempty"`
    External                  *bool                     `json:"external,omitempty"`
    AdditionalProperties      map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApplyPaymentEventData.
// It customizes the JSON marshaling process for ApplyPaymentEventData objects.
func (a ApplyPaymentEventData) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApplyPaymentEventData object to a map representation for JSON marshaling.
func (a ApplyPaymentEventData) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["consolidation_level"] = a.ConsolidationLevel
    structMap["memo"] = a.Memo
    structMap["original_amount"] = a.OriginalAmount
    structMap["applied_amount"] = a.AppliedAmount
    structMap["transaction_time"] = a.TransactionTime.Format(time.RFC3339)
    structMap["payment_method"] = a.PaymentMethod.toMap()
    if a.TransactionId != nil {
        structMap["transaction_id"] = a.TransactionId
    }
    if a.ParentInvoiceNumber.IsValueSet() {
        if a.ParentInvoiceNumber.Value() != nil {
            structMap["parent_invoice_number"] = a.ParentInvoiceNumber.Value()
        } else {
            structMap["parent_invoice_number"] = nil
        }
    }
    if a.RemainingPrepaymentAmount.IsValueSet() {
        if a.RemainingPrepaymentAmount.Value() != nil {
            structMap["remaining_prepayment_amount"] = a.RemainingPrepaymentAmount.Value()
        } else {
            structMap["remaining_prepayment_amount"] = nil
        }
    }
    if a.Prepayment != nil {
        structMap["prepayment"] = a.Prepayment
    }
    if a.External != nil {
        structMap["external"] = a.External
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApplyPaymentEventData.
// It customizes the JSON unmarshaling process for ApplyPaymentEventData objects.
func (a *ApplyPaymentEventData) UnmarshalJSON(input []byte) error {
    var temp applyPaymentEventData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "consolidation_level", "memo", "original_amount", "applied_amount", "transaction_time", "payment_method", "transaction_id", "parent_invoice_number", "remaining_prepayment_amount", "prepayment", "external")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.ConsolidationLevel = *temp.ConsolidationLevel
    a.Memo = *temp.Memo
    a.OriginalAmount = *temp.OriginalAmount
    a.AppliedAmount = *temp.AppliedAmount
    TransactionTimeVal, err := time.Parse(time.RFC3339, *temp.TransactionTime)
    if err != nil {
        log.Fatalf("Cannot Parse transaction_time as % s format.", time.RFC3339)
    }
    a.TransactionTime = TransactionTimeVal
    a.PaymentMethod = *temp.PaymentMethod
    a.TransactionId = temp.TransactionId
    a.ParentInvoiceNumber = temp.ParentInvoiceNumber
    a.RemainingPrepaymentAmount = temp.RemainingPrepaymentAmount
    a.Prepayment = temp.Prepayment
    a.External = temp.External
    return nil
}

// applyPaymentEventData is a temporary struct used for validating the fields of ApplyPaymentEventData.
type applyPaymentEventData  struct {
    ConsolidationLevel        *InvoiceConsolidationLevel `json:"consolidation_level"`
    Memo                      *string                    `json:"memo"`
    OriginalAmount            *string                    `json:"original_amount"`
    AppliedAmount             *string                    `json:"applied_amount"`
    TransactionTime           *string                    `json:"transaction_time"`
    PaymentMethod             *InvoiceEventPayment       `json:"payment_method"`
    TransactionId             *int                       `json:"transaction_id,omitempty"`
    ParentInvoiceNumber       Optional[int]              `json:"parent_invoice_number"`
    RemainingPrepaymentAmount Optional[string]           `json:"remaining_prepayment_amount"`
    Prepayment                *bool                      `json:"prepayment,omitempty"`
    External                  *bool                      `json:"external,omitempty"`
}

func (a *applyPaymentEventData) validate() error {
    var errs []string
    if a.ConsolidationLevel == nil {
        errs = append(errs, "required field `consolidation_level` is missing for type `Apply Payment Event Data`")
    }
    if a.Memo == nil {
        errs = append(errs, "required field `memo` is missing for type `Apply Payment Event Data`")
    }
    if a.OriginalAmount == nil {
        errs = append(errs, "required field `original_amount` is missing for type `Apply Payment Event Data`")
    }
    if a.AppliedAmount == nil {
        errs = append(errs, "required field `applied_amount` is missing for type `Apply Payment Event Data`")
    }
    if a.TransactionTime == nil {
        errs = append(errs, "required field `transaction_time` is missing for type `Apply Payment Event Data`")
    }
    if a.PaymentMethod == nil {
        errs = append(errs, "required field `payment_method` is missing for type `Apply Payment Event Data`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
