package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// PaymentMethodExternal represents a PaymentMethodExternal struct.
type PaymentMethodExternal struct {
    Details              *string                   `json:"details"`
    Kind                 string                    `json:"kind"`
    Memo                 *string                   `json:"memo"`
    Type                 InvoiceEventPaymentMethod `json:"type"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentMethodExternal.
// It customizes the JSON marshaling process for PaymentMethodExternal objects.
func (p PaymentMethodExternal) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentMethodExternal object to a map representation for JSON marshaling.
func (p PaymentMethodExternal) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
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
    var temp paymentMethodExternal
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "details", "kind", "memo", "type")
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

// paymentMethodExternal is a temporary struct used for validating the fields of PaymentMethodExternal.
type paymentMethodExternal  struct {
    Details *string                    `json:"details"`
    Kind    *string                    `json:"kind"`
    Memo    *string                    `json:"memo"`
    Type    *InvoiceEventPaymentMethod `json:"type"`
}

func (p *paymentMethodExternal) validate() error {
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
    return errors.New(strings.Join(errs, "\n"))
}
