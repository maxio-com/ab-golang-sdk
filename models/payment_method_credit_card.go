package models

import (
    "encoding/json"
)

// PaymentMethodCreditCard represents a PaymentMethodCreditCard struct.
type PaymentMethodCreditCard struct {
    CardBrand        string                    `json:"card_brand"`
    CardExpiration   *string                   `json:"card_expiration,omitempty"`
    LastFour         Optional[string]          `json:"last_four"`
    MaskedCardNumber string                    `json:"masked_card_number"`
    Type             InvoiceEventPaymentMethod `json:"type"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentMethodCreditCard.
// It customizes the JSON marshaling process for PaymentMethodCreditCard objects.
func (p *PaymentMethodCreditCard) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentMethodCreditCard object to a map representation for JSON marshaling.
func (p *PaymentMethodCreditCard) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["card_brand"] = p.CardBrand
    if p.CardExpiration != nil {
        structMap["card_expiration"] = p.CardExpiration
    }
    if p.LastFour.IsValueSet() {
        structMap["last_four"] = p.LastFour.Value()
    }
    structMap["masked_card_number"] = p.MaskedCardNumber
    structMap["type"] = p.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentMethodCreditCard.
// It customizes the JSON unmarshaling process for PaymentMethodCreditCard objects.
func (p *PaymentMethodCreditCard) UnmarshalJSON(input []byte) error {
    temp := &struct {
        CardBrand        string                    `json:"card_brand"`
        CardExpiration   *string                   `json:"card_expiration,omitempty"`
        LastFour         Optional[string]          `json:"last_four"`
        MaskedCardNumber string                    `json:"masked_card_number"`
        Type             InvoiceEventPaymentMethod `json:"type"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.CardBrand = temp.CardBrand
    p.CardExpiration = temp.CardExpiration
    p.LastFour = temp.LastFour
    p.MaskedCardNumber = temp.MaskedCardNumber
    p.Type = temp.Type
    return nil
}
