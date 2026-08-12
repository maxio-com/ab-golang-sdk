// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// GetOneTimeTokenBankAccountPaymentProfile represents a GetOneTimeTokenBankAccountPaymentProfile struct.
type GetOneTimeTokenBankAccountPaymentProfile struct {
    Id                      Optional[string]       `json:"id"`
    FirstName               string                 `json:"first_name"`
    LastName                string                 `json:"last_name"`
    CustomerId              Optional[string]       `json:"customer_id"`
    // The vault that stores the payment profile with the provided vault_token. Use `bogus` for testing.
    CurrentVault            BankAccountVault       `json:"current_vault"`
    VaultToken              string                 `json:"vault_token"`
    BillingAddress          string                 `json:"billing_address"`
    BillingAddress2         *string                `json:"billing_address_2,omitempty"`
    BillingCity             string                 `json:"billing_city"`
    BillingCountry          string                 `json:"billing_country"`
    BillingState            string                 `json:"billing_state"`
    BillingZip              string                 `json:"billing_zip"`
    BankName                string                 `json:"bank_name"`
    MaskedBankRoutingNumber string                 `json:"masked_bank_routing_number"`
    MaskedBankAccountNumber string                 `json:"masked_bank_account_number"`
    // Defaults to checking
    BankAccountType         BankAccountType        `json:"bank_account_type"`
    // Defaults to personal
    BankAccountHolderType   BankAccountHolderType  `json:"bank_account_holder_type"`
    PaymentType             string                 `json:"payment_type"`
    Disabled                bool                   `json:"disabled"`
    SiteGatewaySettingId    int                    `json:"site_gateway_setting_id"`
    CustomerVaultToken      Optional[string]       `json:"customer_vault_token"`
    GatewayHandle           Optional[string]       `json:"gateway_handle"`
    Verified                Optional[bool]         `json:"verified"`
    AdditionalProperties    map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GetOneTimeTokenBankAccountPaymentProfile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GetOneTimeTokenBankAccountPaymentProfile) String() string {
    return fmt.Sprintf(
    	"GetOneTimeTokenBankAccountPaymentProfile[Id=%v, FirstName=%v, LastName=%v, CustomerId=%v, CurrentVault=%v, VaultToken=%v, BillingAddress=%v, BillingAddress2=%v, BillingCity=%v, BillingCountry=%v, BillingState=%v, BillingZip=%v, BankName=%v, MaskedBankRoutingNumber=%v, MaskedBankAccountNumber=%v, BankAccountType=%v, BankAccountHolderType=%v, PaymentType=%v, Disabled=%v, SiteGatewaySettingId=%v, CustomerVaultToken=%v, GatewayHandle=%v, Verified=%v, AdditionalProperties=%v]",
    	g.Id, g.FirstName, g.LastName, g.CustomerId, g.CurrentVault, g.VaultToken, g.BillingAddress, g.BillingAddress2, g.BillingCity, g.BillingCountry, g.BillingState, g.BillingZip, g.BankName, g.MaskedBankRoutingNumber, g.MaskedBankAccountNumber, g.BankAccountType, g.BankAccountHolderType, g.PaymentType, g.Disabled, g.SiteGatewaySettingId, g.CustomerVaultToken, g.GatewayHandle, g.Verified, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GetOneTimeTokenBankAccountPaymentProfile.
// It customizes the JSON marshaling process for GetOneTimeTokenBankAccountPaymentProfile objects.
func (g GetOneTimeTokenBankAccountPaymentProfile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "id", "first_name", "last_name", "customer_id", "current_vault", "vault_token", "billing_address", "billing_address_2", "billing_city", "billing_country", "billing_state", "billing_zip", "bank_name", "masked_bank_routing_number", "masked_bank_account_number", "bank_account_type", "bank_account_holder_type", "payment_type", "disabled", "site_gateway_setting_id", "customer_vault_token", "gateway_handle", "verified"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GetOneTimeTokenBankAccountPaymentProfile object to a map representation for JSON marshaling.
func (g GetOneTimeTokenBankAccountPaymentProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Id.IsValueSet() {
        if g.Id.Value() != nil {
            structMap["id"] = g.Id.Value()
        } else {
            structMap["id"] = nil
        }
    }
    structMap["first_name"] = g.FirstName
    structMap["last_name"] = g.LastName
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
    structMap["bank_name"] = g.BankName
    structMap["masked_bank_routing_number"] = g.MaskedBankRoutingNumber
    structMap["masked_bank_account_number"] = g.MaskedBankAccountNumber
    structMap["bank_account_type"] = g.BankAccountType
    structMap["bank_account_holder_type"] = g.BankAccountHolderType
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
    if g.Verified.IsValueSet() {
        if g.Verified.Value() != nil {
            structMap["verified"] = g.Verified.Value()
        } else {
            structMap["verified"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetOneTimeTokenBankAccountPaymentProfile.
// It customizes the JSON unmarshaling process for GetOneTimeTokenBankAccountPaymentProfile objects.
func (g *GetOneTimeTokenBankAccountPaymentProfile) UnmarshalJSON(input []byte) error {
    var temp tempGetOneTimeTokenBankAccountPaymentProfile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "first_name", "last_name", "customer_id", "current_vault", "vault_token", "billing_address", "billing_address_2", "billing_city", "billing_country", "billing_state", "billing_zip", "bank_name", "masked_bank_routing_number", "masked_bank_account_number", "bank_account_type", "bank_account_holder_type", "payment_type", "disabled", "site_gateway_setting_id", "customer_vault_token", "gateway_handle", "verified")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.Id = temp.Id
    g.FirstName = *temp.FirstName
    g.LastName = *temp.LastName
    g.CustomerId = temp.CustomerId
    g.CurrentVault = *temp.CurrentVault
    g.VaultToken = *temp.VaultToken
    g.BillingAddress = *temp.BillingAddress
    g.BillingAddress2 = temp.BillingAddress2
    g.BillingCity = *temp.BillingCity
    g.BillingCountry = *temp.BillingCountry
    g.BillingState = *temp.BillingState
    g.BillingZip = *temp.BillingZip
    g.BankName = *temp.BankName
    g.MaskedBankRoutingNumber = *temp.MaskedBankRoutingNumber
    g.MaskedBankAccountNumber = *temp.MaskedBankAccountNumber
    g.BankAccountType = *temp.BankAccountType
    g.BankAccountHolderType = *temp.BankAccountHolderType
    g.PaymentType = *temp.PaymentType
    g.Disabled = *temp.Disabled
    g.SiteGatewaySettingId = *temp.SiteGatewaySettingId
    g.CustomerVaultToken = temp.CustomerVaultToken
    g.GatewayHandle = temp.GatewayHandle
    g.Verified = temp.Verified
    return nil
}

// tempGetOneTimeTokenBankAccountPaymentProfile is a temporary struct used for validating the fields of GetOneTimeTokenBankAccountPaymentProfile.
type tempGetOneTimeTokenBankAccountPaymentProfile  struct {
    Id                      Optional[string]       `json:"id"`
    FirstName               *string                `json:"first_name"`
    LastName                *string                `json:"last_name"`
    CustomerId              Optional[string]       `json:"customer_id"`
    CurrentVault            *BankAccountVault      `json:"current_vault"`
    VaultToken              *string                `json:"vault_token"`
    BillingAddress          *string                `json:"billing_address"`
    BillingAddress2         *string                `json:"billing_address_2,omitempty"`
    BillingCity             *string                `json:"billing_city"`
    BillingCountry          *string                `json:"billing_country"`
    BillingState            *string                `json:"billing_state"`
    BillingZip              *string                `json:"billing_zip"`
    BankName                *string                `json:"bank_name"`
    MaskedBankRoutingNumber *string                `json:"masked_bank_routing_number"`
    MaskedBankAccountNumber *string                `json:"masked_bank_account_number"`
    BankAccountType         *BankAccountType       `json:"bank_account_type"`
    BankAccountHolderType   *BankAccountHolderType `json:"bank_account_holder_type"`
    PaymentType             *string                `json:"payment_type"`
    Disabled                *bool                  `json:"disabled"`
    SiteGatewaySettingId    *int                   `json:"site_gateway_setting_id"`
    CustomerVaultToken      Optional[string]       `json:"customer_vault_token"`
    GatewayHandle           Optional[string]       `json:"gateway_handle"`
    Verified                Optional[bool]         `json:"verified"`
}

func (g *tempGetOneTimeTokenBankAccountPaymentProfile) validate() error {
    var errs []string
    if g.FirstName == nil {
        errs = append(errs, "required field `first_name` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.LastName == nil {
        errs = append(errs, "required field `last_name` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.CurrentVault == nil {
        errs = append(errs, "required field `current_vault` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.VaultToken == nil {
        errs = append(errs, "required field `vault_token` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.BillingAddress == nil {
        errs = append(errs, "required field `billing_address` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.BillingCity == nil {
        errs = append(errs, "required field `billing_city` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.BillingCountry == nil {
        errs = append(errs, "required field `billing_country` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.BillingState == nil {
        errs = append(errs, "required field `billing_state` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.BillingZip == nil {
        errs = append(errs, "required field `billing_zip` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.BankName == nil {
        errs = append(errs, "required field `bank_name` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.MaskedBankRoutingNumber == nil {
        errs = append(errs, "required field `masked_bank_routing_number` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.MaskedBankAccountNumber == nil {
        errs = append(errs, "required field `masked_bank_account_number` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.BankAccountType == nil {
        errs = append(errs, "required field `bank_account_type` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.BankAccountHolderType == nil {
        errs = append(errs, "required field `bank_account_holder_type` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.PaymentType == nil {
        errs = append(errs, "required field `payment_type` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.Disabled == nil {
        errs = append(errs, "required field `disabled` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if g.SiteGatewaySettingId == nil {
        errs = append(errs, "required field `site_gateway_setting_id` is missing for type `Get One Time Token Bank Account Payment Profile`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
