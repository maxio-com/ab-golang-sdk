package models

import (
    "encoding/json"
)

// Refund represents a Refund struct.
type Refund struct {
    // The amount to be refunded in decimal format as a string. Example: "10.50". Must not exceed the remaining refundable balance of the payment.
    Amount               *string            `json:"amount,omitempty"`
    // A description that will be attached to the refund
    Memo                 *string            `json:"memo,omitempty"`
    // The ID of the payment to be refunded
    PaymentId            *int               `json:"payment_id,omitempty"`
    // Flag that marks refund as external (no money is returned to the customer). Defaults to `false`.
    External             *bool              `json:"external,omitempty"`
    // If set to true, creates credit and applies it to an invoice. Defaults to `false`.
    ApplyCredit          *bool              `json:"apply_credit,omitempty"`
    // If `apply_credit` set to false and refunding full amount, if `void_invoice` set to true, invoice will be voided after refund. Defaults to `false`.
    VoidInvoice          *bool              `json:"void_invoice,omitempty"`
    // An array of segment uids to refund or the string 'all' to indicate that all segments should be refunded
    SegmentUids          *RefundSegmentUids `json:"segment_uids,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Refund.
// It customizes the JSON marshaling process for Refund objects.
func (r Refund) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the Refund object to a map representation for JSON marshaling.
func (r Refund) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Amount != nil {
        structMap["amount"] = r.Amount
    }
    if r.Memo != nil {
        structMap["memo"] = r.Memo
    }
    if r.PaymentId != nil {
        structMap["payment_id"] = r.PaymentId
    }
    if r.External != nil {
        structMap["external"] = r.External
    }
    if r.ApplyCredit != nil {
        structMap["apply_credit"] = r.ApplyCredit
    }
    if r.VoidInvoice != nil {
        structMap["void_invoice"] = r.VoidInvoice
    }
    if r.SegmentUids != nil {
        structMap["segment_uids"] = r.SegmentUids.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Refund.
// It customizes the JSON unmarshaling process for Refund objects.
func (r *Refund) UnmarshalJSON(input []byte) error {
    var temp refund
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "amount", "memo", "payment_id", "external", "apply_credit", "void_invoice", "segment_uids")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Amount = temp.Amount
    r.Memo = temp.Memo
    r.PaymentId = temp.PaymentId
    r.External = temp.External
    r.ApplyCredit = temp.ApplyCredit
    r.VoidInvoice = temp.VoidInvoice
    r.SegmentUids = temp.SegmentUids
    return nil
}

// refund is a temporary struct used for validating the fields of Refund.
type refund  struct {
    Amount      *string            `json:"amount,omitempty"`
    Memo        *string            `json:"memo,omitempty"`
    PaymentId   *int               `json:"payment_id,omitempty"`
    External    *bool              `json:"external,omitempty"`
    ApplyCredit *bool              `json:"apply_credit,omitempty"`
    VoidInvoice *bool              `json:"void_invoice,omitempty"`
    SegmentUids *RefundSegmentUids `json:"segment_uids,omitempty"`
}
