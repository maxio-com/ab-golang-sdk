/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "log"
    "strings"
    "time"
)

// RemovePaymentEventData represents a RemovePaymentEventData struct.
// Example schema for an `remove_payment` event
type RemovePaymentEventData struct {
    // Transaction ID of the original payment that was removed
    TransactionId        int                 `json:"transaction_id"`
    // Memo of the original payment
    Memo                 string              `json:"memo"`
    // Full amount of the original payment
    OriginalAmount       *string             `json:"original_amount,omitempty"`
    // Applied amount of the original payment
    AppliedAmount        string              `json:"applied_amount"`
    // Transaction time of the original payment, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z"
    TransactionTime      time.Time           `json:"transaction_time"`
    // A nested data structure detailing the method of payment
    PaymentMethod        InvoiceEventPayment `json:"payment_method"`
    // The flag that shows whether the original payment was a prepayment or not
    Prepayment           bool                `json:"prepayment"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RemovePaymentEventData.
// It customizes the JSON marshaling process for RemovePaymentEventData objects.
func (r RemovePaymentEventData) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RemovePaymentEventData object to a map representation for JSON marshaling.
func (r RemovePaymentEventData) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["transaction_id"] = r.TransactionId
    structMap["memo"] = r.Memo
    if r.OriginalAmount != nil {
        structMap["original_amount"] = r.OriginalAmount
    }
    structMap["applied_amount"] = r.AppliedAmount
    structMap["transaction_time"] = r.TransactionTime.Format(time.RFC3339)
    structMap["payment_method"] = r.PaymentMethod.toMap()
    structMap["prepayment"] = r.Prepayment
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RemovePaymentEventData.
// It customizes the JSON unmarshaling process for RemovePaymentEventData objects.
func (r *RemovePaymentEventData) UnmarshalJSON(input []byte) error {
    var temp tempRemovePaymentEventData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "transaction_id", "memo", "original_amount", "applied_amount", "transaction_time", "payment_method", "prepayment")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.TransactionId = *temp.TransactionId
    r.Memo = *temp.Memo
    r.OriginalAmount = temp.OriginalAmount
    r.AppliedAmount = *temp.AppliedAmount
    TransactionTimeVal, err := time.Parse(time.RFC3339, *temp.TransactionTime)
    if err != nil {
        log.Fatalf("Cannot Parse transaction_time as % s format.", time.RFC3339)
    }
    r.TransactionTime = TransactionTimeVal
    r.PaymentMethod = *temp.PaymentMethod
    r.Prepayment = *temp.Prepayment
    return nil
}

// tempRemovePaymentEventData is a temporary struct used for validating the fields of RemovePaymentEventData.
type tempRemovePaymentEventData  struct {
    TransactionId   *int                 `json:"transaction_id"`
    Memo            *string              `json:"memo"`
    OriginalAmount  *string              `json:"original_amount,omitempty"`
    AppliedAmount   *string              `json:"applied_amount"`
    TransactionTime *string              `json:"transaction_time"`
    PaymentMethod   *InvoiceEventPayment `json:"payment_method"`
    Prepayment      *bool                `json:"prepayment"`
}

func (r *tempRemovePaymentEventData) validate() error {
    var errs []string
    if r.TransactionId == nil {
        errs = append(errs, "required field `transaction_id` is missing for type `Remove Payment Event Data`")
    }
    if r.Memo == nil {
        errs = append(errs, "required field `memo` is missing for type `Remove Payment Event Data`")
    }
    if r.AppliedAmount == nil {
        errs = append(errs, "required field `applied_amount` is missing for type `Remove Payment Event Data`")
    }
    if r.TransactionTime == nil {
        errs = append(errs, "required field `transaction_time` is missing for type `Remove Payment Event Data`")
    }
    if r.PaymentMethod == nil {
        errs = append(errs, "required field `payment_method` is missing for type `Remove Payment Event Data`")
    }
    if r.Prepayment == nil {
        errs = append(errs, "required field `prepayment` is missing for type `Remove Payment Event Data`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
