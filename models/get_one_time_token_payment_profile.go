package models

import (
    "encoding/json"
)

// GetOneTimeTokenPaymentProfile represents a GetOneTimeTokenPaymentProfile struct.
type GetOneTimeTokenPaymentProfile struct {
    Id                   Optional[string] `json:"id"`
    FirstName            string           `json:"first_name"`
    LastName             string           `json:"last_name"`
    MaskedCardNumber     string           `json:"masked_card_number"`
    // The type of card used.
    CardType             CardType         `json:"card_type"`
    ExpirationMonth      float64          `json:"expiration_month"`
    ExpirationYear       float64          `json:"expiration_year"`
    CustomerId           Optional[string] `json:"customer_id"`
    // The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing.
    CurrentVault         CurrentVault     `json:"current_vault"`
    VaultToken           string           `json:"vault_token"`
    BillingAddress       string           `json:"billing_address"`
    BillingAddress2      *string          `json:"billing_address_2,omitempty"`
    BillingCity          string           `json:"billing_city"`
    BillingCountry       string           `json:"billing_country"`
    BillingState         string           `json:"billing_state"`
    BillingZip           string           `json:"billing_zip"`
    PaymentType          string           `json:"payment_type"`
    Disabled             bool             `json:"disabled"`
    SiteGatewaySettingId int              `json:"site_gateway_setting_id"`
    CustomerVaultToken   Optional[string] `json:"customer_vault_token"`
    GatewayHandle        Optional[string] `json:"gateway_handle"`
}

// MarshalJSON implements the json.Marshaler interface for GetOneTimeTokenPaymentProfile.
// It customizes the JSON marshaling process for GetOneTimeTokenPaymentProfile objects.
func (g *GetOneTimeTokenPaymentProfile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetOneTimeTokenPaymentProfile object to a map representation for JSON marshaling.
func (g *GetOneTimeTokenPaymentProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    if g.Id.IsValueSet() {
        structMap["id"] = g.Id.Value()
    }
    structMap["first_name"] = g.FirstName
    structMap["last_name"] = g.LastName
    structMap["masked_card_number"] = g.MaskedCardNumber
    structMap["card_type"] = g.CardType
    structMap["expiration_month"] = g.ExpirationMonth
    structMap["expiration_year"] = g.ExpirationYear
    if g.CustomerId.IsValueSet() {
        structMap["customer_id"] = g.CustomerId.Value()
    }
    structMap["current_vault"] = g.CurrentVault
    structMap["vault_token"] = g.VaultToken
    structMap["billing_address"] = g.BillingAddress
    if g.BillingAddress2 != nil {
        structMap["billing_address_2"] = g.BillingAddress2
    }
    structMap["billing_city"] = g.BillingCity
    structMap["billing_country"] = g.BillingCountry
    structMap["billing_state"] = g.BillingState
    structMap["billing_zip"] = g.BillingZip
    structMap["payment_type"] = g.PaymentType
    structMap["disabled"] = g.Disabled
    structMap["site_gateway_setting_id"] = g.SiteGatewaySettingId
    if g.CustomerVaultToken.IsValueSet() {
        structMap["customer_vault_token"] = g.CustomerVaultToken.Value()
    }
    if g.GatewayHandle.IsValueSet() {
        structMap["gateway_handle"] = g.GatewayHandle.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetOneTimeTokenPaymentProfile.
// It customizes the JSON unmarshaling process for GetOneTimeTokenPaymentProfile objects.
func (g *GetOneTimeTokenPaymentProfile) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                   Optional[string] `json:"id"`
        FirstName            string           `json:"first_name"`
        LastName             string           `json:"last_name"`
        MaskedCardNumber     string           `json:"masked_card_number"`
        CardType             CardType         `json:"card_type"`
        ExpirationMonth      float64          `json:"expiration_month"`
        ExpirationYear       float64          `json:"expiration_year"`
        CustomerId           Optional[string] `json:"customer_id"`
        CurrentVault         CurrentVault     `json:"current_vault"`
        VaultToken           string           `json:"vault_token"`
        BillingAddress       string           `json:"billing_address"`
        BillingAddress2      *string          `json:"billing_address_2,omitempty"`
        BillingCity          string           `json:"billing_city"`
        BillingCountry       string           `json:"billing_country"`
        BillingState         string           `json:"billing_state"`
        BillingZip           string           `json:"billing_zip"`
        PaymentType          string           `json:"payment_type"`
        Disabled             bool             `json:"disabled"`
        SiteGatewaySettingId int              `json:"site_gateway_setting_id"`
        CustomerVaultToken   Optional[string] `json:"customer_vault_token"`
        GatewayHandle        Optional[string] `json:"gateway_handle"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.Id = temp.Id
    g.FirstName = temp.FirstName
    g.LastName = temp.LastName
    g.MaskedCardNumber = temp.MaskedCardNumber
    g.CardType = temp.CardType
    g.ExpirationMonth = temp.ExpirationMonth
    g.ExpirationYear = temp.ExpirationYear
    g.CustomerId = temp.CustomerId
    g.CurrentVault = temp.CurrentVault
    g.VaultToken = temp.VaultToken
    g.BillingAddress = temp.BillingAddress
    g.BillingAddress2 = temp.BillingAddress2
    g.BillingCity = temp.BillingCity
    g.BillingCountry = temp.BillingCountry
    g.BillingState = temp.BillingState
    g.BillingZip = temp.BillingZip
    g.PaymentType = temp.PaymentType
    g.Disabled = temp.Disabled
    g.SiteGatewaySettingId = temp.SiteGatewaySettingId
    g.CustomerVaultToken = temp.CustomerVaultToken
    g.GatewayHandle = temp.GatewayHandle
    return nil
}
