package models

import (
	"encoding/json"
	"log"
	"time"
)

// InvoicePayment represents a InvoicePayment struct.
type InvoicePayment struct {
	TransactionTime *time.Time            `json:"transaction_time,omitempty"`
	Memo            *string               `json:"memo,omitempty"`
	OriginalAmount  *string               `json:"original_amount,omitempty"`
	AppliedAmount   *string               `json:"applied_amount,omitempty"`
	PaymentMethod   *InvoicePaymentMethod `json:"payment_method,omitempty"`
	TransactionId   *int                  `json:"transaction_id,omitempty"`
	Prepayment      *bool                 `json:"prepayment,omitempty"`
	GatewayHandle   Optional[string]      `json:"gateway_handle"`
	GatewayUsed     *string               `json:"gateway_used,omitempty"`
	// The transaction ID for the payment as returned from the payment gateway
	GatewayTransactionId Optional[string] `json:"gateway_transaction_id"`
}

// MarshalJSON implements the json.Marshaler interface for InvoicePayment.
// It customizes the JSON marshaling process for InvoicePayment objects.
func (i *InvoicePayment) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InvoicePayment object to a map representation for JSON marshaling.
func (i *InvoicePayment) toMap() map[string]any {
	structMap := make(map[string]any)
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
		structMap["payment_method"] = i.PaymentMethod
	}
	if i.TransactionId != nil {
		structMap["transaction_id"] = i.TransactionId
	}
	if i.Prepayment != nil {
		structMap["prepayment"] = i.Prepayment
	}
	if i.GatewayHandle.IsValueSet() {
		structMap["gateway_handle"] = i.GatewayHandle.Value()
	}
	if i.GatewayUsed != nil {
		structMap["gateway_used"] = i.GatewayUsed
	}
	if i.GatewayTransactionId.IsValueSet() {
		structMap["gateway_transaction_id"] = i.GatewayTransactionId.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoicePayment.
// It customizes the JSON unmarshaling process for InvoicePayment objects.
func (i *InvoicePayment) UnmarshalJSON(input []byte) error {
	temp := &struct {
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
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

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
	return nil
}
