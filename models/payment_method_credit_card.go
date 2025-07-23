// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// PaymentMethodCreditCard represents a PaymentMethodCreditCard struct.
type PaymentMethodCreditCard struct {
	CardBrand            string                    `json:"card_brand"`
	CardExpiration       *string                   `json:"card_expiration,omitempty"`
	LastFour             Optional[string]          `json:"last_four"`
	MaskedCardNumber     string                    `json:"masked_card_number"`
	Type                 InvoiceEventPaymentMethod `json:"type"`
	AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for PaymentMethodCreditCard,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PaymentMethodCreditCard) String() string {
	return fmt.Sprintf(
		"PaymentMethodCreditCard[CardBrand=%v, CardExpiration=%v, LastFour=%v, MaskedCardNumber=%v, Type=%v, AdditionalProperties=%v]",
		p.CardBrand, p.CardExpiration, p.LastFour, p.MaskedCardNumber, p.Type, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PaymentMethodCreditCard.
// It customizes the JSON marshaling process for PaymentMethodCreditCard objects.
func (p PaymentMethodCreditCard) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(p.AdditionalProperties,
		"card_brand", "card_expiration", "last_four", "masked_card_number", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(p.toMap())
}

// toMap converts the PaymentMethodCreditCard object to a map representation for JSON marshaling.
func (p PaymentMethodCreditCard) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, p.AdditionalProperties)
	structMap["card_brand"] = p.CardBrand
	if p.CardExpiration != nil {
		structMap["card_expiration"] = p.CardExpiration
	}
	if p.LastFour.IsValueSet() {
		if p.LastFour.Value() != nil {
			structMap["last_four"] = p.LastFour.Value()
		} else {
			structMap["last_four"] = nil
		}
	}
	structMap["masked_card_number"] = p.MaskedCardNumber
	structMap["type"] = p.Type
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentMethodCreditCard.
// It customizes the JSON unmarshaling process for PaymentMethodCreditCard objects.
func (p *PaymentMethodCreditCard) UnmarshalJSON(input []byte) error {
	var temp tempPaymentMethodCreditCard
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "card_brand", "card_expiration", "last_four", "masked_card_number", "type")
	if err != nil {
		return err
	}
	p.AdditionalProperties = additionalProperties

	p.CardBrand = *temp.CardBrand
	p.CardExpiration = temp.CardExpiration
	p.LastFour = temp.LastFour
	p.MaskedCardNumber = *temp.MaskedCardNumber
	p.Type = *temp.Type
	return nil
}

// tempPaymentMethodCreditCard is a temporary struct used for validating the fields of PaymentMethodCreditCard.
type tempPaymentMethodCreditCard struct {
	CardBrand        *string                    `json:"card_brand"`
	CardExpiration   *string                    `json:"card_expiration,omitempty"`
	LastFour         Optional[string]           `json:"last_four"`
	MaskedCardNumber *string                    `json:"masked_card_number"`
	Type             *InvoiceEventPaymentMethod `json:"type"`
}

func (p *tempPaymentMethodCreditCard) validate() error {
	var errs []string
	if p.CardBrand == nil {
		errs = append(errs, "required field `card_brand` is missing for type `Payment Method Credit Card`")
	}
	if p.MaskedCardNumber == nil {
		errs = append(errs, "required field `masked_card_number` is missing for type `Payment Method Credit Card`")
	}
	if p.Type == nil {
		errs = append(errs, "required field `type` is missing for type `Payment Method Credit Card`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
