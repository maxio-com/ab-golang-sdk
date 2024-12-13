/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// InvoiceRefund represents a InvoiceRefund struct.
type InvoiceRefund struct {
    TransactionId        *int                   `json:"transaction_id,omitempty"`
    PaymentId            *int                   `json:"payment_id,omitempty"`
    Memo                 *string                `json:"memo,omitempty"`
    OriginalAmount       *string                `json:"original_amount,omitempty"`
    AppliedAmount        *string                `json:"applied_amount,omitempty"`
    // The transaction ID for the refund as returned from the payment gateway
    GatewayTransactionId Optional[string]       `json:"gateway_transaction_id"`
    GatewayUsed          *string                `json:"gateway_used,omitempty"`
    GatewayHandle        Optional[string]       `json:"gateway_handle"`
    AchLateReject        Optional[bool]         `json:"ach_late_reject"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceRefund.
// It customizes the JSON marshaling process for InvoiceRefund objects.
func (i InvoiceRefund) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "transaction_id", "payment_id", "memo", "original_amount", "applied_amount", "gateway_transaction_id", "gateway_used", "gateway_handle", "ach_late_reject"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceRefund object to a map representation for JSON marshaling.
func (i InvoiceRefund) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.TransactionId != nil {
        structMap["transaction_id"] = i.TransactionId
    }
    if i.PaymentId != nil {
        structMap["payment_id"] = i.PaymentId
    }
    if i.Memo != nil {
        structMap["memo"] = i.Memo
    }
    if i.OriginalAmount != nil {
        structMap["original_amount"] = i.OriginalAmount
    }
    if i.AppliedAmount != nil {
        structMap["applied_amount"] = i.AppliedAmount
    }
    if i.GatewayTransactionId.IsValueSet() {
        if i.GatewayTransactionId.Value() != nil {
            structMap["gateway_transaction_id"] = i.GatewayTransactionId.Value()
        } else {
            structMap["gateway_transaction_id"] = nil
        }
    }
    if i.GatewayUsed != nil {
        structMap["gateway_used"] = i.GatewayUsed
    }
    if i.GatewayHandle.IsValueSet() {
        if i.GatewayHandle.Value() != nil {
            structMap["gateway_handle"] = i.GatewayHandle.Value()
        } else {
            structMap["gateway_handle"] = nil
        }
    }
    if i.AchLateReject.IsValueSet() {
        if i.AchLateReject.Value() != nil {
            structMap["ach_late_reject"] = i.AchLateReject.Value()
        } else {
            structMap["ach_late_reject"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceRefund.
// It customizes the JSON unmarshaling process for InvoiceRefund objects.
func (i *InvoiceRefund) UnmarshalJSON(input []byte) error {
    var temp tempInvoiceRefund
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "transaction_id", "payment_id", "memo", "original_amount", "applied_amount", "gateway_transaction_id", "gateway_used", "gateway_handle", "ach_late_reject")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.TransactionId = temp.TransactionId
    i.PaymentId = temp.PaymentId
    i.Memo = temp.Memo
    i.OriginalAmount = temp.OriginalAmount
    i.AppliedAmount = temp.AppliedAmount
    i.GatewayTransactionId = temp.GatewayTransactionId
    i.GatewayUsed = temp.GatewayUsed
    i.GatewayHandle = temp.GatewayHandle
    i.AchLateReject = temp.AchLateReject
    return nil
}

// tempInvoiceRefund is a temporary struct used for validating the fields of InvoiceRefund.
type tempInvoiceRefund  struct {
    TransactionId        *int             `json:"transaction_id,omitempty"`
    PaymentId            *int             `json:"payment_id,omitempty"`
    Memo                 *string          `json:"memo,omitempty"`
    OriginalAmount       *string          `json:"original_amount,omitempty"`
    AppliedAmount        *string          `json:"applied_amount,omitempty"`
    GatewayTransactionId Optional[string] `json:"gateway_transaction_id"`
    GatewayUsed          *string          `json:"gateway_used,omitempty"`
    GatewayHandle        Optional[string] `json:"gateway_handle"`
    AchLateReject        Optional[bool]   `json:"ach_late_reject"`
}
