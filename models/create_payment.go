/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreatePayment represents a CreatePayment struct.
type CreatePayment struct {
    Amount               string                   `json:"amount"`
    Memo                 string                   `json:"memo"`
    PaymentDetails       string                   `json:"payment_details"`
    // The type of payment method used. Defaults to other.
    PaymentMethod        InvoicePaymentMethodType `json:"payment_method"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePayment.
// It customizes the JSON marshaling process for CreatePayment objects.
func (c CreatePayment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "amount", "memo", "payment_details", "payment_method"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePayment object to a map representation for JSON marshaling.
func (c CreatePayment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["amount"] = c.Amount
    structMap["memo"] = c.Memo
    structMap["payment_details"] = c.PaymentDetails
    structMap["payment_method"] = c.PaymentMethod
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePayment.
// It customizes the JSON unmarshaling process for CreatePayment objects.
func (c *CreatePayment) UnmarshalJSON(input []byte) error {
    var temp tempCreatePayment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "amount", "memo", "payment_details", "payment_method")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Amount = *temp.Amount
    c.Memo = *temp.Memo
    c.PaymentDetails = *temp.PaymentDetails
    c.PaymentMethod = *temp.PaymentMethod
    return nil
}

// tempCreatePayment is a temporary struct used for validating the fields of CreatePayment.
type tempCreatePayment  struct {
    Amount         *string                   `json:"amount"`
    Memo           *string                   `json:"memo"`
    PaymentDetails *string                   `json:"payment_details"`
    PaymentMethod  *InvoicePaymentMethodType `json:"payment_method"`
}

func (c *tempCreatePayment) validate() error {
    var errs []string
    if c.Amount == nil {
        errs = append(errs, "required field `amount` is missing for type `Create Payment`")
    }
    if c.Memo == nil {
        errs = append(errs, "required field `memo` is missing for type `Create Payment`")
    }
    if c.PaymentDetails == nil {
        errs = append(errs, "required field `payment_details` is missing for type `Create Payment`")
    }
    if c.PaymentMethod == nil {
        errs = append(errs, "required field `payment_method` is missing for type `Create Payment`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
