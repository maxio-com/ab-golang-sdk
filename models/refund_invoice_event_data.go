package models

import (
    "encoding/json"
    "errors"
    "log"
    "strings"
    "time"
)

// RefundInvoiceEventData represents a RefundInvoiceEventData struct.
// Example schema for an `refund_invoice` event
type RefundInvoiceEventData struct {
    // If true, credit was created and applied it to the invoice.
    ApplyCredit          bool                       `json:"apply_credit"`
    // Consolidation level of the invoice, which is applicable to invoice consolidation.  It will hold one of the following values:
    // * "none": A normal invoice with no consolidation.
    // * "child": An invoice segment which has been combined into a consolidated invoice.
    // * "parent": A consolidated invoice, whose contents are composed of invoice segments.
    // "Parent" invoices do not have lines of their own, but they have subtotals and totals which aggregate the member invoice segments.
    // See also the [invoice consolidation documentation](https://chargify.zendesk.com/hc/en-us/articles/4407746391835).
    ConsolidationLevel   *InvoiceConsolidationLevel `json:"consolidation_level,omitempty"`
    CreditNoteAttributes CreditNote                 `json:"credit_note_attributes"`
    // The refund memo.
    Memo                 *string                    `json:"memo,omitempty"`
    // The full, original amount of the refund.
    OriginalAmount       *string                    `json:"original_amount,omitempty"`
    // The ID of the payment transaction to be refunded.
    PaymentId            int                        `json:"payment_id"`
    // The amount of the refund.
    RefundAmount         string                     `json:"refund_amount"`
    // The ID of the refund transaction.
    RefundId             int                        `json:"refund_id"`
    // The time the refund was applied, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z"
    TransactionTime      time.Time                  `json:"transaction_time"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RefundInvoiceEventData.
// It customizes the JSON marshaling process for RefundInvoiceEventData objects.
func (r RefundInvoiceEventData) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RefundInvoiceEventData object to a map representation for JSON marshaling.
func (r RefundInvoiceEventData) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["apply_credit"] = r.ApplyCredit
    if r.ConsolidationLevel != nil {
        structMap["consolidation_level"] = r.ConsolidationLevel
    }
    structMap["credit_note_attributes"] = r.CreditNoteAttributes.toMap()
    if r.Memo != nil {
        structMap["memo"] = r.Memo
    }
    if r.OriginalAmount != nil {
        structMap["original_amount"] = r.OriginalAmount
    }
    structMap["payment_id"] = r.PaymentId
    structMap["refund_amount"] = r.RefundAmount
    structMap["refund_id"] = r.RefundId
    structMap["transaction_time"] = r.TransactionTime.Format(time.RFC3339)
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundInvoiceEventData.
// It customizes the JSON unmarshaling process for RefundInvoiceEventData objects.
func (r *RefundInvoiceEventData) UnmarshalJSON(input []byte) error {
    var temp refundInvoiceEventData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "apply_credit", "consolidation_level", "credit_note_attributes", "memo", "original_amount", "payment_id", "refund_amount", "refund_id", "transaction_time")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.ApplyCredit = *temp.ApplyCredit
    r.ConsolidationLevel = temp.ConsolidationLevel
    r.CreditNoteAttributes = *temp.CreditNoteAttributes
    r.Memo = temp.Memo
    r.OriginalAmount = temp.OriginalAmount
    r.PaymentId = *temp.PaymentId
    r.RefundAmount = *temp.RefundAmount
    r.RefundId = *temp.RefundId
    TransactionTimeVal, err := time.Parse(time.RFC3339, *temp.TransactionTime)
    if err != nil {
        log.Fatalf("Cannot Parse transaction_time as % s format.", time.RFC3339)
    }
    r.TransactionTime = TransactionTimeVal
    return nil
}

// TODO
type refundInvoiceEventData  struct {
    ApplyCredit          *bool                      `json:"apply_credit"`
    ConsolidationLevel   *InvoiceConsolidationLevel `json:"consolidation_level,omitempty"`
    CreditNoteAttributes *CreditNote                `json:"credit_note_attributes"`
    Memo                 *string                    `json:"memo,omitempty"`
    OriginalAmount       *string                    `json:"original_amount,omitempty"`
    PaymentId            *int                       `json:"payment_id"`
    RefundAmount         *string                    `json:"refund_amount"`
    RefundId             *int                       `json:"refund_id"`
    TransactionTime      *string                    `json:"transaction_time"`
}

func (r *refundInvoiceEventData) validate() error {
    var errs []string
    if r.ApplyCredit == nil {
        errs = append(errs, "required field `apply_credit` is missing for type `Refund Invoice Event Data`")
    }
    if r.CreditNoteAttributes == nil {
        errs = append(errs, "required field `credit_note_attributes` is missing for type `Refund Invoice Event Data`")
    }
    if r.PaymentId == nil {
        errs = append(errs, "required field `payment_id` is missing for type `Refund Invoice Event Data`")
    }
    if r.RefundAmount == nil {
        errs = append(errs, "required field `refund_amount` is missing for type `Refund Invoice Event Data`")
    }
    if r.RefundId == nil {
        errs = append(errs, "required field `refund_id` is missing for type `Refund Invoice Event Data`")
    }
    if r.TransactionTime == nil {
        errs = append(errs, "required field `transaction_time` is missing for type `Refund Invoice Event Data`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
