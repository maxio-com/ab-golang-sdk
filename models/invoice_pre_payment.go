/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// InvoicePrePayment represents a InvoicePrePayment struct.
type InvoicePrePayment struct {
    // The subscription id for the prepayment account
    SubscriptionId       *int                   `json:"subscription_id,omitempty"`
    // The amount in cents of the prepayment that was created as a result of this payment.
    AmountInCents        *int64                 `json:"amount_in_cents,omitempty"`
    // The total balance of the prepayment account for this subscription including any prior prepayments
    EndingBalanceInCents *int64                 `json:"ending_balance_in_cents,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InvoicePrePayment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoicePrePayment) String() string {
    return fmt.Sprintf(
    	"InvoicePrePayment[SubscriptionId=%v, AmountInCents=%v, EndingBalanceInCents=%v, AdditionalProperties=%v]",
    	i.SubscriptionId, i.AmountInCents, i.EndingBalanceInCents, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoicePrePayment.
// It customizes the JSON marshaling process for InvoicePrePayment objects.
func (i InvoicePrePayment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "subscription_id", "amount_in_cents", "ending_balance_in_cents"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoicePrePayment object to a map representation for JSON marshaling.
func (i InvoicePrePayment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.SubscriptionId != nil {
        structMap["subscription_id"] = i.SubscriptionId
    }
    if i.AmountInCents != nil {
        structMap["amount_in_cents"] = i.AmountInCents
    }
    if i.EndingBalanceInCents != nil {
        structMap["ending_balance_in_cents"] = i.EndingBalanceInCents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoicePrePayment.
// It customizes the JSON unmarshaling process for InvoicePrePayment objects.
func (i *InvoicePrePayment) UnmarshalJSON(input []byte) error {
    var temp tempInvoicePrePayment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscription_id", "amount_in_cents", "ending_balance_in_cents")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.SubscriptionId = temp.SubscriptionId
    i.AmountInCents = temp.AmountInCents
    i.EndingBalanceInCents = temp.EndingBalanceInCents
    return nil
}

// tempInvoicePrePayment is a temporary struct used for validating the fields of InvoicePrePayment.
type tempInvoicePrePayment  struct {
    SubscriptionId       *int   `json:"subscription_id,omitempty"`
    AmountInCents        *int64 `json:"amount_in_cents,omitempty"`
    EndingBalanceInCents *int64 `json:"ending_balance_in_cents,omitempty"`
}
