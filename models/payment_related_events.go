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

// PaymentRelatedEvents represents a PaymentRelatedEvents struct.
type PaymentRelatedEvents struct {
    ProductId            int                    `json:"product_id"`
    AccountTransactionId int                    `json:"account_transaction_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PaymentRelatedEvents,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PaymentRelatedEvents) String() string {
    return fmt.Sprintf(
    	"PaymentRelatedEvents[ProductId=%v, AccountTransactionId=%v, AdditionalProperties=%v]",
    	p.ProductId, p.AccountTransactionId, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PaymentRelatedEvents.
// It customizes the JSON marshaling process for PaymentRelatedEvents objects.
func (p PaymentRelatedEvents) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "product_id", "account_transaction_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentRelatedEvents object to a map representation for JSON marshaling.
func (p PaymentRelatedEvents) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["product_id"] = p.ProductId
    structMap["account_transaction_id"] = p.AccountTransactionId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentRelatedEvents.
// It customizes the JSON unmarshaling process for PaymentRelatedEvents objects.
func (p *PaymentRelatedEvents) UnmarshalJSON(input []byte) error {
    var temp tempPaymentRelatedEvents
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "product_id", "account_transaction_id")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.ProductId = *temp.ProductId
    p.AccountTransactionId = *temp.AccountTransactionId
    return nil
}

// tempPaymentRelatedEvents is a temporary struct used for validating the fields of PaymentRelatedEvents.
type tempPaymentRelatedEvents  struct {
    ProductId            *int `json:"product_id"`
    AccountTransactionId *int `json:"account_transaction_id"`
}

func (p *tempPaymentRelatedEvents) validate() error {
    var errs []string
    if p.ProductId == nil {
        errs = append(errs, "required field `product_id` is missing for type `Payment Related Events`")
    }
    if p.AccountTransactionId == nil {
        errs = append(errs, "required field `account_transaction_id` is missing for type `Payment Related Events`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
