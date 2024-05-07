package models

import (
    "encoding/json"
    "log"
    "time"
)

// InvoiceEventData represents a InvoiceEventData struct.
// The event data is the data that, when combined with the command, results in the output invoice found in the invoice field.
type InvoiceEventData struct {
    // Unique identifier for the credit note application. It is generated automatically by Chargify and has the prefix "cdt_" followed by alphanumeric characters.
    Uid                       *string                        `json:"uid,omitempty"`
    // A unique, identifying string that appears on the credit note and in places it is referenced.
    CreditNoteNumber          *string                        `json:"credit_note_number,omitempty"`
    // Unique identifier for the credit note. It is generated automatically by Chargify and has the prefix "cn_" followed by alphanumeric characters.
    CreditNoteUid             *string                        `json:"credit_note_uid,omitempty"`
    // The full, original amount of the credit note.
    OriginalAmount            *string                        `json:"original_amount,omitempty"`
    // The amount of the credit note applied to invoice.
    AppliedAmount             *string                        `json:"applied_amount,omitempty"`
    // The time the credit note was applied, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z"
    TransactionTime           *time.Time                     `json:"transaction_time,omitempty"`
    // The credit note memo.
    Memo                      *string                        `json:"memo,omitempty"`
    // The role of the credit note (e.g. 'general')
    Role                      *string                        `json:"role,omitempty"`
    // Shows whether it was applied to consolidated invoice or not
    ConsolidatedInvoice       *bool                          `json:"consolidated_invoice,omitempty"`
    // List of credit notes applied to children invoices (if consolidated invoice)
    AppliedCreditNotes        []AppliedCreditNoteData        `json:"applied_credit_notes,omitempty"`
    // A unique, identifying string that appears on the debit note and in places it is referenced.
    DebitNoteNumber           *string                        `json:"debit_note_number,omitempty"`
    // Unique identifier for the debit note. It is generated automatically by Chargify and has the prefix "db_" followed by alphanumeric characters.
    DebitNoteUid              *string                        `json:"debit_note_uid,omitempty"`
    // A nested data structure detailing the method of payment
    PaymentMethod             *InvoiceEventDataPaymentMethod `json:"payment_method,omitempty"`
    // The Chargify id of the original payment
    TransactionId             *int                           `json:"transaction_id,omitempty"`
    ParentInvoiceNumber       Optional[int]                  `json:"parent_invoice_number"`
    RemainingPrepaymentAmount Optional[string]               `json:"remaining_prepayment_amount"`
    // The flag that shows whether the original payment was a prepayment or not
    Prepayment                *bool                          `json:"prepayment,omitempty"`
    External                  *bool                          `json:"external,omitempty"`
    // The previous collection method of the invoice.
    FromCollectionMethod      *string                        `json:"from_collection_method,omitempty"`
    // The new collection method of the invoice.
    ToCollectionMethod        *string                        `json:"to_collection_method,omitempty"`
    // Consolidation level of the invoice, which is applicable to invoice consolidation.  It will hold one of the following values:
    // * "none": A normal invoice with no consolidation.
    // * "child": An invoice segment which has been combined into a consolidated invoice.
    // * "parent": A consolidated invoice, whose contents are composed of invoice segments.
    // "Parent" invoices do not have lines of their own, but they have subtotals and totals which aggregate the member invoice segments.
    // See also the [invoice consolidation documentation](https://chargify.zendesk.com/hc/en-us/articles/4407746391835).
    ConsolidationLevel        *InvoiceConsolidationLevel     `json:"consolidation_level,omitempty"`
    // The status of the invoice before event occurence. See [Invoice Statuses](https://chargify.zendesk.com/hc/en-us/articles/4407737494171#line-item-breakdowns) for more.
    FromStatus                *InvoiceStatus                 `json:"from_status,omitempty"`
    // The status of the invoice after event occurence. See [Invoice Statuses](https://chargify.zendesk.com/hc/en-us/articles/4407737494171#line-item-breakdowns) for more.
    ToStatus                  *InvoiceStatus                 `json:"to_status,omitempty"`
    // Amount due on the invoice, which is `total_amount - credit_amount - paid_amount`.
    DueAmount                 *string                        `json:"due_amount,omitempty"`
    // The invoice total, which is `subtotal_amount - discount_amount + tax_amount`.'
    TotalAmount               *string                        `json:"total_amount,omitempty"`
    // If true, credit was created and applied it to the invoice.
    ApplyCredit               *bool                          `json:"apply_credit,omitempty"`
    CreditNoteAttributes      *CreditNote1                   `json:"credit_note_attributes,omitempty"`
    // The ID of the payment transaction to be refunded.
    PaymentId                 *int                           `json:"payment_id,omitempty"`
    // The amount of the refund.
    RefundAmount              *string                        `json:"refund_amount,omitempty"`
    // The ID of the refund transaction.
    RefundId                  *int                           `json:"refund_id,omitempty"`
    // If true, the invoice is an advance invoice.
    IsAdvanceInvoice          *bool                          `json:"is_advance_invoice,omitempty"`
    // The reason for the void.
    Reason                    *string                        `json:"reason,omitempty"`
    AdditionalProperties      map[string]any                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceEventData.
// It customizes the JSON marshaling process for InvoiceEventData objects.
func (i InvoiceEventData) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceEventData object to a map representation for JSON marshaling.
func (i InvoiceEventData) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Uid != nil {
        structMap["uid"] = i.Uid
    }
    if i.CreditNoteNumber != nil {
        structMap["credit_note_number"] = i.CreditNoteNumber
    }
    if i.CreditNoteUid != nil {
        structMap["credit_note_uid"] = i.CreditNoteUid
    }
    if i.OriginalAmount != nil {
        structMap["original_amount"] = i.OriginalAmount
    }
    if i.AppliedAmount != nil {
        structMap["applied_amount"] = i.AppliedAmount
    }
    if i.TransactionTime != nil {
        structMap["transaction_time"] = i.TransactionTime.Format(time.RFC3339)
    }
    if i.Memo != nil {
        structMap["memo"] = i.Memo
    }
    if i.Role != nil {
        structMap["role"] = i.Role
    }
    if i.ConsolidatedInvoice != nil {
        structMap["consolidated_invoice"] = i.ConsolidatedInvoice
    }
    if i.AppliedCreditNotes != nil {
        structMap["applied_credit_notes"] = i.AppliedCreditNotes
    }
    if i.DebitNoteNumber != nil {
        structMap["debit_note_number"] = i.DebitNoteNumber
    }
    if i.DebitNoteUid != nil {
        structMap["debit_note_uid"] = i.DebitNoteUid
    }
    if i.PaymentMethod != nil {
        structMap["payment_method"] = i.PaymentMethod.toMap()
    }
    if i.TransactionId != nil {
        structMap["transaction_id"] = i.TransactionId
    }
    if i.ParentInvoiceNumber.IsValueSet() {
        if i.ParentInvoiceNumber.Value() != nil {
            structMap["parent_invoice_number"] = i.ParentInvoiceNumber.Value()
        } else {
            structMap["parent_invoice_number"] = nil
        }
    }
    if i.RemainingPrepaymentAmount.IsValueSet() {
        if i.RemainingPrepaymentAmount.Value() != nil {
            structMap["remaining_prepayment_amount"] = i.RemainingPrepaymentAmount.Value()
        } else {
            structMap["remaining_prepayment_amount"] = nil
        }
    }
    if i.Prepayment != nil {
        structMap["prepayment"] = i.Prepayment
    }
    if i.External != nil {
        structMap["external"] = i.External
    }
    if i.FromCollectionMethod != nil {
        structMap["from_collection_method"] = i.FromCollectionMethod
    }
    if i.ToCollectionMethod != nil {
        structMap["to_collection_method"] = i.ToCollectionMethod
    }
    if i.ConsolidationLevel != nil {
        structMap["consolidation_level"] = i.ConsolidationLevel
    }
    if i.FromStatus != nil {
        structMap["from_status"] = i.FromStatus
    }
    if i.ToStatus != nil {
        structMap["to_status"] = i.ToStatus
    }
    if i.DueAmount != nil {
        structMap["due_amount"] = i.DueAmount
    }
    if i.TotalAmount != nil {
        structMap["total_amount"] = i.TotalAmount
    }
    if i.ApplyCredit != nil {
        structMap["apply_credit"] = i.ApplyCredit
    }
    if i.CreditNoteAttributes != nil {
        structMap["credit_note_attributes"] = i.CreditNoteAttributes.toMap()
    }
    if i.PaymentId != nil {
        structMap["payment_id"] = i.PaymentId
    }
    if i.RefundAmount != nil {
        structMap["refund_amount"] = i.RefundAmount
    }
    if i.RefundId != nil {
        structMap["refund_id"] = i.RefundId
    }
    if i.IsAdvanceInvoice != nil {
        structMap["is_advance_invoice"] = i.IsAdvanceInvoice
    }
    if i.Reason != nil {
        structMap["reason"] = i.Reason
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceEventData.
// It customizes the JSON unmarshaling process for InvoiceEventData objects.
func (i *InvoiceEventData) UnmarshalJSON(input []byte) error {
    var temp invoiceEventData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "credit_note_number", "credit_note_uid", "original_amount", "applied_amount", "transaction_time", "memo", "role", "consolidated_invoice", "applied_credit_notes", "debit_note_number", "debit_note_uid", "payment_method", "transaction_id", "parent_invoice_number", "remaining_prepayment_amount", "prepayment", "external", "from_collection_method", "to_collection_method", "consolidation_level", "from_status", "to_status", "due_amount", "total_amount", "apply_credit", "credit_note_attributes", "payment_id", "refund_amount", "refund_id", "is_advance_invoice", "reason")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Uid = temp.Uid
    i.CreditNoteNumber = temp.CreditNoteNumber
    i.CreditNoteUid = temp.CreditNoteUid
    i.OriginalAmount = temp.OriginalAmount
    i.AppliedAmount = temp.AppliedAmount
    if temp.TransactionTime != nil {
        TransactionTimeVal, err := time.Parse(time.RFC3339, *temp.TransactionTime)
        if err != nil {
            log.Fatalf("Cannot Parse transaction_time as % s format.", time.RFC3339)
        }
        i.TransactionTime = &TransactionTimeVal
    }
    i.Memo = temp.Memo
    i.Role = temp.Role
    i.ConsolidatedInvoice = temp.ConsolidatedInvoice
    i.AppliedCreditNotes = temp.AppliedCreditNotes
    i.DebitNoteNumber = temp.DebitNoteNumber
    i.DebitNoteUid = temp.DebitNoteUid
    i.PaymentMethod = temp.PaymentMethod
    i.TransactionId = temp.TransactionId
    i.ParentInvoiceNumber = temp.ParentInvoiceNumber
    i.RemainingPrepaymentAmount = temp.RemainingPrepaymentAmount
    i.Prepayment = temp.Prepayment
    i.External = temp.External
    i.FromCollectionMethod = temp.FromCollectionMethod
    i.ToCollectionMethod = temp.ToCollectionMethod
    i.ConsolidationLevel = temp.ConsolidationLevel
    i.FromStatus = temp.FromStatus
    i.ToStatus = temp.ToStatus
    i.DueAmount = temp.DueAmount
    i.TotalAmount = temp.TotalAmount
    i.ApplyCredit = temp.ApplyCredit
    i.CreditNoteAttributes = temp.CreditNoteAttributes
    i.PaymentId = temp.PaymentId
    i.RefundAmount = temp.RefundAmount
    i.RefundId = temp.RefundId
    i.IsAdvanceInvoice = temp.IsAdvanceInvoice
    i.Reason = temp.Reason
    return nil
}

// invoiceEventData is a temporary struct used for validating the fields of InvoiceEventData.
type invoiceEventData  struct {
    Uid                       *string                        `json:"uid,omitempty"`
    CreditNoteNumber          *string                        `json:"credit_note_number,omitempty"`
    CreditNoteUid             *string                        `json:"credit_note_uid,omitempty"`
    OriginalAmount            *string                        `json:"original_amount,omitempty"`
    AppliedAmount             *string                        `json:"applied_amount,omitempty"`
    TransactionTime           *string                        `json:"transaction_time,omitempty"`
    Memo                      *string                        `json:"memo,omitempty"`
    Role                      *string                        `json:"role,omitempty"`
    ConsolidatedInvoice       *bool                          `json:"consolidated_invoice,omitempty"`
    AppliedCreditNotes        []AppliedCreditNoteData        `json:"applied_credit_notes,omitempty"`
    DebitNoteNumber           *string                        `json:"debit_note_number,omitempty"`
    DebitNoteUid              *string                        `json:"debit_note_uid,omitempty"`
    PaymentMethod             *InvoiceEventDataPaymentMethod `json:"payment_method,omitempty"`
    TransactionId             *int                           `json:"transaction_id,omitempty"`
    ParentInvoiceNumber       Optional[int]                  `json:"parent_invoice_number"`
    RemainingPrepaymentAmount Optional[string]               `json:"remaining_prepayment_amount"`
    Prepayment                *bool                          `json:"prepayment,omitempty"`
    External                  *bool                          `json:"external,omitempty"`
    FromCollectionMethod      *string                        `json:"from_collection_method,omitempty"`
    ToCollectionMethod        *string                        `json:"to_collection_method,omitempty"`
    ConsolidationLevel        *InvoiceConsolidationLevel     `json:"consolidation_level,omitempty"`
    FromStatus                *InvoiceStatus                 `json:"from_status,omitempty"`
    ToStatus                  *InvoiceStatus                 `json:"to_status,omitempty"`
    DueAmount                 *string                        `json:"due_amount,omitempty"`
    TotalAmount               *string                        `json:"total_amount,omitempty"`
    ApplyCredit               *bool                          `json:"apply_credit,omitempty"`
    CreditNoteAttributes      *CreditNote1                   `json:"credit_note_attributes,omitempty"`
    PaymentId                 *int                           `json:"payment_id,omitempty"`
    RefundAmount              *string                        `json:"refund_amount,omitempty"`
    RefundId                  *int                           `json:"refund_id,omitempty"`
    IsAdvanceInvoice          *bool                          `json:"is_advance_invoice,omitempty"`
    Reason                    *string                        `json:"reason,omitempty"`
}
