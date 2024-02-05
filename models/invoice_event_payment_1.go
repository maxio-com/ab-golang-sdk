package models

import (
    "encoding/json"
)

// InvoiceEventPayment1 represents a InvoiceEventPayment1 struct.
// A nested data structure detailing the method of payment
type InvoiceEventPayment1 struct {
    Type                *string          `json:"type,omitempty"`
    MaskedAccountNumber string           `json:"masked_account_number"`
    MaskedRoutingNumber string           `json:"masked_routing_number"`
    CardBrand           string           `json:"card_brand"`
    CardExpiration      *string          `json:"card_expiration,omitempty"`
    LastFour            Optional[string] `json:"last_four"`
    MaskedCardNumber    string           `json:"masked_card_number"`
    Details             *string          `json:"details"`
    Kind                string           `json:"kind"`
    Memo                *string          `json:"memo"`
    Email               string           `json:"email"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceEventPayment1.
// It customizes the JSON marshaling process for InvoiceEventPayment1 objects.
func (i *InvoiceEventPayment1) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceEventPayment1 object to a map representation for JSON marshaling.
func (i *InvoiceEventPayment1) toMap() map[string]any {
    structMap := make(map[string]any)
    if i.Type != nil {
        structMap["type"] = *i.Type
    } else {
        structMap["type"] = "Invoice Event Payment1"
    }
    structMap["masked_account_number"] = i.MaskedAccountNumber
    structMap["masked_routing_number"] = i.MaskedRoutingNumber
    structMap["card_brand"] = i.CardBrand
    if i.CardExpiration != nil {
        structMap["card_expiration"] = i.CardExpiration
    }
    if i.LastFour.IsValueSet() {
        structMap["last_four"] = i.LastFour.Value()
    }
    structMap["masked_card_number"] = i.MaskedCardNumber
    structMap["details"] = i.Details
    structMap["kind"] = i.Kind
    structMap["memo"] = i.Memo
    structMap["email"] = i.Email
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceEventPayment1.
// It customizes the JSON unmarshaling process for InvoiceEventPayment1 objects.
func (i *InvoiceEventPayment1) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Type                *string          `json:"type,omitempty"`
        MaskedAccountNumber string           `json:"masked_account_number"`
        MaskedRoutingNumber string           `json:"masked_routing_number"`
        CardBrand           string           `json:"card_brand"`
        CardExpiration      *string          `json:"card_expiration,omitempty"`
        LastFour            Optional[string] `json:"last_four"`
        MaskedCardNumber    string           `json:"masked_card_number"`
        Details             *string          `json:"details"`
        Kind                string           `json:"kind"`
        Memo                *string          `json:"memo"`
        Email               string           `json:"email"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
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
