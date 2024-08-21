/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// RefundInvoice represents a RefundInvoice struct.
// Refund an invoice or a segment of a consolidated invoice.
type RefundInvoice struct {
    // The amount to be refunded in decimal format as a string. Example: "10.50". Must not exceed the remaining refundable balance of the payment.
    Amount               string         `json:"amount"`
    // A description that will be attached to the refund
    Memo                 string         `json:"memo"`
    // The ID of the payment to be refunded
    PaymentId            int            `json:"payment_id"`
    // Flag that marks refund as external (no money is returned to the customer). Defaults to `false`.
    External             *bool          `json:"external,omitempty"`
    // If set to true, creates credit and applies it to an invoice. Defaults to `false`.
    ApplyCredit          *bool          `json:"apply_credit,omitempty"`
    // If `apply_credit` set to false and refunding full amount, if `void_invoice` set to true, invoice will be voided after refund. Defaults to `false`.
    VoidInvoice          *bool          `json:"void_invoice,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RefundInvoice.
// It customizes the JSON marshaling process for RefundInvoice objects.
func (r RefundInvoice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RefundInvoice object to a map representation for JSON marshaling.
func (r RefundInvoice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["amount"] = r.Amount
    structMap["memo"] = r.Memo
    structMap["payment_id"] = r.PaymentId
    if r.External != nil {
        structMap["external"] = r.External
    }
    if r.ApplyCredit != nil {
        structMap["apply_credit"] = r.ApplyCredit
    }
    if r.VoidInvoice != nil {
        structMap["void_invoice"] = r.VoidInvoice
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundInvoice.
// It customizes the JSON unmarshaling process for RefundInvoice objects.
func (r *RefundInvoice) UnmarshalJSON(input []byte) error {
    var temp tempRefundInvoice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "amount", "memo", "payment_id", "external", "apply_credit", "void_invoice")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Amount = *temp.Amount
    r.Memo = *temp.Memo
    r.PaymentId = *temp.PaymentId
    r.External = temp.External
    r.ApplyCredit = temp.ApplyCredit
    r.VoidInvoice = temp.VoidInvoice
    return nil
}

// tempRefundInvoice is a temporary struct used for validating the fields of RefundInvoice.
type tempRefundInvoice  struct {
    Amount      *string `json:"amount"`
    Memo        *string `json:"memo"`
    PaymentId   *int    `json:"payment_id"`
    External    *bool   `json:"external,omitempty"`
    ApplyCredit *bool   `json:"apply_credit,omitempty"`
    VoidInvoice *bool   `json:"void_invoice,omitempty"`
}

func (r *tempRefundInvoice) validate() error {
    var errs []string
    if r.Amount == nil {
        errs = append(errs, "required field `amount` is missing for type `Refund Invoice`")
    }
    if r.Memo == nil {
        errs = append(errs, "required field `memo` is missing for type `Refund Invoice`")
    }
    if r.PaymentId == nil {
        errs = append(errs, "required field `payment_id` is missing for type `Refund Invoice`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
