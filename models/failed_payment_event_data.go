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

// FailedPaymentEventData represents a FailedPaymentEventData struct.
// Example schema for an `failed_payment` event
type FailedPaymentEventData struct {
    // The monetary value of the payment, expressed in cents.
    AmountInCents        int                      `json:"amount_in_cents"`
    // The monetary value of the payment, expressed in dollars.
    AppliedAmount        int                      `json:"applied_amount"`
    // The memo passed when the payment was created.
    Memo                 Optional[string]         `json:"memo"`
    PaymentMethod        InvoicePaymentMethodType `json:"payment_method"`
    // The transaction ID of the failed payment.
    TransactionId        int                      `json:"transaction_id"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for FailedPaymentEventData,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f FailedPaymentEventData) String() string {
    return fmt.Sprintf(
    	"FailedPaymentEventData[AmountInCents=%v, AppliedAmount=%v, Memo=%v, PaymentMethod=%v, TransactionId=%v, AdditionalProperties=%v]",
    	f.AmountInCents, f.AppliedAmount, f.Memo, f.PaymentMethod, f.TransactionId, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for FailedPaymentEventData.
// It customizes the JSON marshaling process for FailedPaymentEventData objects.
func (f FailedPaymentEventData) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(f.AdditionalProperties,
        "amount_in_cents", "applied_amount", "memo", "payment_method", "transaction_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(f.toMap())
}

// toMap converts the FailedPaymentEventData object to a map representation for JSON marshaling.
func (f FailedPaymentEventData) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, f.AdditionalProperties)
    structMap["amount_in_cents"] = f.AmountInCents
    structMap["applied_amount"] = f.AppliedAmount
    if f.Memo.IsValueSet() {
        if f.Memo.Value() != nil {
            structMap["memo"] = f.Memo.Value()
        } else {
            structMap["memo"] = nil
        }
    }
    structMap["payment_method"] = f.PaymentMethod
    structMap["transaction_id"] = f.TransactionId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FailedPaymentEventData.
// It customizes the JSON unmarshaling process for FailedPaymentEventData objects.
func (f *FailedPaymentEventData) UnmarshalJSON(input []byte) error {
    var temp tempFailedPaymentEventData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "amount_in_cents", "applied_amount", "memo", "payment_method", "transaction_id")
    if err != nil {
    	return err
    }
    f.AdditionalProperties = additionalProperties
    
    f.AmountInCents = *temp.AmountInCents
    f.AppliedAmount = *temp.AppliedAmount
    f.Memo = temp.Memo
    f.PaymentMethod = *temp.PaymentMethod
    f.TransactionId = *temp.TransactionId
    return nil
}

// tempFailedPaymentEventData is a temporary struct used for validating the fields of FailedPaymentEventData.
type tempFailedPaymentEventData  struct {
    AmountInCents *int                      `json:"amount_in_cents"`
    AppliedAmount *int                      `json:"applied_amount"`
    Memo          Optional[string]          `json:"memo"`
    PaymentMethod *InvoicePaymentMethodType `json:"payment_method"`
    TransactionId *int                      `json:"transaction_id"`
}

func (f *tempFailedPaymentEventData) validate() error {
    var errs []string
    if f.AmountInCents == nil {
        errs = append(errs, "required field `amount_in_cents` is missing for type `Failed Payment Event Data`")
    }
    if f.AppliedAmount == nil {
        errs = append(errs, "required field `applied_amount` is missing for type `Failed Payment Event Data`")
    }
    if f.PaymentMethod == nil {
        errs = append(errs, "required field `payment_method` is missing for type `Failed Payment Event Data`")
    }
    if f.TransactionId == nil {
        errs = append(errs, "required field `transaction_id` is missing for type `Failed Payment Event Data`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
