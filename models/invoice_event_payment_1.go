package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// InvoiceEventPayment1 represents a InvoiceEventPayment1 struct.
// A nested data structure detailing the method of payment
type InvoiceEventPayment1 struct {
    Type                 *string          `json:"type,omitempty"`
    MaskedAccountNumber  string           `json:"masked_account_number"`
    MaskedRoutingNumber  string           `json:"masked_routing_number"`
    CardBrand            string           `json:"card_brand"`
    CardExpiration       *string          `json:"card_expiration,omitempty"`
    LastFour             Optional[string] `json:"last_four"`
    MaskedCardNumber     string           `json:"masked_card_number"`
    Details              *string          `json:"details"`
    Kind                 string           `json:"kind"`
    Memo                 *string          `json:"memo"`
    Email                string           `json:"email"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceEventPayment1.
// It customizes the JSON marshaling process for InvoiceEventPayment1 objects.
func (i InvoiceEventPayment1) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceEventPayment1 object to a map representation for JSON marshaling.
func (i InvoiceEventPayment1) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
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
        if i.LastFour.Value() != nil {
            structMap["last_four"] = i.LastFour.Value()
        } else {
            structMap["last_four"] = nil
        }
    }
    structMap["masked_card_number"] = i.MaskedCardNumber
    if i.Details != nil {
        structMap["details"] = i.Details
    } else {
        structMap["details"] = nil
    }
    structMap["kind"] = i.Kind
    if i.Memo != nil {
        structMap["memo"] = i.Memo
    } else {
        structMap["memo"] = nil
    }
    structMap["email"] = i.Email
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceEventPayment1.
// It customizes the JSON unmarshaling process for InvoiceEventPayment1 objects.
func (i *InvoiceEventPayment1) UnmarshalJSON(input []byte) error {
    var temp invoiceEventPayment1
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "type", "masked_account_number", "masked_routing_number", "card_brand", "card_expiration", "last_four", "masked_card_number", "details", "kind", "memo", "email")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Type = temp.Type
    i.MaskedAccountNumber = *temp.MaskedAccountNumber
    i.MaskedRoutingNumber = *temp.MaskedRoutingNumber
    i.CardBrand = *temp.CardBrand
    i.CardExpiration = temp.CardExpiration
    i.LastFour = temp.LastFour
    i.MaskedCardNumber = *temp.MaskedCardNumber
    i.Details = temp.Details
    i.Kind = *temp.Kind
    i.Memo = temp.Memo
    i.Email = *temp.Email
    return nil
}

// invoiceEventPayment1 is a temporary struct used for validating the fields of InvoiceEventPayment1.
type invoiceEventPayment1  struct {
    Type                *string          `json:"type,omitempty"`
    MaskedAccountNumber *string          `json:"masked_account_number"`
    MaskedRoutingNumber *string          `json:"masked_routing_number"`
    CardBrand           *string          `json:"card_brand"`
    CardExpiration      *string          `json:"card_expiration,omitempty"`
    LastFour            Optional[string] `json:"last_four"`
    MaskedCardNumber    *string          `json:"masked_card_number"`
    Details             *string          `json:"details"`
    Kind                *string          `json:"kind"`
    Memo                *string          `json:"memo"`
    Email               *string          `json:"email"`
}

func (i *invoiceEventPayment1) validate() error {
    var errs []string
    if i.MaskedAccountNumber == nil {
        errs = append(errs, "required field `masked_account_number` is missing for type `Invoice Event Payment1`")
    }
    if i.MaskedRoutingNumber == nil {
        errs = append(errs, "required field `masked_routing_number` is missing for type `Invoice Event Payment1`")
    }
    if i.CardBrand == nil {
        errs = append(errs, "required field `card_brand` is missing for type `Invoice Event Payment1`")
    }
    if i.MaskedCardNumber == nil {
        errs = append(errs, "required field `masked_card_number` is missing for type `Invoice Event Payment1`")
    }
    if i.Kind == nil {
        errs = append(errs, "required field `kind` is missing for type `Invoice Event Payment1`")
    }
    if i.Email == nil {
        errs = append(errs, "required field `email` is missing for type `Invoice Event Payment1`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
