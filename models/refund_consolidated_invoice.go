/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// RefundConsolidatedInvoice represents a RefundConsolidatedInvoice struct.
// Refund consolidated invoice
type RefundConsolidatedInvoice struct {
    // A description for the refund
    Memo                 string                               `json:"memo"`
    // The ID of the payment to be refunded
    PaymentId            int                                  `json:"payment_id"`
    // An array of segment uids to refund or the string 'all' to indicate that all segments should be refunded
    SegmentUids          RefundConsolidatedInvoiceSegmentUids `json:"segment_uids"`
    // Flag that marks refund as external (no money is returned to the customer). Defaults to `false`.
    External             *bool                                `json:"external,omitempty"`
    // If set to true, creates credit and applies it to an invoice. Defaults to `false`.
    ApplyCredit          *bool                                `json:"apply_credit,omitempty"`
    // The amount of payment to be refunded in decimal format. Example: "10.50". This will default to the full amount of the payment if not provided.
    Amount               *string                              `json:"amount,omitempty"`
    AdditionalProperties map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for RefundConsolidatedInvoice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RefundConsolidatedInvoice) String() string {
    return fmt.Sprintf(
    	"RefundConsolidatedInvoice[Memo=%v, PaymentId=%v, SegmentUids=%v, External=%v, ApplyCredit=%v, Amount=%v, AdditionalProperties=%v]",
    	r.Memo, r.PaymentId, r.SegmentUids, r.External, r.ApplyCredit, r.Amount, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RefundConsolidatedInvoice.
// It customizes the JSON marshaling process for RefundConsolidatedInvoice objects.
func (r RefundConsolidatedInvoice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "memo", "payment_id", "segment_uids", "external", "apply_credit", "amount"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RefundConsolidatedInvoice object to a map representation for JSON marshaling.
func (r RefundConsolidatedInvoice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["memo"] = r.Memo
    structMap["payment_id"] = r.PaymentId
    structMap["segment_uids"] = r.SegmentUids.toMap()
    if r.External != nil {
        structMap["external"] = r.External
    }
    if r.ApplyCredit != nil {
        structMap["apply_credit"] = r.ApplyCredit
    }
    if r.Amount != nil {
        structMap["amount"] = r.Amount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundConsolidatedInvoice.
// It customizes the JSON unmarshaling process for RefundConsolidatedInvoice objects.
func (r *RefundConsolidatedInvoice) UnmarshalJSON(input []byte) error {
    var temp tempRefundConsolidatedInvoice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "memo", "payment_id", "segment_uids", "external", "apply_credit", "amount")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Memo = *temp.Memo
    r.PaymentId = *temp.PaymentId
    r.SegmentUids = *temp.SegmentUids
    r.External = temp.External
    r.ApplyCredit = temp.ApplyCredit
    r.Amount = temp.Amount
    return nil
}

// tempRefundConsolidatedInvoice is a temporary struct used for validating the fields of RefundConsolidatedInvoice.
type tempRefundConsolidatedInvoice  struct {
    Memo        *string                               `json:"memo"`
    PaymentId   *int                                  `json:"payment_id"`
    SegmentUids *RefundConsolidatedInvoiceSegmentUids `json:"segment_uids"`
    External    *bool                                 `json:"external,omitempty"`
    ApplyCredit *bool                                 `json:"apply_credit,omitempty"`
    Amount      *string                               `json:"amount,omitempty"`
}

func (r *tempRefundConsolidatedInvoice) validate() error {
    var errs []string
    if r.Memo == nil {
        errs = append(errs, "required field `memo` is missing for type `Refund Consolidated Invoice`")
    }
    if r.PaymentId == nil {
        errs = append(errs, "required field `payment_id` is missing for type `Refund Consolidated Invoice`")
    }
    if r.SegmentUids == nil {
        errs = append(errs, "required field `segment_uids` is missing for type `Refund Consolidated Invoice`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
