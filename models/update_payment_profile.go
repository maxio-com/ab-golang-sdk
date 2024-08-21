/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// UpdatePaymentProfile represents a UpdatePaymentProfile struct.
type UpdatePaymentProfile struct {
    // The first name of the card holder.
    FirstName            *string          `json:"first_name,omitempty"`
    // The last name of the card holder.
    LastName             *string          `json:"last_name,omitempty"`
    // The full credit card number
    FullNumber           *string          `json:"full_number,omitempty"`
    // The type of card used.
    CardType             *CardType        `json:"card_type,omitempty"`
    // (Optional when performing an Import via vault_token, required otherwise) The 1- or 2-digit credit card expiration month, as an integer or string, i.e. 5
    ExpirationMonth      *string          `json:"expiration_month,omitempty"`
    // (Optional when performing a Import via vault_token, required otherwise) The 4-digit credit card expiration year, as an integer or string, i.e. 2012
    ExpirationYear       *string          `json:"expiration_year,omitempty"`
    // The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing.
    CurrentVault         *AllVaults       `json:"current_vault,omitempty"`
    // The credit card or bank account billing street address (i.e. 123 Main St.). This value is merely passed through to the payment gateway.
    BillingAddress       *string          `json:"billing_address,omitempty"`
    // The credit card or bank account billing address city (i.e. “Boston”). This value is merely passed through to the payment gateway.
    BillingCity          *string          `json:"billing_city,omitempty"`
    // The credit card or bank account billing address state (i.e. MA). This value is merely passed through to the payment gateway. This must conform to the [ISO_3166-1](https://en.wikipedia.org/wiki/ISO_3166-1#Current_codes) in order to be valid for tax locale purposes.
    BillingState         *string          `json:"billing_state,omitempty"`
    // The credit card or bank account billing address zip code (i.e. 12345). This value is merely passed through to the payment gateway.
    BillingZip           *string          `json:"billing_zip,omitempty"`
    // The credit card or bank account billing address country, required in [ISO_3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format (i.e. “US”). This value is merely passed through to the payment gateway. Some gateways require country codes in a specific format. Please check your gateway’s documentation. If creating an ACH subscription, only US is supported at this time.
    BillingCountry       *string          `json:"billing_country,omitempty"`
    // Second line of the customer’s billing address i.e. Apt. 100
    BillingAddress2      Optional[string] `json:"billing_address_2"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdatePaymentProfile.
// It customizes the JSON marshaling process for UpdatePaymentProfile objects.
func (u UpdatePaymentProfile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdatePaymentProfile object to a map representation for JSON marshaling.
func (u UpdatePaymentProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.FirstName != nil {
        structMap["first_name"] = u.FirstName
    }
    if u.LastName != nil {
        structMap["last_name"] = u.LastName
    }
    if u.FullNumber != nil {
        structMap["full_number"] = u.FullNumber
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
    if u.CurrentVault != nil {
        structMap["current_vault"] = u.CurrentVault
    }
    if u.BillingAddress != nil {
        structMap["billing_address"] = u.BillingAddress
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
    if u.BillingAddress2.IsValueSet() {
        if u.BillingAddress2.Value() != nil {
            structMap["billing_address_2"] = u.BillingAddress2.Value()
        } else {
            structMap["billing_address_2"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdatePaymentProfile.
// It customizes the JSON unmarshaling process for UpdatePaymentProfile objects.
func (u *UpdatePaymentProfile) UnmarshalJSON(input []byte) error {
    var temp tempUpdatePaymentProfile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "first_name", "last_name", "full_number", "card_type", "expiration_month", "expiration_year", "current_vault", "billing_address", "billing_city", "billing_state", "billing_zip", "billing_country", "billing_address_2")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.FirstName = temp.FirstName
    u.LastName = temp.LastName
    u.FullNumber = temp.FullNumber
    u.CardType = temp.CardType
    u.ExpirationMonth = temp.ExpirationMonth
    u.ExpirationYear = temp.ExpirationYear
    u.CurrentVault = temp.CurrentVault
    u.BillingAddress = temp.BillingAddress
    u.BillingCity = temp.BillingCity
    u.BillingState = temp.BillingState
    u.BillingZip = temp.BillingZip
    u.BillingCountry = temp.BillingCountry
    u.BillingAddress2 = temp.BillingAddress2
    return nil
}

// tempUpdatePaymentProfile is a temporary struct used for validating the fields of UpdatePaymentProfile.
type tempUpdatePaymentProfile  struct {
    FirstName       *string          `json:"first_name,omitempty"`
    LastName        *string          `json:"last_name,omitempty"`
    FullNumber      *string          `json:"full_number,omitempty"`
    CardType        *CardType        `json:"card_type,omitempty"`
    ExpirationMonth *string          `json:"expiration_month,omitempty"`
    ExpirationYear  *string          `json:"expiration_year,omitempty"`
    CurrentVault    *AllVaults       `json:"current_vault,omitempty"`
    BillingAddress  *string          `json:"billing_address,omitempty"`
    BillingCity     *string          `json:"billing_city,omitempty"`
    BillingState    *string          `json:"billing_state,omitempty"`
    BillingZip      *string          `json:"billing_zip,omitempty"`
    BillingCountry  *string          `json:"billing_country,omitempty"`
    BillingAddress2 Optional[string] `json:"billing_address_2"`
}
