// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// PaymentMethodExternal represents a PaymentMethodExternal struct.
type PaymentMethodExternal struct {
    Details              *string                   `json:"details"`
    Kind                 string                    `json:"kind"`
    Memo                 *string                   `json:"memo"`
    Type                 InvoiceEventPaymentMethod `json:"type"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for PaymentMethodExternal,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PaymentMethodExternal) String() string {
    return fmt.Sprintf(
    	"PaymentMethodExternal[Details=%v, Kind=%v, Memo=%v, Type=%v, AdditionalProperties=%v]",
    	p.Details, p.Kind, p.Memo, p.Type, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PaymentMethodExternal.
// It customizes the JSON marshaling process for PaymentMethodExternal objects.
func (p PaymentMethodExternal) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "details", "kind", "memo", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentMethodExternal object to a map representation for JSON marshaling.
func (p PaymentMethodExternal) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Details != nil {
        structMap["details"] = p.Details
    } else {
        structMap["details"] = nil
    }
    structMap["kind"] = p.Kind
    if p.Memo != nil {
        structMap["memo"] = p.Memo
    } else {
        structMap["memo"] = nil
    }
    structMap["type"] = p.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentMethodExternal.
// It customizes the JSON unmarshaling process for PaymentMethodExternal objects.
func (p *PaymentMethodExternal) UnmarshalJSON(input []byte) error {
    var temp tempPaymentMethodExternal
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "details", "kind", "memo", "type")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Details = temp.Details
    p.Kind = *temp.Kind
    p.Memo = temp.Memo
    p.Type = *temp.Type
    return nil
}

// tempPaymentMethodExternal is a temporary struct used for validating the fields of PaymentMethodExternal.
type tempPaymentMethodExternal  struct {
    Details *string                    `json:"details"`
    Kind    *string                    `json:"kind"`
    Memo    *string                    `json:"memo"`
    Type    *InvoiceEventPaymentMethod `json:"type"`
}

func (p *tempPaymentMethodExternal) validate() error {
    var errs []string
    if p.Kind == nil {
        errs = append(errs, "required field `kind` is missing for type `Payment Method External`")
    }
    if p.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Payment Method External`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
