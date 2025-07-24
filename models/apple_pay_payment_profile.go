// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "strings"
    "time"
)

// ApplePayPaymentProfile represents a ApplePayPaymentProfile struct.
type ApplePayPaymentProfile struct {
    // The Chargify-assigned ID of the Apple Pay payment profile.
    Id                   *int                   `json:"id,omitempty"`
    // The first name of the Apple Pay account holder
    FirstName            *string                `json:"first_name,omitempty"`
    // The last name of the Apple Pay account holder
    LastName             *string                `json:"last_name,omitempty"`
    // The Chargify-assigned id for the customer record to which the Apple Pay account belongs
    CustomerId           *int                   `json:"customer_id,omitempty"`
    // The vault that stores the payment profile with the provided vault_token.
    CurrentVault         *ApplePayVault         `json:"current_vault,omitempty"`
    // The “token” provided by your vault storage for an already stored payment profile
    VaultToken           *string                `json:"vault_token,omitempty"`
    // The current billing street address for the Apple Pay account
    BillingAddress       Optional[string]       `json:"billing_address"`
    // The current billing address city for the Apple Pay account
    BillingCity          Optional[string]       `json:"billing_city"`
    // The current billing address state for the Apple Pay account
    BillingState         Optional[string]       `json:"billing_state"`
    // The current billing address zip code for the Apple Pay account
    BillingZip           Optional[string]       `json:"billing_zip"`
    // The current billing address country for the Apple Pay account
    BillingCountry       Optional[string]       `json:"billing_country"`
    CustomerVaultToken   Optional[string]       `json:"customer_vault_token"`
    // The current billing street address, second line, for the Apple Pay account
    BillingAddress2      Optional[string]       `json:"billing_address_2"`
    PaymentType          PaymentType            `json:"payment_type"`
    SiteGatewaySettingId Optional[int]          `json:"site_gateway_setting_id"`
    GatewayHandle        Optional[string]       `json:"gateway_handle"`
    // A timestamp indicating when this payment profile was created
    CreatedAt            *time.Time             `json:"created_at,omitempty"`
    // A timestamp indicating when this payment profile was last updated
    UpdatedAt            *time.Time             `json:"updated_at,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApplePayPaymentProfile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApplePayPaymentProfile) String() string {
    return fmt.Sprintf(
    	"ApplePayPaymentProfile[Id=%v, FirstName=%v, LastName=%v, CustomerId=%v, CurrentVault=%v, VaultToken=%v, BillingAddress=%v, BillingCity=%v, BillingState=%v, BillingZip=%v, BillingCountry=%v, CustomerVaultToken=%v, BillingAddress2=%v, PaymentType=%v, SiteGatewaySettingId=%v, GatewayHandle=%v, CreatedAt=%v, UpdatedAt=%v, AdditionalProperties=%v]",
    	a.Id, a.FirstName, a.LastName, a.CustomerId, a.CurrentVault, a.VaultToken, a.BillingAddress, a.BillingCity, a.BillingState, a.BillingZip, a.BillingCountry, a.CustomerVaultToken, a.BillingAddress2, a.PaymentType, a.SiteGatewaySettingId, a.GatewayHandle, a.CreatedAt, a.UpdatedAt, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApplePayPaymentProfile.
// It customizes the JSON marshaling process for ApplePayPaymentProfile objects.
func (a ApplePayPaymentProfile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "id", "first_name", "last_name", "customer_id", "current_vault", "vault_token", "billing_address", "billing_city", "billing_state", "billing_zip", "billing_country", "customer_vault_token", "billing_address_2", "payment_type", "site_gateway_setting_id", "gateway_handle", "created_at", "updated_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApplePayPaymentProfile object to a map representation for JSON marshaling.
func (a ApplePayPaymentProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Id != nil {
        structMap["id"] = a.Id
    }
    if a.FirstName != nil {
        structMap["first_name"] = a.FirstName
    }
    if a.LastName != nil {
        structMap["last_name"] = a.LastName
    }
    if a.CustomerId != nil {
        structMap["customer_id"] = a.CustomerId
    }
    if a.CurrentVault != nil {
        structMap["current_vault"] = a.CurrentVault
    }
    if a.VaultToken != nil {
        structMap["vault_token"] = a.VaultToken
    }
    if a.BillingAddress.IsValueSet() {
        if a.BillingAddress.Value() != nil {
            structMap["billing_address"] = a.BillingAddress.Value()
        } else {
            structMap["billing_address"] = nil
        }
    }
    if a.BillingCity.IsValueSet() {
        if a.BillingCity.Value() != nil {
            structMap["billing_city"] = a.BillingCity.Value()
        } else {
            structMap["billing_city"] = nil
        }
    }
    if a.BillingState.IsValueSet() {
        if a.BillingState.Value() != nil {
            structMap["billing_state"] = a.BillingState.Value()
        } else {
            structMap["billing_state"] = nil
        }
    }
    if a.BillingZip.IsValueSet() {
        if a.BillingZip.Value() != nil {
            structMap["billing_zip"] = a.BillingZip.Value()
        } else {
            structMap["billing_zip"] = nil
        }
    }
    if a.BillingCountry.IsValueSet() {
        if a.BillingCountry.Value() != nil {
            structMap["billing_country"] = a.BillingCountry.Value()
        } else {
            structMap["billing_country"] = nil
        }
    }
    if a.CustomerVaultToken.IsValueSet() {
        if a.CustomerVaultToken.Value() != nil {
            structMap["customer_vault_token"] = a.CustomerVaultToken.Value()
        } else {
            structMap["customer_vault_token"] = nil
        }
    }
    if a.BillingAddress2.IsValueSet() {
        if a.BillingAddress2.Value() != nil {
            structMap["billing_address_2"] = a.BillingAddress2.Value()
        } else {
            structMap["billing_address_2"] = nil
        }
    }
    structMap["payment_type"] = a.PaymentType
    if a.SiteGatewaySettingId.IsValueSet() {
        if a.SiteGatewaySettingId.Value() != nil {
            structMap["site_gateway_setting_id"] = a.SiteGatewaySettingId.Value()
        } else {
            structMap["site_gateway_setting_id"] = nil
        }
    }
    if a.GatewayHandle.IsValueSet() {
        if a.GatewayHandle.Value() != nil {
            structMap["gateway_handle"] = a.GatewayHandle.Value()
        } else {
            structMap["gateway_handle"] = nil
        }
    }
    if a.CreatedAt != nil {
        structMap["created_at"] = a.CreatedAt.Format(time.RFC3339)
    }
    if a.UpdatedAt != nil {
        structMap["updated_at"] = a.UpdatedAt.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApplePayPaymentProfile.
// It customizes the JSON unmarshaling process for ApplePayPaymentProfile objects.
func (a *ApplePayPaymentProfile) UnmarshalJSON(input []byte) error {
    var temp tempApplePayPaymentProfile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "first_name", "last_name", "customer_id", "current_vault", "vault_token", "billing_address", "billing_city", "billing_state", "billing_zip", "billing_country", "customer_vault_token", "billing_address_2", "payment_type", "site_gateway_setting_id", "gateway_handle", "created_at", "updated_at")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Id = temp.Id
    a.FirstName = temp.FirstName
    a.LastName = temp.LastName
    a.CustomerId = temp.CustomerId
    a.CurrentVault = temp.CurrentVault
    a.VaultToken = temp.VaultToken
    a.BillingAddress = temp.BillingAddress
    a.BillingCity = temp.BillingCity
    a.BillingState = temp.BillingState
    a.BillingZip = temp.BillingZip
    a.BillingCountry = temp.BillingCountry
    a.CustomerVaultToken = temp.CustomerVaultToken
    a.BillingAddress2 = temp.BillingAddress2
    a.PaymentType = *temp.PaymentType
    a.SiteGatewaySettingId = temp.SiteGatewaySettingId
    a.GatewayHandle = temp.GatewayHandle
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        a.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        a.UpdatedAt = &UpdatedAtVal
    }
    return nil
}

// tempApplePayPaymentProfile is a temporary struct used for validating the fields of ApplePayPaymentProfile.
type tempApplePayPaymentProfile  struct {
    Id                   *int             `json:"id,omitempty"`
    FirstName            *string          `json:"first_name,omitempty"`
    LastName             *string          `json:"last_name,omitempty"`
    CustomerId           *int             `json:"customer_id,omitempty"`
    CurrentVault         *ApplePayVault   `json:"current_vault,omitempty"`
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
    CreatedAt            *string          `json:"created_at,omitempty"`
    UpdatedAt            *string          `json:"updated_at,omitempty"`
}

func (a *tempApplePayPaymentProfile) validate() error {
    var errs []string
    if a.PaymentType == nil {
        errs = append(errs, "required field `payment_type` is missing for type `ApplePay Payment Profile`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
