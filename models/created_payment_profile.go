package models

import (
    "encoding/json"
)

// CreatedPaymentProfile represents a CreatedPaymentProfile struct.
type CreatedPaymentProfile struct {
    Id                      *int             `json:"id,omitempty"`
    FirstName               *string          `json:"first_name,omitempty"`
    LastName                *string          `json:"last_name,omitempty"`
    MaskedCardNumber        Optional[string] `json:"masked_card_number"`
    CardType                *string          `json:"card_type,omitempty"`
    ExpirationMonth         *int             `json:"expiration_month,omitempty"`
    ExpirationYear          *int             `json:"expiration_year,omitempty"`
    CustomerId              *int             `json:"customer_id,omitempty"`
    // The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing.
    CurrentVault            *CurrentVault    `json:"current_vault,omitempty"`
    VaultToken              *string          `json:"vault_token,omitempty"`
    BillingAddress          *string          `json:"billing_address,omitempty"`
    BillingCity             *string          `json:"billing_city,omitempty"`
    BillingState            *string          `json:"billing_state,omitempty"`
    BillingZip              *string          `json:"billing_zip,omitempty"`
    BillingCountry          *string          `json:"billing_country,omitempty"`
    CustomerVaultToken      Optional[string] `json:"customer_vault_token"`
    BillingAddress2         Optional[string] `json:"billing_address_2"`
    PaymentType             *string          `json:"payment_type,omitempty"`
    BankName                *string          `json:"bank_name,omitempty"`
    MaskedBankRoutingNumber *string          `json:"masked_bank_routing_number,omitempty"`
    MaskedBankAccountNumber *string          `json:"masked_bank_account_number,omitempty"`
    BankAccountType         *string          `json:"bank_account_type,omitempty"`
    BankAccountHolderType   *string          `json:"bank_account_holder_type,omitempty"`
    Verified                *bool            `json:"verified,omitempty"`
    SiteGatewaySettingId    *int             `json:"site_gateway_setting_id,omitempty"`
    GatewayHandle           *string          `json:"gateway_handle,omitempty"`
    Disabled                *bool            `json:"disabled,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreatedPaymentProfile.
// It customizes the JSON marshaling process for CreatedPaymentProfile objects.
func (c *CreatedPaymentProfile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreatedPaymentProfile object to a map representation for JSON marshaling.
func (c *CreatedPaymentProfile) toMap() map[string]any {
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
    if c.MaskedCardNumber.IsValueSet() {
        structMap["masked_card_number"] = c.MaskedCardNumber.Value()
    }
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
    if c.VaultToken != nil {
        structMap["vault_token"] = c.VaultToken
    }
    if c.BillingAddress != nil {
        structMap["billing_address"] = c.BillingAddress
    }
    if c.BillingCity != nil {
        structMap["billing_city"] = c.BillingCity
    }
    if c.BillingState != nil {
        structMap["billing_state"] = c.BillingState
    }
    if c.BillingZip != nil {
        structMap["billing_zip"] = c.BillingZip
    }
    if c.BillingCountry != nil {
        structMap["billing_country"] = c.BillingCountry
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
    if c.BankName != nil {
        structMap["bank_name"] = c.BankName
    }
    if c.MaskedBankRoutingNumber != nil {
        structMap["masked_bank_routing_number"] = c.MaskedBankRoutingNumber
    }
    if c.MaskedBankAccountNumber != nil {
        structMap["masked_bank_account_number"] = c.MaskedBankAccountNumber
    }
    if c.BankAccountType != nil {
        structMap["bank_account_type"] = c.BankAccountType
    }
    if c.BankAccountHolderType != nil {
        structMap["bank_account_holder_type"] = c.BankAccountHolderType
    }
    if c.Verified != nil {
        structMap["verified"] = c.Verified
    }
    if c.SiteGatewaySettingId != nil {
        structMap["site_gateway_setting_id"] = c.SiteGatewaySettingId
    }
    if c.GatewayHandle != nil {
        structMap["gateway_handle"] = c.GatewayHandle
    }
    if c.Disabled != nil {
        structMap["disabled"] = c.Disabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatedPaymentProfile.
// It customizes the JSON unmarshaling process for CreatedPaymentProfile objects.
func (c *CreatedPaymentProfile) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                      *int             `json:"id,omitempty"`
        FirstName               *string          `json:"first_name,omitempty"`
        LastName                *string          `json:"last_name,omitempty"`
        MaskedCardNumber        Optional[string] `json:"masked_card_number"`
        CardType                *string          `json:"card_type,omitempty"`
        ExpirationMonth         *int             `json:"expiration_month,omitempty"`
        ExpirationYear          *int             `json:"expiration_year,omitempty"`
        CustomerId              *int             `json:"customer_id,omitempty"`
        CurrentVault            *CurrentVault    `json:"current_vault,omitempty"`
        VaultToken              *string          `json:"vault_token,omitempty"`
        BillingAddress          *string          `json:"billing_address,omitempty"`
        BillingCity             *string          `json:"billing_city,omitempty"`
        BillingState            *string          `json:"billing_state,omitempty"`
        BillingZip              *string          `json:"billing_zip,omitempty"`
        BillingCountry          *string          `json:"billing_country,omitempty"`
        CustomerVaultToken      Optional[string] `json:"customer_vault_token"`
        BillingAddress2         Optional[string] `json:"billing_address_2"`
        PaymentType             *string          `json:"payment_type,omitempty"`
        BankName                *string          `json:"bank_name,omitempty"`
        MaskedBankRoutingNumber *string          `json:"masked_bank_routing_number,omitempty"`
        MaskedBankAccountNumber *string          `json:"masked_bank_account_number,omitempty"`
        BankAccountType         *string          `json:"bank_account_type,omitempty"`
        BankAccountHolderType   *string          `json:"bank_account_holder_type,omitempty"`
        Verified                *bool            `json:"verified,omitempty"`
        SiteGatewaySettingId    *int             `json:"site_gateway_setting_id,omitempty"`
        GatewayHandle           *string          `json:"gateway_handle,omitempty"`
        Disabled                *bool            `json:"disabled,omitempty"`
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
    c.BankName = temp.BankName
    c.MaskedBankRoutingNumber = temp.MaskedBankRoutingNumber
    c.MaskedBankAccountNumber = temp.MaskedBankAccountNumber
    c.BankAccountType = temp.BankAccountType
    c.BankAccountHolderType = temp.BankAccountHolderType
    c.Verified = temp.Verified
    c.SiteGatewaySettingId = temp.SiteGatewaySettingId
    c.GatewayHandle = temp.GatewayHandle
    c.Disabled = temp.Disabled
    return nil
}
