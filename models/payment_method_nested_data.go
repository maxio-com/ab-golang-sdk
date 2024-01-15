package models

import (
    "encoding/json"
)

// PaymentMethodNestedData represents a PaymentMethodNestedData struct.
// A nested data structure detailing the method of payment
type PaymentMethodNestedData struct {
    Type                *string          `json:"type,omitempty"`
    MaskedAccountNumber *string          `json:"masked_account_number,omitempty"`
    MaskedRoutingNumber *string          `json:"masked_routing_number,omitempty"`
    CardBrand           *string          `json:"card_brand,omitempty"`
    CardExpiration      *string          `json:"card_expiration,omitempty"`
    LastFour            Optional[string] `json:"last_four"`
    MaskedCardNumber    *string          `json:"masked_card_number,omitempty"`
    Details             *string          `json:"details,omitempty"`
    Kind                *string          `json:"kind,omitempty"`
    Memo                *string          `json:"memo,omitempty"`
    Email               *string          `json:"email,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentMethodNestedData.
// It customizes the JSON marshaling process for PaymentMethodNestedData objects.
func (p *PaymentMethodNestedData) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentMethodNestedData object to a map representation for JSON marshaling.
func (p *PaymentMethodNestedData) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.Type != nil {
        structMap["type"] = p.Type
    }
    if p.MaskedAccountNumber != nil {
        structMap["masked_account_number"] = p.MaskedAccountNumber
    }
    if p.MaskedRoutingNumber != nil {
        structMap["masked_routing_number"] = p.MaskedRoutingNumber
    }
    if p.CardBrand != nil {
        structMap["card_brand"] = p.CardBrand
    }
    if p.CardExpiration != nil {
        structMap["card_expiration"] = p.CardExpiration
    }
    if p.LastFour.IsValueSet() {
        structMap["last_four"] = p.LastFour.Value()
    }
    if p.MaskedCardNumber != nil {
        structMap["masked_card_number"] = p.MaskedCardNumber
    }
    if p.Details != nil {
        structMap["details"] = p.Details
    }
    if p.Kind != nil {
        structMap["kind"] = p.Kind
    }
    if p.Memo != nil {
        structMap["memo"] = p.Memo
    }
    if p.Email != nil {
        structMap["email"] = p.Email
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentMethodNestedData.
// It customizes the JSON unmarshaling process for PaymentMethodNestedData objects.
func (p *PaymentMethodNestedData) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Type                *string          `json:"type,omitempty"`
        MaskedAccountNumber *string          `json:"masked_account_number,omitempty"`
        MaskedRoutingNumber *string          `json:"masked_routing_number,omitempty"`
        CardBrand           *string          `json:"card_brand,omitempty"`
        CardExpiration      *string          `json:"card_expiration,omitempty"`
        LastFour            Optional[string] `json:"last_four"`
        MaskedCardNumber    *string          `json:"masked_card_number,omitempty"`
        Details             *string          `json:"details,omitempty"`
        Kind                *string          `json:"kind,omitempty"`
        Memo                *string          `json:"memo,omitempty"`
        Email               *string          `json:"email,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Type = temp.Type
    p.MaskedAccountNumber = temp.MaskedAccountNumber
    p.MaskedRoutingNumber = temp.MaskedRoutingNumber
    p.CardBrand = temp.CardBrand
    p.CardExpiration = temp.CardExpiration
    p.LastFour = temp.LastFour
    p.MaskedCardNumber = temp.MaskedCardNumber
    p.Details = temp.Details
    p.Kind = temp.Kind
    p.Memo = temp.Memo
    p.Email = temp.Email
    return nil
}
