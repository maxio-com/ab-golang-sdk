/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// InvoiceEventPayment represents a InvoiceEventPayment struct.
// A nested data structure detailing the method of payment
type InvoiceEventPayment struct {
    Type                 *string          `json:"type,omitempty"`
    MaskedAccountNumber  *string          `json:"masked_account_number,omitempty"`
    MaskedRoutingNumber  *string          `json:"masked_routing_number,omitempty"`
    CardBrand            *string          `json:"card_brand,omitempty"`
    CardExpiration       *string          `json:"card_expiration,omitempty"`
    LastFour             Optional[string] `json:"last_four"`
    MaskedCardNumber     *string          `json:"masked_card_number,omitempty"`
    Details              Optional[string] `json:"details"`
    Kind                 *string          `json:"kind,omitempty"`
    Memo                 Optional[string] `json:"memo"`
    Email                *string          `json:"email,omitempty"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceEventPayment.
// It customizes the JSON marshaling process for InvoiceEventPayment objects.
func (i InvoiceEventPayment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceEventPayment object to a map representation for JSON marshaling.
func (i InvoiceEventPayment) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Type != nil {
        structMap["type"] = *i.Type
    } else {
        structMap["type"] = "Invoice Event Payment"
    }
    if i.MaskedAccountNumber != nil {
        structMap["masked_account_number"] = i.MaskedAccountNumber
    }
    if i.MaskedRoutingNumber != nil {
        structMap["masked_routing_number"] = i.MaskedRoutingNumber
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
    if i.Details.IsValueSet() {
        if i.Details.Value() != nil {
            structMap["details"] = i.Details.Value()
        } else {
            structMap["details"] = nil
        }
    }
    if i.Kind != nil {
        structMap["kind"] = i.Kind
    }
    if i.Memo.IsValueSet() {
        if i.Memo.Value() != nil {
            structMap["memo"] = i.Memo.Value()
        } else {
            structMap["memo"] = nil
        }
    }
    if i.Email != nil {
        structMap["email"] = i.Email
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceEventPayment.
// It customizes the JSON unmarshaling process for InvoiceEventPayment objects.
func (i *InvoiceEventPayment) UnmarshalJSON(input []byte) error {
    var temp invoiceEventPayment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "type", "masked_account_number", "masked_routing_number", "card_brand", "card_expiration", "last_four", "masked_card_number", "details", "kind", "memo", "email")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Type = temp.Type
    i.MaskedAccountNumber = temp.MaskedAccountNumber
    i.MaskedRoutingNumber = temp.MaskedRoutingNumber
    i.CardBrand = temp.CardBrand
    i.CardExpiration = temp.CardExpiration
    i.LastFour = temp.LastFour
    i.MaskedCardNumber = temp.MaskedCardNumber
    i.Details = temp.Details
    i.Kind = temp.Kind
    i.Memo = temp.Memo
    i.Email = temp.Email
    return nil
}

// TODO
type invoiceEventPayment  struct {
    Type                *string          `json:"type,omitempty"`
    MaskedAccountNumber *string          `json:"masked_account_number,omitempty"`
    MaskedRoutingNumber *string          `json:"masked_routing_number,omitempty"`
    CardBrand           *string          `json:"card_brand,omitempty"`
    CardExpiration      *string          `json:"card_expiration,omitempty"`
    LastFour            Optional[string] `json:"last_four"`
    MaskedCardNumber    *string          `json:"masked_card_number,omitempty"`
    Details             Optional[string] `json:"details"`
    Kind                *string          `json:"kind,omitempty"`
    Memo                Optional[string] `json:"memo"`
    Email               *string          `json:"email,omitempty"`
}
