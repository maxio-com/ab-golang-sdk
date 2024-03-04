package models

import (
    "encoding/json"
)

// InvoiceRefund represents a InvoiceRefund struct.
type InvoiceRefund struct {
    TransactionId        *int             `json:"transaction_id,omitempty"`
    PaymentId            *int             `json:"payment_id,omitempty"`
    Memo                 *string          `json:"memo,omitempty"`
    OriginalAmount       *string          `json:"original_amount,omitempty"`
    AppliedAmount        *string          `json:"applied_amount,omitempty"`
    // The transaction ID for the refund as returned from the payment gateway
    GatewayTransactionId Optional[string] `json:"gateway_transaction_id"`
    GatewayUsed          *string          `json:"gateway_used,omitempty"`
    GatewayHandle        Optional[string] `json:"gateway_handle"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceRefund.
// It customizes the JSON marshaling process for InvoiceRefund objects.
func (i *InvoiceRefund) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceRefund object to a map representation for JSON marshaling.
func (i *InvoiceRefund) toMap() map[string]any {
    structMap := make(map[string]any)
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
        structMap["gateway_transaction_id"] = i.GatewayTransactionId.Value()
    }
    if i.GatewayUsed != nil {
        structMap["gateway_used"] = i.GatewayUsed
    }
    if i.GatewayHandle.IsValueSet() {
        structMap["gateway_handle"] = i.GatewayHandle.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceRefund.
// It customizes the JSON unmarshaling process for InvoiceRefund objects.
func (i *InvoiceRefund) UnmarshalJSON(input []byte) error {
    temp := &struct {
        TransactionId        *int             `json:"transaction_id,omitempty"`
        PaymentId            *int             `json:"payment_id,omitempty"`
        Memo                 *string          `json:"memo,omitempty"`
        OriginalAmount       *string          `json:"original_amount,omitempty"`
        AppliedAmount        *string          `json:"applied_amount,omitempty"`
        GatewayTransactionId Optional[string] `json:"gateway_transaction_id"`
        GatewayUsed          *string          `json:"gateway_used,omitempty"`
        GatewayHandle        Optional[string] `json:"gateway_handle"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    i.TransactionId = temp.TransactionId
    i.PaymentId = temp.PaymentId
    i.Memo = temp.Memo
    i.OriginalAmount = temp.OriginalAmount
    i.AppliedAmount = temp.AppliedAmount
    i.GatewayTransactionId = temp.GatewayTransactionId
    i.GatewayUsed = temp.GatewayUsed
    i.GatewayHandle = temp.GatewayHandle
    return nil
}
