/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// PaypalPaymentProfile represents a PaypalPaymentProfile struct.
type PaypalPaymentProfile struct {
    // The Chargify-assigned ID of the stored PayPal payment profile.
    Id                   *int                   `json:"id,omitempty"`
    // The first name of the PayPal account holder
    FirstName            *string                `json:"first_name,omitempty"`
    // The last name of the PayPal account holder
    LastName             *string                `json:"last_name,omitempty"`
    // The Chargify-assigned id for the customer record to which the PayPal account belongs
    CustomerId           *int                   `json:"customer_id,omitempty"`
    // The vault that stores the payment profile with the provided vault_token.
    CurrentVault         *PayPalVault           `json:"current_vault,omitempty"`
    // The “token” provided by your vault storage for an already stored payment profile
    VaultToken           *string                `json:"vault_token,omitempty"`
    // The current billing street address for the PayPal account
    BillingAddress       Optional[string]       `json:"billing_address"`
    // The current billing address city for the PayPal account
    BillingCity          Optional[string]       `json:"billing_city"`
    // The current billing address state for the PayPal account
    BillingState         Optional[string]       `json:"billing_state"`
    // The current billing address zip code for the PayPal account
    BillingZip           Optional[string]       `json:"billing_zip"`
    // The current billing address country for the PayPal account
    BillingCountry       Optional[string]       `json:"billing_country"`
    CustomerVaultToken   Optional[string]       `json:"customer_vault_token"`
    // The current billing street address, second line, for the PayPal account
    BillingAddress2      Optional[string]       `json:"billing_address_2"`
    PaymentType          PaymentType            `json:"payment_type"`
    SiteGatewaySettingId Optional[int]          `json:"site_gateway_setting_id"`
    GatewayHandle        Optional[string]       `json:"gateway_handle"`
    PaypalEmail          *string                `json:"paypal_email,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PaypalPaymentProfile.
// It customizes the JSON marshaling process for PaypalPaymentProfile objects.
func (p PaypalPaymentProfile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "id", "first_name", "last_name", "customer_id", "current_vault", "vault_token", "billing_address", "billing_city", "billing_state", "billing_zip", "billing_country", "customer_vault_token", "billing_address_2", "payment_type", "site_gateway_setting_id", "gateway_handle", "paypal_email"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PaypalPaymentProfile object to a map representation for JSON marshaling.
func (p PaypalPaymentProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Id != nil {
        structMap["id"] = p.Id
    }
    if p.FirstName != nil {
        structMap["first_name"] = p.FirstName
    }
    if p.LastName != nil {
        structMap["last_name"] = p.LastName
    }
    if p.CustomerId != nil {
        structMap["customer_id"] = p.CustomerId
    }
    if p.CurrentVault != nil {
        structMap["current_vault"] = p.CurrentVault
    }
    if p.VaultToken != nil {
        structMap["vault_token"] = p.VaultToken
    }
    if p.BillingAddress.IsValueSet() {
        if p.BillingAddress.Value() != nil {
            structMap["billing_address"] = p.BillingAddress.Value()
        } else {
            structMap["billing_address"] = nil
        }
    }
    if p.BillingCity.IsValueSet() {
        if p.BillingCity.Value() != nil {
            structMap["billing_city"] = p.BillingCity.Value()
        } else {
            structMap["billing_city"] = nil
        }
    }
    if p.BillingState.IsValueSet() {
        if p.BillingState.Value() != nil {
            structMap["billing_state"] = p.BillingState.Value()
        } else {
            structMap["billing_state"] = nil
        }
    }
    if p.BillingZip.IsValueSet() {
        if p.BillingZip.Value() != nil {
            structMap["billing_zip"] = p.BillingZip.Value()
        } else {
            structMap["billing_zip"] = nil
        }
    }
    if p.BillingCountry.IsValueSet() {
        if p.BillingCountry.Value() != nil {
            structMap["billing_country"] = p.BillingCountry.Value()
        } else {
            structMap["billing_country"] = nil
        }
    }
    if p.CustomerVaultToken.IsValueSet() {
        if p.CustomerVaultToken.Value() != nil {
            structMap["customer_vault_token"] = p.CustomerVaultToken.Value()
        } else {
            structMap["customer_vault_token"] = nil
        }
    }
    if p.BillingAddress2.IsValueSet() {
        if p.BillingAddress2.Value() != nil {
            structMap["billing_address_2"] = p.BillingAddress2.Value()
        } else {
            structMap["billing_address_2"] = nil
        }
    }
    structMap["payment_type"] = p.PaymentType
    if p.SiteGatewaySettingId.IsValueSet() {
        if p.SiteGatewaySettingId.Value() != nil {
            structMap["site_gateway_setting_id"] = p.SiteGatewaySettingId.Value()
        } else {
            structMap["site_gateway_setting_id"] = nil
        }
    }
    if p.GatewayHandle.IsValueSet() {
        if p.GatewayHandle.Value() != nil {
            structMap["gateway_handle"] = p.GatewayHandle.Value()
        } else {
            structMap["gateway_handle"] = nil
        }
    }
    if p.PaypalEmail != nil {
        structMap["paypal_email"] = p.PaypalEmail
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaypalPaymentProfile.
// It customizes the JSON unmarshaling process for PaypalPaymentProfile objects.
func (p *PaypalPaymentProfile) UnmarshalJSON(input []byte) error {
    var temp tempPaypalPaymentProfile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "first_name", "last_name", "customer_id", "current_vault", "vault_token", "billing_address", "billing_city", "billing_state", "billing_zip", "billing_country", "customer_vault_token", "billing_address_2", "payment_type", "site_gateway_setting_id", "gateway_handle", "paypal_email")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Id = temp.Id
    p.FirstName = temp.FirstName
    p.LastName = temp.LastName
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
    p.PaymentType = *temp.PaymentType
    p.SiteGatewaySettingId = temp.SiteGatewaySettingId
    p.GatewayHandle = temp.GatewayHandle
    p.PaypalEmail = temp.PaypalEmail
    return nil
}

// tempPaypalPaymentProfile is a temporary struct used for validating the fields of PaypalPaymentProfile.
type tempPaypalPaymentProfile  struct {
    Id                   *int             `json:"id,omitempty"`
    FirstName            *string          `json:"first_name,omitempty"`
    LastName             *string          `json:"last_name,omitempty"`
    CustomerId           *int             `json:"customer_id,omitempty"`
    CurrentVault         *PayPalVault     `json:"current_vault,omitempty"`
    VaultToken           *string          `json:"vault_token,omitempty"`
    BillingAddress       Optional[string] `json:"billing_address"`
    BillingCity          Optional[string] `json:"billing_city"`
    BillingState         Optional[string] `json:"billing_state"`
    BillingZip           Optional[string] `json:"billing_zip"`
    BillingCountry       Optional[string] `json:"billing_country"`
    CustomerVaultToken   Optional[string] `json:"customer_vault_token"`
    BillingAddress2      Optional[string] `json:"billing_address_2"`
    PaymentType          *PaymentType     `json:"payment_type"`
    SiteGatewaySettingId Optional[int]    `json:"site_gateway_setting_id"`
    GatewayHandle        Optional[string] `json:"gateway_handle"`
    PaypalEmail          *string          `json:"paypal_email,omitempty"`
}

func (p *tempPaypalPaymentProfile) validate() error {
    var errs []string
    if p.PaymentType == nil {
        errs = append(errs, "required field `payment_type` is missing for type `Paypal Payment Profile`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
