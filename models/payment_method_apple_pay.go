// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// PaymentMethodApplePay represents a PaymentMethodApplePay struct.
type PaymentMethodApplePay struct {
    Type                 InvoiceEventPaymentMethod `json:"type"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for PaymentMethodApplePay,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PaymentMethodApplePay) String() string {
    return fmt.Sprintf(
    	"PaymentMethodApplePay[Type=%v, AdditionalProperties=%v]",
    	p.Type, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PaymentMethodApplePay.
// It customizes the JSON marshaling process for PaymentMethodApplePay objects.
func (p PaymentMethodApplePay) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentMethodApplePay object to a map representation for JSON marshaling.
func (p PaymentMethodApplePay) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["type"] = p.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentMethodApplePay.
// It customizes the JSON unmarshaling process for PaymentMethodApplePay objects.
func (p *PaymentMethodApplePay) UnmarshalJSON(input []byte) error {
    var temp tempPaymentMethodApplePay
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "type")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Type = *temp.Type
    return nil
}

// tempPaymentMethodApplePay is a temporary struct used for validating the fields of PaymentMethodApplePay.
type tempPaymentMethodApplePay  struct {
    Type *InvoiceEventPaymentMethod `json:"type"`
}

func (p *tempPaymentMethodApplePay) validate() error {
    var errs []string
    if p.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Payment Method Apple Pay`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
