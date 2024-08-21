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

// PaymentMethodPaypal represents a PaymentMethodPaypal struct.
type PaymentMethodPaypal struct {
    Email                string                    `json:"email"`
    Type                 InvoiceEventPaymentMethod `json:"type"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentMethodPaypal.
// It customizes the JSON marshaling process for PaymentMethodPaypal objects.
func (p PaymentMethodPaypal) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentMethodPaypal object to a map representation for JSON marshaling.
func (p PaymentMethodPaypal) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["email"] = p.Email
    structMap["type"] = p.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentMethodPaypal.
// It customizes the JSON unmarshaling process for PaymentMethodPaypal objects.
func (p *PaymentMethodPaypal) UnmarshalJSON(input []byte) error {
    var temp tempPaymentMethodPaypal
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "email", "type")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Email = *temp.Email
    p.Type = *temp.Type
    return nil
}

// tempPaymentMethodPaypal is a temporary struct used for validating the fields of PaymentMethodPaypal.
type tempPaymentMethodPaypal  struct {
    Email *string                    `json:"email"`
    Type  *InvoiceEventPaymentMethod `json:"type"`
}

func (p *tempPaymentMethodPaypal) validate() error {
    var errs []string
    if p.Email == nil {
        errs = append(errs, "required field `email` is missing for type `Payment Method Paypal`")
    }
    if p.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Payment Method Paypal`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
