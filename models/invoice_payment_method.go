/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// InvoicePaymentMethod represents a InvoicePaymentMethod struct.
type InvoicePaymentMethod struct {
    Details              *string                `json:"details,omitempty"`
    Kind                 *string                `json:"kind,omitempty"`
    Memo                 *string                `json:"memo,omitempty"`
    Type                 *string                `json:"type,omitempty"`
    CardBrand            *string                `json:"card_brand,omitempty"`
    CardExpiration       *string                `json:"card_expiration,omitempty"`
    LastFour             Optional[string]       `json:"last_four"`
    MaskedCardNumber     *string                `json:"masked_card_number,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoicePaymentMethod.
// It customizes the JSON marshaling process for InvoicePaymentMethod objects.
func (i InvoicePaymentMethod) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "details", "kind", "memo", "type", "card_brand", "card_expiration", "last_four", "masked_card_number"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoicePaymentMethod object to a map representation for JSON marshaling.
func (i InvoicePaymentMethod) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Details != nil {
        structMap["details"] = i.Details
    }
    if i.Kind != nil {
        structMap["kind"] = i.Kind
    }
    if i.Memo != nil {
        structMap["memo"] = i.Memo
    }
    if i.Type != nil {
        structMap["type"] = i.Type
    }
    if i.CardBrand != nil {
        structMap["card_brand"] = i.CardBrand
    }
    if i.CardExpiration != nil {
        structMap["card_expiration"] = i.CardExpiration
    }
    if i.LastFour.IsValueSet() {
        if i.LastFour.Value() != nil {
            structMap["last_four"] = i.LastFour.Value()
        } else {
            structMap["last_four"] = nil
        }
    }
    if i.MaskedCardNumber != nil {
        structMap["masked_card_number"] = i.MaskedCardNumber
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoicePaymentMethod.
// It customizes the JSON unmarshaling process for InvoicePaymentMethod objects.
func (i *InvoicePaymentMethod) UnmarshalJSON(input []byte) error {
    var temp tempInvoicePaymentMethod
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "details", "kind", "memo", "type", "card_brand", "card_expiration", "last_four", "masked_card_number")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Details = temp.Details
    i.Kind = temp.Kind
    i.Memo = temp.Memo
    i.Type = temp.Type
    i.CardBrand = temp.CardBrand
    i.CardExpiration = temp.CardExpiration
    i.LastFour = temp.LastFour
    i.MaskedCardNumber = temp.MaskedCardNumber
    return nil
}

// tempInvoicePaymentMethod is a temporary struct used for validating the fields of InvoicePaymentMethod.
type tempInvoicePaymentMethod  struct {
    Details          *string          `json:"details,omitempty"`
    Kind             *string          `json:"kind,omitempty"`
    Memo             *string          `json:"memo,omitempty"`
    Type             *string          `json:"type,omitempty"`
    CardBrand        *string          `json:"card_brand,omitempty"`
    CardExpiration   *string          `json:"card_expiration,omitempty"`
    LastFour         Optional[string] `json:"last_four"`
    MaskedCardNumber *string          `json:"masked_card_number,omitempty"`
}
