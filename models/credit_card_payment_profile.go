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

// CreditCardPaymentProfile represents a CreditCardPaymentProfile struct.
type CreditCardPaymentProfile struct {
    // The Chargify-assigned ID of the stored card. This value can be used as an input to payment_profile_id when creating a subscription, in order to re-use a stored payment profile for the same customer.
    Id                   *int                   `json:"id,omitempty"`
    // The first name of the card holder.
    FirstName            *string                `json:"first_name,omitempty"`
    // The last name of the card holder.
    LastName             *string                `json:"last_name,omitempty"`
    // A string representation of the credit card number with all but the last 4 digits masked with X’s (i.e. ‘XXXX-XXXX-XXXX-1234’).
    MaskedCardNumber     *string                `json:"masked_card_number,omitempty"`
    // The type of card used.
    CardType             *CardType              `json:"card_type,omitempty"`
    // An integer representing the expiration month of the card(1 – 12).
    ExpirationMonth      *int                   `json:"expiration_month,omitempty"`
    // An integer representing the 4-digit expiration year of the card(i.e. ‘2012’).
    ExpirationYear       *int                   `json:"expiration_year,omitempty"`
    // The Chargify-assigned id for the customer record to which the card belongs.
    CustomerId           *int                   `json:"customer_id,omitempty"`
    // The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing.
    CurrentVault         *CreditCardVault       `json:"current_vault,omitempty"`
    // The “token” provided by your vault storage for an already stored payment profile.
    VaultToken           Optional[string]       `json:"vault_token"`
    // The current billing street address for the card.
    BillingAddress       Optional[string]       `json:"billing_address"`
    // The current billing address city for the card.
    BillingCity          Optional[string]       `json:"billing_city"`
    // The current billing address state for the card.
    BillingState         Optional[string]       `json:"billing_state"`
    // The current billing address zip code for the card.
    BillingZip           Optional[string]       `json:"billing_zip"`
    // The current billing address country for the card.
    BillingCountry       Optional[string]       `json:"billing_country"`
    // (only for Authorize.Net CIM storage): the customerProfileId for the owner of the customerPaymentProfileId provided as the vault_token.
    CustomerVaultToken   Optional[string]       `json:"customer_vault_token"`
    // The current billing street address, second line, for the card.
    BillingAddress2      Optional[string]       `json:"billing_address_2"`
    PaymentType          PaymentType            `json:"payment_type"`
    Disabled             *bool                  `json:"disabled,omitempty"`
    // Token received after sending billing information using chargify.js. This token will only be received if passed as a sole attribute of credit_card_attributes (i.e. tok_9g6hw85pnpt6knmskpwp4ttt)
    ChargifyToken        *string                `json:"chargify_token,omitempty"`
    SiteGatewaySettingId Optional[int]          `json:"site_gateway_setting_id"`
    // An identifier of connected gateway.
    GatewayHandle        Optional[string]       `json:"gateway_handle"`
    // A timestamp indicating when this payment profile was created
    CreatedAt            *time.Time             `json:"created_at,omitempty"`
    // A timestamp indicating when this payment profile was last updated
    UpdatedAt            *time.Time             `json:"updated_at,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreditCardPaymentProfile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreditCardPaymentProfile) String() string {
    return fmt.Sprintf(
    	"CreditCardPaymentProfile[Id=%v, FirstName=%v, LastName=%v, MaskedCardNumber=%v, CardType=%v, ExpirationMonth=%v, ExpirationYear=%v, CustomerId=%v, CurrentVault=%v, VaultToken=%v, BillingAddress=%v, BillingCity=%v, BillingState=%v, BillingZip=%v, BillingCountry=%v, CustomerVaultToken=%v, BillingAddress2=%v, PaymentType=%v, Disabled=%v, ChargifyToken=%v, SiteGatewaySettingId=%v, GatewayHandle=%v, CreatedAt=%v, UpdatedAt=%v, AdditionalProperties=%v]",
    	c.Id, c.FirstName, c.LastName, c.MaskedCardNumber, c.CardType, c.ExpirationMonth, c.ExpirationYear, c.CustomerId, c.CurrentVault, c.VaultToken, c.BillingAddress, c.BillingCity, c.BillingState, c.BillingZip, c.BillingCountry, c.CustomerVaultToken, c.BillingAddress2, c.PaymentType, c.Disabled, c.ChargifyToken, c.SiteGatewaySettingId, c.GatewayHandle, c.CreatedAt, c.UpdatedAt, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreditCardPaymentProfile.
// It customizes the JSON marshaling process for CreditCardPaymentProfile objects.
func (c CreditCardPaymentProfile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "id", "first_name", "last_name", "masked_card_number", "card_type", "expiration_month", "expiration_year", "customer_id", "current_vault", "vault_token", "billing_address", "billing_city", "billing_state", "billing_zip", "billing_country", "customer_vault_token", "billing_address_2", "payment_type", "disabled", "chargify_token", "site_gateway_setting_id", "gateway_handle", "created_at", "updated_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreditCardPaymentProfile object to a map representation for JSON marshaling.
func (c CreditCardPaymentProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.FirstName != nil {
        structMap["first_name"] = c.FirstName
    }
    if c.LastName != nil {
        structMap["last_name"] = c.LastName
    }
    if c.MaskedCardNumber != nil {
        structMap["masked_card_number"] = c.MaskedCardNumber
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
    if c.VaultToken.IsValueSet() {
        if c.VaultToken.Value() != nil {
            structMap["vault_token"] = c.VaultToken.Value()
        } else {
            structMap["vault_token"] = nil
        }
    }
    if c.BillingAddress.IsValueSet() {
        if c.BillingAddress.Value() != nil {
            structMap["billing_address"] = c.BillingAddress.Value()
        } else {
            structMap["billing_address"] = nil
        }
    }
    if c.BillingCity.IsValueSet() {
        if c.BillingCity.Value() != nil {
            structMap["billing_city"] = c.BillingCity.Value()
        } else {
            structMap["billing_city"] = nil
        }
    }
    if c.BillingState.IsValueSet() {
        if c.BillingState.Value() != nil {
            structMap["billing_state"] = c.BillingState.Value()
        } else {
            structMap["billing_state"] = nil
        }
    }
    if c.BillingZip.IsValueSet() {
        if c.BillingZip.Value() != nil {
            structMap["billing_zip"] = c.BillingZip.Value()
        } else {
            structMap["billing_zip"] = nil
        }
    }
    if c.BillingCountry.IsValueSet() {
        if c.BillingCountry.Value() != nil {
            structMap["billing_country"] = c.BillingCountry.Value()
        } else {
            structMap["billing_country"] = nil
        }
    }
    if c.CustomerVaultToken.IsValueSet() {
        if c.CustomerVaultToken.Value() != nil {
            structMap["customer_vault_token"] = c.CustomerVaultToken.Value()
        } else {
            structMap["customer_vault_token"] = nil
        }
    }
    if c.BillingAddress2.IsValueSet() {
        if c.BillingAddress2.Value() != nil {
            structMap["billing_address_2"] = c.BillingAddress2.Value()
        } else {
            structMap["billing_address_2"] = nil
        }
    }
    structMap["payment_type"] = c.PaymentType
    if c.Disabled != nil {
        structMap["disabled"] = c.Disabled
    }
    if c.ChargifyToken != nil {
        structMap["chargify_token"] = c.ChargifyToken
    }
    if c.SiteGatewaySettingId.IsValueSet() {
        if c.SiteGatewaySettingId.Value() != nil {
            structMap["site_gateway_setting_id"] = c.SiteGatewaySettingId.Value()
        } else {
            structMap["site_gateway_setting_id"] = nil
        }
    }
    if c.GatewayHandle.IsValueSet() {
        if c.GatewayHandle.Value() != nil {
            structMap["gateway_handle"] = c.GatewayHandle.Value()
        } else {
            structMap["gateway_handle"] = nil
        }
    }
    if c.CreatedAt != nil {
        structMap["created_at"] = c.CreatedAt.Format(time.RFC3339)
    }
    if c.UpdatedAt != nil {
        structMap["updated_at"] = c.UpdatedAt.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreditCardPaymentProfile.
// It customizes the JSON unmarshaling process for CreditCardPaymentProfile objects.
func (c *CreditCardPaymentProfile) UnmarshalJSON(input []byte) error {
    var temp tempCreditCardPaymentProfile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "first_name", "last_name", "masked_card_number", "card_type", "expiration_month", "expiration_year", "customer_id", "current_vault", "vault_token", "billing_address", "billing_city", "billing_state", "billing_zip", "billing_country", "customer_vault_token", "billing_address_2", "payment_type", "disabled", "chargify_token", "site_gateway_setting_id", "gateway_handle", "created_at", "updated_at")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
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
    c.PaymentType = *temp.PaymentType
    c.Disabled = temp.Disabled
    c.ChargifyToken = temp.ChargifyToken
    c.SiteGatewaySettingId = temp.SiteGatewaySettingId
    c.GatewayHandle = temp.GatewayHandle
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        c.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        c.UpdatedAt = &UpdatedAtVal
    }
    return nil
}

// tempCreditCardPaymentProfile is a temporary struct used for validating the fields of CreditCardPaymentProfile.
type tempCreditCardPaymentProfile  struct {
    Id                   *int             `json:"id,omitempty"`
    FirstName            *string          `json:"first_name,omitempty"`
    LastName             *string          `json:"last_name,omitempty"`
    MaskedCardNumber     *string          `json:"masked_card_number,omitempty"`
    CardType             *CardType        `json:"card_type,omitempty"`
    ExpirationMonth      *int             `json:"expiration_month,omitempty"`
    ExpirationYear       *int             `json:"expiration_year,omitempty"`
    CustomerId           *int             `json:"customer_id,omitempty"`
    CurrentVault         *CreditCardVault `json:"current_vault,omitempty"`
    VaultToken           Optional[string] `json:"vault_token"`
    BillingAddress       Optional[string] `json:"billing_address"`
    BillingCity          Optional[string] `json:"billing_city"`
    BillingState         Optional[string] `json:"billing_state"`
    BillingZip           Optional[string] `json:"billing_zip"`
    BillingCountry       Optional[string] `json:"billing_country"`
    CustomerVaultToken   Optional[string] `json:"customer_vault_token"`
    BillingAddress2      Optional[string] `json:"billing_address_2"`
    PaymentType          *PaymentType     `json:"payment_type"`
    Disabled             *bool            `json:"disabled,omitempty"`
    ChargifyToken        *string          `json:"chargify_token,omitempty"`
    SiteGatewaySettingId Optional[int]    `json:"site_gateway_setting_id"`
    GatewayHandle        Optional[string] `json:"gateway_handle"`
    CreatedAt            *string          `json:"created_at,omitempty"`
    UpdatedAt            *string          `json:"updated_at,omitempty"`
}

func (c *tempCreditCardPaymentProfile) validate() error {
    var errs []string
    if c.PaymentType == nil {
        errs = append(errs, "required field `payment_type` is missing for type `Credit Card Payment Profile`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
