package models

import (
    "encoding/json"
)

// PaymentProfile represents a PaymentProfile struct.
type PaymentProfile struct {
    // The Chargify-assigned ID of the stored card. This value can be used as an input to payment_profile_id when creating a subscription, in order to re-use a stored payment profile for the same customer.
    Id                   *int             `json:"id,omitempty"`
    // The first name of the card holder.
    FirstName            *string          `json:"first_name,omitempty"`
    // The last name of the card holder.
    LastName             *string          `json:"last_name,omitempty"`
    // A string representation of the credit card number with all but the last 4 digits masked with X’s (i.e. ‘XXXX-XXXX-XXXX-1234’).
    MaskedCardNumber     *string          `json:"masked_card_number,omitempty"`
    // The type of card used.
    CardType             *CardType        `json:"card_type,omitempty"`
    // An integer representing the expiration month of the card(1 – 12).
    ExpirationMonth      *int             `json:"expiration_month,omitempty"`
    // An integer representing the 4-digit expiration year of the card(i.e. ‘2012’).
    ExpirationYear       *int             `json:"expiration_year,omitempty"`
    // The Chargify-assigned id for the customer record to which the card belongs.
    CustomerId           *int             `json:"customer_id,omitempty"`
    // The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing.
    CurrentVault         *CurrentVault    `json:"current_vault,omitempty"`
    // The “token” provided by your vault storage for an already stored payment profile.
    VaultToken           Optional[string] `json:"vault_token"`
    // The current billing street address for the card.
    BillingAddress       Optional[string] `json:"billing_address"`
    // The current billing address city for the card.
    BillingCity          Optional[string] `json:"billing_city"`
    // The current billing address state for the card.
    BillingState         Optional[string] `json:"billing_state"`
    // The current billing address zip code for the card.
    BillingZip           Optional[string] `json:"billing_zip"`
    // The current billing address country for the card.
    BillingCountry       Optional[string] `json:"billing_country"`
    // (only for Authorize.Net CIM storage): the customerProfileId for the owner of the customerPaymentProfileId provided as the vault_token.
    CustomerVaultToken   Optional[string] `json:"customer_vault_token"`
    // The current billing street address, second line, for the card.
    BillingAddress2      Optional[string] `json:"billing_address_2"`
    PaymentType          *PaymentType     `json:"payment_type,omitempty"`
    Disabled             *bool            `json:"disabled,omitempty"`
    // Token received after sending billing informations using chargify.js.
    ChargifyToken        *string          `json:"chargify_token,omitempty"`
    SiteGatewaySettingId Optional[int]    `json:"site_gateway_setting_id"`
    // An identifier of connected gateway.
    GatewayHandle        Optional[string] `json:"gateway_handle"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentProfile.
// It customizes the JSON marshaling process for PaymentProfile objects.
func (p *PaymentProfile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentProfile object to a map representation for JSON marshaling.
func (p *PaymentProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.Id != nil {
        structMap["id"] = p.Id
    }
    if p.FirstName != nil {
        structMap["first_name"] = p.FirstName
    }
    if p.LastName != nil {
        structMap["last_name"] = p.LastName
    }
    if p.MaskedCardNumber != nil {
        structMap["masked_card_number"] = p.MaskedCardNumber
    }
    if p.CardType != nil {
        structMap["card_type"] = p.CardType
    }
    if p.ExpirationMonth != nil {
        structMap["expiration_month"] = p.ExpirationMonth
    }
    if p.ExpirationYear != nil {
        structMap["expiration_year"] = p.ExpirationYear
    }
    if p.CustomerId != nil {
        structMap["customer_id"] = p.CustomerId
    }
    if p.CurrentVault != nil {
        structMap["current_vault"] = p.CurrentVault
    }
    if p.VaultToken.IsValueSet() {
        structMap["vault_token"] = p.VaultToken.Value()
    }
    if p.BillingAddress.IsValueSet() {
        structMap["billing_address"] = p.BillingAddress.Value()
    }
    if p.BillingCity.IsValueSet() {
        structMap["billing_city"] = p.BillingCity.Value()
    }
    if p.BillingState.IsValueSet() {
        structMap["billing_state"] = p.BillingState.Value()
    }
    if p.BillingZip.IsValueSet() {
        structMap["billing_zip"] = p.BillingZip.Value()
    }
    if p.BillingCountry.IsValueSet() {
        structMap["billing_country"] = p.BillingCountry.Value()
    }
    if p.CustomerVaultToken.IsValueSet() {
        structMap["customer_vault_token"] = p.CustomerVaultToken.Value()
    }
    if p.BillingAddress2.IsValueSet() {
        structMap["billing_address_2"] = p.BillingAddress2.Value()
    }
    if p.PaymentType != nil {
        structMap["payment_type"] = p.PaymentType
    }
    if p.Disabled != nil {
        structMap["disabled"] = p.Disabled
    }
    if p.ChargifyToken != nil {
        structMap["chargify_token"] = p.ChargifyToken
    }
    if p.SiteGatewaySettingId.IsValueSet() {
        structMap["site_gateway_setting_id"] = p.SiteGatewaySettingId.Value()
    }
    if p.GatewayHandle.IsValueSet() {
        structMap["gateway_handle"] = p.GatewayHandle.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentProfile.
// It customizes the JSON unmarshaling process for PaymentProfile objects.
func (p *PaymentProfile) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                   *int             `json:"id,omitempty"`
        FirstName            *string          `json:"first_name,omitempty"`
        LastName             *string          `json:"last_name,omitempty"`
        MaskedCardNumber     *string          `json:"masked_card_number,omitempty"`
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
    
    p.Id = temp.Id
    p.FirstName = temp.FirstName
    p.LastName = temp.LastName
    p.MaskedCardNumber = temp.MaskedCardNumber
    p.CardType = temp.CardType
    p.ExpirationMonth = temp.ExpirationMonth
    p.ExpirationYear = temp.ExpirationYear
    p.CustomerId = temp.CustomerId
    p.CurrentVault = temp.CurrentVault
    p.VaultToken = temp.VaultToken
    p.BillingAddress = temp.BillingAddress
    p.BillingCity = temp.BillingCity
    p.BillingState = temp.BillingState
    p.BillingZip = temp.BillingZip
    p.BillingCountry = temp.BillingCountry
    p.CustomerVaultToken = temp.CustomerVaultToken
    p.BillingAddress2 = temp.BillingAddress2
    p.PaymentType = temp.PaymentType
    p.Disabled = temp.Disabled
    p.ChargifyToken = temp.ChargifyToken
    p.SiteGatewaySettingId = temp.SiteGatewaySettingId
    p.GatewayHandle = temp.GatewayHandle
    return nil
}
