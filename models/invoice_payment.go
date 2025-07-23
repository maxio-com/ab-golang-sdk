// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// InvoicePayment represents a InvoicePayment struct.
type InvoicePayment struct {
    TransactionTime      *time.Time             `json:"transaction_time,omitempty"`
    Memo                 *string                `json:"memo,omitempty"`
    OriginalAmount       *string                `json:"original_amount,omitempty"`
    AppliedAmount        *string                `json:"applied_amount,omitempty"`
    PaymentMethod        *InvoicePaymentMethod  `json:"payment_method,omitempty"`
    TransactionId        *int                   `json:"transaction_id,omitempty"`
    Prepayment           *bool                  `json:"prepayment,omitempty"`
    GatewayHandle        Optional[string]       `json:"gateway_handle"`
    GatewayUsed          *string                `json:"gateway_used,omitempty"`
    // The transaction ID for the payment as returned from the payment gateway
    GatewayTransactionId Optional[string]       `json:"gateway_transaction_id"`
    // Date reflecting when the payment was received from a customer. Must be in the past. Applicable only to
    // `external` payments.
    ReceivedOn           Optional[time.Time]    `json:"received_on"`
    Uid                  *string                `json:"uid,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InvoicePayment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoicePayment) String() string {
    return fmt.Sprintf(
    	"InvoicePayment[TransactionTime=%v, Memo=%v, OriginalAmount=%v, AppliedAmount=%v, PaymentMethod=%v, TransactionId=%v, Prepayment=%v, GatewayHandle=%v, GatewayUsed=%v, GatewayTransactionId=%v, ReceivedOn=%v, Uid=%v, AdditionalProperties=%v]",
    	i.TransactionTime, i.Memo, i.OriginalAmount, i.AppliedAmount, i.PaymentMethod, i.TransactionId, i.Prepayment, i.GatewayHandle, i.GatewayUsed, i.GatewayTransactionId, i.ReceivedOn, i.Uid, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoicePayment.
// It customizes the JSON marshaling process for InvoicePayment objects.
func (i InvoicePayment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "transaction_time", "memo", "original_amount", "applied_amount", "payment_method", "transaction_id", "prepayment", "gateway_handle", "gateway_used", "gateway_transaction_id", "received_on", "uid"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoicePayment object to a map representation for JSON marshaling.
func (i InvoicePayment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.TransactionTime != nil {
        structMap["transaction_time"] = i.TransactionTime.Format(time.RFC3339)
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
    if i.PaymentMethod != nil {
        structMap["payment_method"] = i.PaymentMethod.toMap()
    }
    if i.TransactionId != nil {
        structMap["transaction_id"] = i.TransactionId
    }
    if i.Prepayment != nil {
        structMap["prepayment"] = i.Prepayment
    }
    if i.GatewayHandle.IsValueSet() {
        if i.GatewayHandle.Value() != nil {
            structMap["gateway_handle"] = i.GatewayHandle.Value()
        } else {
            structMap["gateway_handle"] = nil
        }
    }
    if i.GatewayUsed != nil {
        structMap["gateway_used"] = i.GatewayUsed
    }
    if i.GatewayTransactionId.IsValueSet() {
        if i.GatewayTransactionId.Value() != nil {
            structMap["gateway_transaction_id"] = i.GatewayTransactionId.Value()
        } else {
            structMap["gateway_transaction_id"] = nil
        }
    }
    if i.ReceivedOn.IsValueSet() {
        var ReceivedOnVal *string = nil
        if i.ReceivedOn.Value() != nil {
            val := i.ReceivedOn.Value().Format(DEFAULT_DATE)
            ReceivedOnVal = &val
        }
        if i.ReceivedOn.Value() != nil {
            structMap["received_on"] = ReceivedOnVal
        } else {
            structMap["received_on"] = nil
        }
    }
    if i.Uid != nil {
        structMap["uid"] = i.Uid
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoicePayment.
// It customizes the JSON unmarshaling process for InvoicePayment objects.
func (i *InvoicePayment) UnmarshalJSON(input []byte) error {
    var temp tempInvoicePayment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "transaction_time", "memo", "original_amount", "applied_amount", "payment_method", "transaction_id", "prepayment", "gateway_handle", "gateway_used", "gateway_transaction_id", "received_on", "uid")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    if temp.TransactionTime != nil {
        TransactionTimeVal, err := time.Parse(time.RFC3339, *temp.TransactionTime)
        if err != nil {
            log.Fatalf("Cannot Parse transaction_time as % s format.", time.RFC3339)
        }
        i.TransactionTime = &TransactionTimeVal
    }
    i.Memo = temp.Memo
    i.OriginalAmount = temp.OriginalAmount
    i.AppliedAmount = temp.AppliedAmount
    i.PaymentMethod = temp.PaymentMethod
    i.TransactionId = temp.TransactionId
    i.Prepayment = temp.Prepayment
    i.GatewayHandle = temp.GatewayHandle
    i.GatewayUsed = temp.GatewayUsed
    i.GatewayTransactionId = temp.GatewayTransactionId
    i.ReceivedOn.ShouldSetValue(temp.ReceivedOn.IsValueSet())
    if temp.ReceivedOn.Value() != nil {
        ReceivedOnVal, err := time.Parse(DEFAULT_DATE, (*temp.ReceivedOn.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse received_on as % s format.", DEFAULT_DATE)
        }
        i.ReceivedOn.SetValue(&ReceivedOnVal)
    }
    i.Uid = temp.Uid
    return nil
}

// tempInvoicePayment is a temporary struct used for validating the fields of InvoicePayment.
type tempInvoicePayment  struct {
    TransactionTime      *string               `json:"transaction_time,omitempty"`
    Memo                 *string               `json:"memo,omitempty"`
    OriginalAmount       *string               `json:"original_amount,omitempty"`
    AppliedAmount        *string               `json:"applied_amount,omitempty"`
    PaymentMethod        *InvoicePaymentMethod `json:"payment_method,omitempty"`
    TransactionId        *int                  `json:"transaction_id,omitempty"`
    Prepayment           *bool                 `json:"prepayment,omitempty"`
    GatewayHandle        Optional[string]      `json:"gateway_handle"`
    GatewayUsed          *string               `json:"gateway_used,omitempty"`
    GatewayTransactionId Optional[string]      `json:"gateway_transaction_id"`
    ReceivedOn           Optional[string]      `json:"received_on"`
    Uid                  *string               `json:"uid,omitempty"`
}
