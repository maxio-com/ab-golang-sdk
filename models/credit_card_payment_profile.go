package models

import (
	"encoding/json"
)

// CreditCardPaymentProfile represents a CreditCardPaymentProfile struct.
type CreditCardPaymentProfile struct {
	// The Chargify-assigned ID of the stored card. This value can be used as an input to payment_profile_id when creating a subscription, in order to re-use a stored payment profile for the same customer.
	Id *int `json:"id,omitempty"`
	// The first name of the card holder.
	FirstName *string `json:"first_name,omitempty"`
	// The last name of the card holder.
	LastName *string `json:"last_name,omitempty"`
	// A string representation of the credit card number with all but the last 4 digits masked with X’s (i.e. ‘XXXX-XXXX-XXXX-1234’).
	MaskedCardNumber string `json:"masked_card_number"`
	// The type of card used.
	CardType *CardType `json:"card_type,omitempty"`
	// An integer representing the expiration month of the card(1 – 12).
	ExpirationMonth *int `json:"expiration_month,omitempty"`
	// An integer representing the 4-digit expiration year of the card(i.e. ‘2012’).
	ExpirationYear *int `json:"expiration_year,omitempty"`
	// The Chargify-assigned id for the customer record to which the card belongs.
	CustomerId *int `json:"customer_id,omitempty"`
	// The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing.
	CurrentVault *CurrentVault `json:"current_vault,omitempty"`
	// The “token” provided by your vault storage for an already stored payment profile.
	VaultToken Optional[string] `json:"vault_token"`
	// The current billing street address for the card.
	BillingAddress Optional[string] `json:"billing_address"`
	// The current billing address city for the card.
	BillingCity Optional[string] `json:"billing_city"`
	// The current billing address state for the card.
	BillingState Optional[string] `json:"billing_state"`
	// The current billing address zip code for the card.
	BillingZip Optional[string] `json:"billing_zip"`
	// The current billing address country for the card.
	BillingCountry Optional[string] `json:"billing_country"`
	// (only for Authorize.Net CIM storage): the customerProfileId for the owner of the customerPaymentProfileId provided as the vault_token.
	CustomerVaultToken Optional[string] `json:"customer_vault_token"`
	// The current billing street address, second line, for the card.
	BillingAddress2 Optional[string] `json:"billing_address_2"`
	PaymentType     *PaymentType     `json:"payment_type,omitempty"`
	Disabled        *bool            `json:"disabled,omitempty"`
	// Token received after sending billing information using chargify.js. This token will only be received if passed as a sole attribute of credit_card_attributes (i.e. tok_9g6hw85pnpt6knmskpwp4ttt)
	ChargifyToken        *string       `json:"chargify_token,omitempty"`
	SiteGatewaySettingId Optional[int] `json:"site_gateway_setting_id"`
	// An identifier of connected gateway.
	GatewayHandle Optional[string] `json:"gateway_handle"`
}

// MarshalJSON implements the json.Marshaler interface for CreditCardPaymentProfile.
// It customizes the JSON marshaling process for CreditCardPaymentProfile objects.
func (c *CreditCardPaymentProfile) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreditCardPaymentProfile object to a map representation for JSON marshaling.
func (c *CreditCardPaymentProfile) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Id != nil {
		structMap["id"] = c.Id
	}
	if c.FirstName != nil {
		structMap["first_name"] = c.FirstName
	}
	if c.LastName != nil {
		structMap["last_name"] = c.LastName
	}
	structMap["masked_card_number"] = c.MaskedCardNumber
	if c.CardType != nil {
		structMap["card_type"] = c.CardType
	}
	if c.ExpirationMonth != nil {
		structMap["expiration_month"] = c.ExpirationMonth
	}
	if c.ExpirationYear != nil {
		structMap["expiration_year"] = c.ExpirationYear
	}
	if c.CustomerId != nil {
		structMap["customer_id"] = c.CustomerId
	}
	if c.CurrentVault != nil {
		structMap["current_vault"] = c.CurrentVault
	}
	if c.VaultToken.IsValueSet() {
		structMap["vault_token"] = c.VaultToken.Value()
	}
	if c.BillingAddress.IsValueSet() {
		structMap["billing_address"] = c.BillingAddress.Value()
	}
	if c.BillingCity.IsValueSet() {
		structMap["billing_city"] = c.BillingCity.Value()
	}
	if c.BillingState.IsValueSet() {
		structMap["billing_state"] = c.BillingState.Value()
	}
	if c.BillingZip.IsValueSet() {
		structMap["billing_zip"] = c.BillingZip.Value()
	}
	if c.BillingCountry.IsValueSet() {
		structMap["billing_country"] = c.BillingCountry.Value()
	}
	if c.CustomerVaultToken.IsValueSet() {
		structMap["customer_vault_token"] = c.CustomerVaultToken.Value()
	}
	if c.BillingAddress2.IsValueSet() {
		structMap["billing_address_2"] = c.BillingAddress2.Value()
	}
	if c.PaymentType != nil {
		structMap["payment_type"] = c.PaymentType
	}
	if c.Disabled != nil {
		structMap["disabled"] = c.Disabled
	}
	if c.ChargifyToken != nil {
		structMap["chargify_token"] = c.ChargifyToken
	}
	if c.SiteGatewaySettingId.IsValueSet() {
		structMap["site_gateway_setting_id"] = c.SiteGatewaySettingId.Value()
	}
	if c.GatewayHandle.IsValueSet() {
		structMap["gateway_handle"] = c.GatewayHandle.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreditCardPaymentProfile.
// It customizes the JSON unmarshaling process for CreditCardPaymentProfile objects.
func (c *CreditCardPaymentProfile) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                   *int             `json:"id,omitempty"`
		FirstName            *string          `json:"first_name,omitempty"`
		LastName             *string          `json:"last_name,omitempty"`
		MaskedCardNumber     string           `json:"masked_card_number"`
		CardType             *CardType        `json:"card_type,omitempty"`
		ExpirationMonth      *int             `json:"expiration_month,omitempty"`
		ExpirationYear       *int             `json:"expiration_year,omitempty"`
		CustomerId           *int             `json:"customer_id,omitempty"`
		CurrentVault         *CurrentVault    `json:"current_vault,omitempty"`
		VaultToken           Optional[string] `json:"vault_token"`
		BillingAddress       Optional[string] `json:"billing_address"`
		BillingCity          Optional[string] `json:"billing_city"`
		BillingState         Optional[string] `json:"billing_state"`
		BillingZip           Optional[string] `json:"billing_zip"`
		BillingCountry       Optional[string] `json:"billing_country"`
		CustomerVaultToken   Optional[string] `json:"customer_vault_token"`
		BillingAddress2      Optional[string] `json:"billing_address_2"`
		PaymentType          *PaymentType     `json:"payment_type,omitempty"`
		Disabled             *bool            `json:"disabled,omitempty"`
		ChargifyToken        *string          `json:"chargify_token,omitempty"`
		SiteGatewaySettingId Optional[int]    `json:"site_gateway_setting_id"`
		GatewayHandle        Optional[string] `json:"gateway_handle"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Id = temp.Id
	c.FirstName = temp.FirstName
	c.LastName = temp.LastName
	c.MaskedCardNumber = temp.MaskedCardNumber
	c.CardType = temp.CardType
	c.ExpirationMonth = temp.ExpirationMonth
	c.ExpirationYear = temp.ExpirationYear
	c.CustomerId = temp.CustomerId
	c.CurrentVault = temp.CurrentVault
	c.VaultToken = temp.VaultToken
	c.BillingAddress = temp.BillingAddress
	c.BillingCity = temp.BillingCity
	c.BillingState = temp.BillingState
	c.BillingZip = temp.BillingZip
	c.BillingCountry = temp.BillingCountry
	c.CustomerVaultToken = temp.CustomerVaultToken
	c.BillingAddress2 = temp.BillingAddress2
	c.PaymentType = temp.PaymentType
	c.Disabled = temp.Disabled
	c.ChargifyToken = temp.ChargifyToken
	c.SiteGatewaySettingId = temp.SiteGatewaySettingId
	c.GatewayHandle = temp.GatewayHandle
	return nil
}
