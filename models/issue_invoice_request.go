package models

import (
    "encoding/json"
)

// IssueInvoiceRequest represents a IssueInvoiceRequest struct.
type IssueInvoiceRequest struct {
    // Action taken when payment for an invoice fails:
    // - `leave_open_invoice` - prepayments and credits applied to invoice; invoice status set to "open"; email sent to the customer for the issued invoice (if setting applies); payment failure recorded in the invoice history. This is the default option.
    // - `rollback_to_pending` - prepayments and credits not applied; invoice remains in "pending" status; no email sent to the customer; payment failure recorded in the invoice history.
    // - `initiate_dunning` - prepayments and credits applied to the invoice; invoice status set to "open"; email sent to the customer for the issued invoice (if setting applies); payment failure recorded in the invoice history; subscription will  most likely go into "past_due" or "canceled" state (depending upon net terms and dunning settings).
    OnFailedPayment      *FailedPaymentAction `json:"on_failed_payment,omitempty"`
    AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for IssueInvoiceRequest.
// It customizes the JSON marshaling process for IssueInvoiceRequest objects.
func (i IssueInvoiceRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the IssueInvoiceRequest object to a map representation for JSON marshaling.
func (i IssueInvoiceRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.OnFailedPayment != nil {
        structMap["on_failed_payment"] = i.OnFailedPayment
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssueInvoiceRequest.
// It customizes the JSON unmarshaling process for IssueInvoiceRequest objects.
func (i *IssueInvoiceRequest) UnmarshalJSON(input []byte) error {
    var temp issueInvoiceRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "on_failed_payment")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.OnFailedPayment = temp.OnFailedPayment
    return nil
}

// issueInvoiceRequest is a temporary struct used for validating the fields of IssueInvoiceRequest.
type issueInvoiceRequest  struct {
    OnFailedPayment *FailedPaymentAction `json:"on_failed_payment,omitempty"`
}
