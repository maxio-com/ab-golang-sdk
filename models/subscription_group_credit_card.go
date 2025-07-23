// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// SubscriptionGroupCreditCard represents a SubscriptionGroupCreditCard struct.
type SubscriptionGroupCreditCard struct {
	ChargifyToken *string `json:"chargify_token,omitempty"`
	VaultToken    *string `json:"vault_token,omitempty"`
	// The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing.
	CurrentVault    *CreditCardVault                            `json:"current_vault,omitempty"`
	GatewayHandle   *string                                     `json:"gateway_handle,omitempty"`
	FirstName       *string                                     `json:"first_name,omitempty"`
	LastName        *string                                     `json:"last_name,omitempty"`
	BillingAddress  *string                                     `json:"billing_address,omitempty"`
	BillingAddress2 *string                                     `json:"billing_address_2,omitempty"`
	BillingCity     *string                                     `json:"billing_city,omitempty"`
	BillingState    *string                                     `json:"billing_state,omitempty"`
	BillingZip      *string                                     `json:"billing_zip,omitempty"`
	BillingCountry  *string                                     `json:"billing_country,omitempty"`
	FullNumber      *SubscriptionGroupCreditCardFullNumber      `json:"full_number,omitempty"`
	ExpirationMonth *SubscriptionGroupCreditCardExpirationMonth `json:"expiration_month,omitempty"`
	ExpirationYear  *SubscriptionGroupCreditCardExpirationYear  `json:"expiration_year,omitempty"`
	LastFour        *string                                     `json:"last_four,omitempty"`
	// The type of card used.
	CardType             *CardType              `json:"card_type,omitempty"`
	CustomerVaultToken   *string                `json:"customer_vault_token,omitempty"`
	Cvv                  *string                `json:"cvv,omitempty"`
	PaymentType          *string                `json:"payment_type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionGroupCreditCard,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroupCreditCard) String() string {
	return fmt.Sprintf(
		"SubscriptionGroupCreditCard[ChargifyToken=%v, VaultToken=%v, CurrentVault=%v, GatewayHandle=%v, FirstName=%v, LastName=%v, BillingAddress=%v, BillingAddress2=%v, BillingCity=%v, BillingState=%v, BillingZip=%v, BillingCountry=%v, FullNumber=%v, ExpirationMonth=%v, ExpirationYear=%v, LastFour=%v, CardType=%v, CustomerVaultToken=%v, Cvv=%v, PaymentType=%v, AdditionalProperties=%v]",
		s.ChargifyToken, s.VaultToken, s.CurrentVault, s.GatewayHandle, s.FirstName, s.LastName, s.BillingAddress, s.BillingAddress2, s.BillingCity, s.BillingState, s.BillingZip, s.BillingCountry, s.FullNumber, s.ExpirationMonth, s.ExpirationYear, s.LastFour, s.CardType, s.CustomerVaultToken, s.Cvv, s.PaymentType, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupCreditCard.
// It customizes the JSON marshaling process for SubscriptionGroupCreditCard objects.
func (s SubscriptionGroupCreditCard) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"chargify_token", "vault_token", "current_vault", "gateway_handle", "first_name", "last_name", "billing_address", "billing_address_2", "billing_city", "billing_state", "billing_zip", "billing_country", "full_number", "expiration_month", "expiration_year", "last_four", "card_type", "customer_vault_token", "cvv", "payment_type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupCreditCard object to a map representation for JSON marshaling.
func (s SubscriptionGroupCreditCard) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.ChargifyToken != nil {
		structMap["chargify_token"] = s.ChargifyToken
	}
	if s.VaultToken != nil {
		structMap["vault_token"] = s.VaultToken
	}
	if s.CurrentVault != nil {
		structMap["current_vault"] = s.CurrentVault
	}
	if s.GatewayHandle != nil {
		structMap["gateway_handle"] = s.GatewayHandle
	}
	if s.FirstName != nil {
		structMap["first_name"] = s.FirstName
	}
	if s.LastName != nil {
		structMap["last_name"] = s.LastName
	}
	if s.BillingAddress != nil {
		structMap["billing_address"] = s.BillingAddress
	}
	if s.BillingAddress2 != nil {
		structMap["billing_address_2"] = s.BillingAddress2
	}
	if s.BillingCity != nil {
		structMap["billing_city"] = s.BillingCity
	}
	if s.BillingState != nil {
		structMap["billing_state"] = s.BillingState
	}
	if s.BillingZip != nil {
		structMap["billing_zip"] = s.BillingZip
	}
	if s.BillingCountry != nil {
		structMap["billing_country"] = s.BillingCountry
	}
	if s.FullNumber != nil {
		structMap["full_number"] = s.FullNumber.toMap()
	}
	if s.ExpirationMonth != nil {
		structMap["expiration_month"] = s.ExpirationMonth.toMap()
	}
	if s.ExpirationYear != nil {
		structMap["expiration_year"] = s.ExpirationYear.toMap()
	}
	if s.LastFour != nil {
		structMap["last_four"] = s.LastFour
	}
	if s.CardType != nil {
		structMap["card_type"] = s.CardType
	}
	if s.CustomerVaultToken != nil {
		structMap["customer_vault_token"] = s.CustomerVaultToken
	}
	if s.Cvv != nil {
		structMap["cvv"] = s.Cvv
	}
	if s.PaymentType != nil {
		structMap["payment_type"] = s.PaymentType
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupCreditCard.
// It customizes the JSON unmarshaling process for SubscriptionGroupCreditCard objects.
func (s *SubscriptionGroupCreditCard) UnmarshalJSON(input []byte) error {
	var temp tempSubscriptionGroupCreditCard
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "chargify_token", "vault_token", "current_vault", "gateway_handle", "first_name", "last_name", "billing_address", "billing_address_2", "billing_city", "billing_state", "billing_zip", "billing_country", "full_number", "expiration_month", "expiration_year", "last_four", "card_type", "customer_vault_token", "cvv", "payment_type")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.ChargifyToken = temp.ChargifyToken
	s.VaultToken = temp.VaultToken
	s.CurrentVault = temp.CurrentVault
	s.GatewayHandle = temp.GatewayHandle
	s.FirstName = temp.FirstName
	s.LastName = temp.LastName
	s.BillingAddress = temp.BillingAddress
	s.BillingAddress2 = temp.BillingAddress2
	s.BillingCity = temp.BillingCity
	s.BillingState = temp.BillingState
	s.BillingZip = temp.BillingZip
	s.BillingCountry = temp.BillingCountry
	s.FullNumber = temp.FullNumber
	s.ExpirationMonth = temp.ExpirationMonth
	s.ExpirationYear = temp.ExpirationYear
	s.LastFour = temp.LastFour
	s.CardType = temp.CardType
	s.CustomerVaultToken = temp.CustomerVaultToken
	s.Cvv = temp.Cvv
	s.PaymentType = temp.PaymentType
	return nil
}

// tempSubscriptionGroupCreditCard is a temporary struct used for validating the fields of SubscriptionGroupCreditCard.
type tempSubscriptionGroupCreditCard struct {
	ChargifyToken      *string                                     `json:"chargify_token,omitempty"`
	VaultToken         *string                                     `json:"vault_token,omitempty"`
	CurrentVault       *CreditCardVault                            `json:"current_vault,omitempty"`
	GatewayHandle      *string                                     `json:"gateway_handle,omitempty"`
	FirstName          *string                                     `json:"first_name,omitempty"`
	LastName           *string                                     `json:"last_name,omitempty"`
	BillingAddress     *string                                     `json:"billing_address,omitempty"`
	BillingAddress2    *string                                     `json:"billing_address_2,omitempty"`
	BillingCity        *string                                     `json:"billing_city,omitempty"`
	BillingState       *string                                     `json:"billing_state,omitempty"`
	BillingZip         *string                                     `json:"billing_zip,omitempty"`
	BillingCountry     *string                                     `json:"billing_country,omitempty"`
	FullNumber         *SubscriptionGroupCreditCardFullNumber      `json:"full_number,omitempty"`
	ExpirationMonth    *SubscriptionGroupCreditCardExpirationMonth `json:"expiration_month,omitempty"`
	ExpirationYear     *SubscriptionGroupCreditCardExpirationYear  `json:"expiration_year,omitempty"`
	LastFour           *string                                     `json:"last_four,omitempty"`
	CardType           *CardType                                   `json:"card_type,omitempty"`
	CustomerVaultToken *string                                     `json:"customer_vault_token,omitempty"`
	Cvv                *string                                     `json:"cvv,omitempty"`
	PaymentType        *string                                     `json:"payment_type,omitempty"`
}
