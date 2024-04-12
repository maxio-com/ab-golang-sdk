package models

import (
    "encoding/json"
    "errors"
    "strings"
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
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GetOneTimeTokenPaymentProfile.
// It customizes the JSON marshaling process for GetOneTimeTokenPaymentProfile objects.
func (g GetOneTimeTokenPaymentProfile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetOneTimeTokenPaymentProfile object to a map representation for JSON marshaling.
func (g GetOneTimeTokenPaymentProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Id.IsValueSet() {
        if g.Id.Value() != nil {
            structMap["id"] = g.Id.Value()
        } else {
            structMap["id"] = nil
        }
    }
    structMap["first_name"] = g.FirstName
    structMap["last_name"] = g.LastName
    structMap["masked_card_number"] = g.MaskedCardNumber
    structMap["card_type"] = g.CardType
    structMap["expiration_month"] = g.ExpirationMonth
    structMap["expiration_year"] = g.ExpirationYear
    if g.CustomerId.IsValueSet() {
        if g.CustomerId.Value() != nil {
            structMap["customer_id"] = g.CustomerId.Value()
        } else {
            structMap["customer_id"] = nil
        }
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
        if g.CustomerVaultToken.Value() != nil {
            structMap["customer_vault_token"] = g.CustomerVaultToken.Value()
        } else {
            structMap["customer_vault_token"] = nil
        }
    }
    if g.GatewayHandle.IsValueSet() {
        if g.GatewayHandle.Value() != nil {
            structMap["gateway_handle"] = g.GatewayHandle.Value()
        } else {
            structMap["gateway_handle"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetOneTimeTokenPaymentProfile.
// It customizes the JSON unmarshaling process for GetOneTimeTokenPaymentProfile objects.
func (g *GetOneTimeTokenPaymentProfile) UnmarshalJSON(input []byte) error {
    var temp getOneTimeTokenPaymentProfile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "first_name", "last_name", "masked_card_number", "card_type", "expiration_month", "expiration_year", "customer_id", "current_vault", "vault_token", "billing_address", "billing_address_2", "billing_city", "billing_country", "billing_state", "billing_zip", "payment_type", "disabled", "site_gateway_setting_id", "customer_vault_token", "gateway_handle")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.Id = temp.Id
    g.FirstName = *temp.FirstName
    g.LastName = *temp.LastName
    g.MaskedCardNumber = *temp.MaskedCardNumber
    g.CardType = *temp.CardType
    g.ExpirationMonth = *temp.ExpirationMonth
    g.ExpirationYear = *temp.ExpirationYear
    g.CustomerId = temp.CustomerId
    g.CurrentVault = *temp.CurrentVault
    g.VaultToken = *temp.VaultToken
    g.BillingAddress = *temp.BillingAddress
    g.BillingAddress2 = temp.BillingAddress2
    g.BillingCity = *temp.BillingCity
    g.BillingCountry = *temp.BillingCountry
    g.BillingState = *temp.BillingState
    g.BillingZip = *temp.BillingZip
    g.PaymentType = *temp.PaymentType
    g.Disabled = *temp.Disabled
    g.SiteGatewaySettingId = *temp.SiteGatewaySettingId
    g.CustomerVaultToken = temp.CustomerVaultToken
    g.GatewayHandle = temp.GatewayHandle
    return nil
}

// TODO
type getOneTimeTokenPaymentProfile  struct {
    Id                   Optional[string] `json:"id"`
    FirstName            *string          `json:"first_name"`
    LastName             *string          `json:"last_name"`
    MaskedCardNumber     *string          `json:"masked_card_number"`
    CardType             *CardType        `json:"card_type"`
    ExpirationMonth      *float64         `json:"expiration_month"`
    ExpirationYear       *float64         `json:"expiration_year"`
    CustomerId           Optional[string] `json:"customer_id"`
    CurrentVault         *CurrentVault    `json:"current_vault"`
    VaultToken           *string          `json:"vault_token"`
    BillingAddress       *string          `json:"billing_address"`
    BillingAddress2      *string          `json:"billing_address_2,omitempty"`
    BillingCity          *string          `json:"billing_city"`
    BillingCountry       *string          `json:"billing_country"`
    BillingState         *string          `json:"billing_state"`
    BillingZip           *string          `json:"billing_zip"`
    PaymentType          *string          `json:"payment_type"`
    Disabled             *bool            `json:"disabled"`
    SiteGatewaySettingId *int             `json:"site_gateway_setting_id"`
    CustomerVaultToken   Optional[string] `json:"customer_vault_token"`
    GatewayHandle        Optional[string] `json:"gateway_handle"`
}

func (g *getOneTimeTokenPaymentProfile) validate() error {
    var errs []string
    if g.FirstName == nil {
        errs = append(errs, "required field `first_name` is missing for type `Get One Time Token Payment Profile`")
    }
    if g.LastName == nil {
        errs = append(errs, "required field `last_name` is missing for type `Get One Time Token Payment Profile`")
    }
    if g.MaskedCardNumber == nil {
        errs = append(errs, "required field `masked_card_number` is missing for type `Get One Time Token Payment Profile`")
    }
    if g.CardType == nil {
        errs = append(errs, "required field `card_type` is missing for type `Get One Time Token Payment Profile`")
    }
    if g.ExpirationMonth == nil {
        errs = append(errs, "required field `expiration_month` is missing for type `Get One Time Token Payment Profile`")
    }
    if g.ExpirationYear == nil {
        errs = append(errs, "required field `expiration_year` is missing for type `Get One Time Token Payment Profile`")
    }
    if g.CurrentVault == nil {
        errs = append(errs, "required field `current_vault` is missing for type `Get One Time Token Payment Profile`")
    }
    if g.VaultToken == nil {
        errs = append(errs, "required field `vault_token` is missing for type `Get One Time Token Payment Profile`")
    }
    if g.BillingAddress == nil {
        errs = append(errs, "required field `billing_address` is missing for type `Get One Time Token Payment Profile`")
    }
    if g.BillingCity == nil {
        errs = append(errs, "required field `billing_city` is missing for type `Get One Time Token Payment Profile`")
    }
    if g.BillingCountry == nil {
        errs = append(errs, "required field `billing_country` is missing for type `Get One Time Token Payment Profile`")
    }
    if g.BillingState == nil {
        errs = append(errs, "required field `billing_state` is missing for type `Get One Time Token Payment Profile`")
    }
    if g.BillingZip == nil {
        errs = append(errs, "required field `billing_zip` is missing for type `Get One Time Token Payment Profile`")
    }
    if g.PaymentType == nil {
        errs = append(errs, "required field `payment_type` is missing for type `Get One Time Token Payment Profile`")
    }
    if g.Disabled == nil {
        errs = append(errs, "required field `disabled` is missing for type `Get One Time Token Payment Profile`")
    }
    if g.SiteGatewaySettingId == nil {
        errs = append(errs, "required field `site_gateway_setting_id` is missing for type `Get One Time Token Payment Profile`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
