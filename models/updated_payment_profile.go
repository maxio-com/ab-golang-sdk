package models

import (
    "encoding/json"
)

// UpdatedPaymentProfile represents a UpdatedPaymentProfile struct.
type UpdatedPaymentProfile struct {
    Id                   *int             `json:"id,omitempty"`
    FirstName            *string          `json:"first_name,omitempty"`
    LastName             *string          `json:"last_name,omitempty"`
    CardType             *string          `json:"card_type,omitempty"`
    ExpirationMonth      *int             `json:"expiration_month,omitempty"`
    ExpirationYear       *int             `json:"expiration_year,omitempty"`
    CustomerId           *int             `json:"customer_id,omitempty"`
    // The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing.
    CurrentVault         *CurrentVault    `json:"current_vault,omitempty"`
    VaultToken           *string          `json:"vault_token,omitempty"`
    BillingAddress       *string          `json:"billing_address,omitempty"`
    BillingAddress2      *string          `json:"billing_address_2,omitempty"`
    BillingCity          *string          `json:"billing_city,omitempty"`
    BillingState         *string          `json:"billing_state,omitempty"`
    BillingZip           *string          `json:"billing_zip,omitempty"`
    BillingCountry       *string          `json:"billing_country,omitempty"`
    PaymentType          *string          `json:"payment_type,omitempty"`
    SiteGatewaySettingId *int             `json:"site_gateway_setting_id,omitempty"`
    GatewayHandle        Optional[string] `json:"gateway_handle"`
    MaskedCardNumber     *string          `json:"masked_card_number,omitempty"`
    CustomerVaultToken   Optional[string] `json:"customer_vault_token"`
}

// MarshalJSON implements the json.Marshaler interface for UpdatedPaymentProfile.
// It customizes the JSON marshaling process for UpdatedPaymentProfile objects.
func (u *UpdatedPaymentProfile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdatedPaymentProfile object to a map representation for JSON marshaling.
func (u *UpdatedPaymentProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    if u.Id != nil {
        structMap["id"] = u.Id
    }
    if u.FirstName != nil {
        structMap["first_name"] = u.FirstName
    }
    if u.LastName != nil {
        structMap["last_name"] = u.LastName
    }
    if u.CardType != nil {
        structMap["card_type"] = u.CardType
    }
    if u.ExpirationMonth != nil {
        structMap["expiration_month"] = u.ExpirationMonth
    }
    if u.ExpirationYear != nil {
        structMap["expiration_year"] = u.ExpirationYear
    }
    if u.CustomerId != nil {
        structMap["customer_id"] = u.CustomerId
    }
    if u.CurrentVault != nil {
        structMap["current_vault"] = u.CurrentVault
    }
    if u.VaultToken != nil {
        structMap["vault_token"] = u.VaultToken
    }
    if u.BillingAddress != nil {
        structMap["billing_address"] = u.BillingAddress
    }
    if u.BillingAddress2 != nil {
        structMap["billing_address_2"] = u.BillingAddress2
    }
    if u.BillingCity != nil {
        structMap["billing_city"] = u.BillingCity
    }
    if u.BillingState != nil {
        structMap["billing_state"] = u.BillingState
    }
    if u.BillingZip != nil {
        structMap["billing_zip"] = u.BillingZip
    }
    if u.BillingCountry != nil {
        structMap["billing_country"] = u.BillingCountry
    }
    if u.PaymentType != nil {
        structMap["payment_type"] = u.PaymentType
    }
    if u.SiteGatewaySettingId != nil {
        structMap["site_gateway_setting_id"] = u.SiteGatewaySettingId
    }
    if u.GatewayHandle.IsValueSet() {
        structMap["gateway_handle"] = u.GatewayHandle.Value()
    }
    if u.MaskedCardNumber != nil {
        structMap["masked_card_number"] = u.MaskedCardNumber
    }
    if u.CustomerVaultToken.IsValueSet() {
        structMap["customer_vault_token"] = u.CustomerVaultToken.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdatedPaymentProfile.
// It customizes the JSON unmarshaling process for UpdatedPaymentProfile objects.
func (u *UpdatedPaymentProfile) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                   *int             `json:"id,omitempty"`
        FirstName            *string          `json:"first_name,omitempty"`
        LastName             *string          `json:"last_name,omitempty"`
        CardType             *string          `json:"card_type,omitempty"`
        ExpirationMonth      *int             `json:"expiration_month,omitempty"`
        ExpirationYear       *int             `json:"expiration_year,omitempty"`
        CustomerId           *int             `json:"customer_id,omitempty"`
        CurrentVault         *CurrentVault    `json:"current_vault,omitempty"`
        VaultToken           *string          `json:"vault_token,omitempty"`
        BillingAddress       *string          `json:"billing_address,omitempty"`
        BillingAddress2      *string          `json:"billing_address_2,omitempty"`
        BillingCity          *string          `json:"billing_city,omitempty"`
        BillingState         *string          `json:"billing_state,omitempty"`
        BillingZip           *string          `json:"billing_zip,omitempty"`
        BillingCountry       *string          `json:"billing_country,omitempty"`
        PaymentType          *string          `json:"payment_type,omitempty"`
        SiteGatewaySettingId *int             `json:"site_gateway_setting_id,omitempty"`
        GatewayHandle        Optional[string] `json:"gateway_handle"`
        MaskedCardNumber     *string          `json:"masked_card_number,omitempty"`
        CustomerVaultToken   Optional[string] `json:"customer_vault_token"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.Id = temp.Id
    u.FirstName = temp.FirstName
    u.LastName = temp.LastName
    u.CardType = temp.CardType
    u.ExpirationMonth = temp.ExpirationMonth
    u.ExpirationYear = temp.ExpirationYear
    u.CustomerId = temp.CustomerId
    u.CurrentVault = temp.CurrentVault
    u.VaultToken = temp.VaultToken
    u.BillingAddress = temp.BillingAddress
    u.BillingAddress2 = temp.BillingAddress2
    u.BillingCity = temp.BillingCity
    u.BillingState = temp.BillingState
    u.BillingZip = temp.BillingZip
    u.BillingCountry = temp.BillingCountry
    u.PaymentType = temp.PaymentType
    u.SiteGatewaySettingId = temp.SiteGatewaySettingId
    u.GatewayHandle = temp.GatewayHandle
    u.MaskedCardNumber = temp.MaskedCardNumber
    u.CustomerVaultToken = temp.CustomerVaultToken
    return nil
}
