/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// IssueInvoiceRequest represents a IssueInvoiceRequest struct.
type IssueInvoiceRequest struct {
    // Action taken when payment for an invoice fails:
    // - `leave_open_invoice` - prepayments and credits applied to invoice; invoice status set to "open"; email sent to the customer for the issued invoice (if setting applies); payment failure recorded in the invoice history. This is the default option.
    // - `rollback_to_pending` - prepayments and credits not applied; invoice remains in "pending" status; no email sent to the customer; payment failure recorded in the invoice history.
    // - `initiate_dunning` - prepayments and credits applied to the invoice; invoice status set to "open"; email sent to the customer for the issued invoice (if setting applies); payment failure recorded in the invoice history; subscription will  most likely go into "past_due" or "canceled" state (depending upon net terms and dunning settings).
    OnFailedPayment      *FailedPaymentAction   `json:"on_failed_payment,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for IssueInvoiceRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IssueInvoiceRequest) String() string {
    return fmt.Sprintf(
    	"IssueInvoiceRequest[OnFailedPayment=%v, AdditionalProperties=%v]",
    	i.OnFailedPayment, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for IssueInvoiceRequest.
// It customizes the JSON marshaling process for IssueInvoiceRequest objects.
func (i IssueInvoiceRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "on_failed_payment"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the IssueInvoiceRequest object to a map representation for JSON marshaling.
func (i IssueInvoiceRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.OnFailedPayment != nil {
        structMap["on_failed_payment"] = i.OnFailedPayment
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssueInvoiceRequest.
// It customizes the JSON unmarshaling process for IssueInvoiceRequest objects.
func (i *IssueInvoiceRequest) UnmarshalJSON(input []byte) error {
    var temp tempIssueInvoiceRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "on_failed_payment")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.OnFailedPayment = temp.OnFailedPayment
    return nil
}

// tempIssueInvoiceRequest is a temporary struct used for validating the fields of IssueInvoiceRequest.
type tempIssueInvoiceRequest  struct {
    OnFailedPayment *FailedPaymentAction `json:"on_failed_payment,omitempty"`
}
