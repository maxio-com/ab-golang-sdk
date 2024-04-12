package models

import (
    "encoding/json"
)

// PaymentProfileAttributes represents a PaymentProfileAttributes struct.
// alias to credit_card_attributes
type PaymentProfileAttributes struct {
    // (Optional) Token received after sending billing informations using chargify.js. This token must be passed as a sole attribute of `payment_profile_attributes` (i.e. tok_9g6hw85pnpt6knmskpwp4ttt)
    ChargifyToken        *string                                  `json:"chargify_token,omitempty"`
    Id                   *int                                     `json:"id,omitempty"`
    PaymentType          *string                                  `json:"payment_type,omitempty"`
    // (Optional) First name on card or bank account. If omitted, the first_name from customer attributes will be used.
    FirstName            *string                                  `json:"first_name,omitempty"`
    // (Optional) Last name on card or bank account. If omitted, the last_name from customer attributes will be used.
    LastName             *string                                  `json:"last_name,omitempty"`
    MaskedCardNumber     *string                                  `json:"masked_card_number,omitempty"`
    // The full credit card number (string representation, i.e. 5424000000000015)
    FullNumber           *string                                  `json:"full_number,omitempty"`
    // (Optional, used only for Subscription Import) If you know the card type (i.e. Visa, MC, etc) you may supply it here so that we may display the card type in the UI.
    CardType             *CardType                                `json:"card_type,omitempty"`
    // (Optional when performing a Subscription Import via vault_token, required otherwise) The 1- or 2-digit credit card expiration month, as an integer or string, i.e. 5
    ExpirationMonth      *PaymentProfileAttributesExpirationMonth `json:"expiration_month,omitempty"`
    // (Optional when performing a Subscription Import via vault_token, required otherwise) The 4-digit credit card expiration year, as an integer or string, i.e. 2012
    ExpirationYear       *PaymentProfileAttributesExpirationYear  `json:"expiration_year,omitempty"`
    // (Optional, may be required by your product configuration or gateway settings) The credit card or bank account billing street address (i.e. 123 Main St.). This value is merely passed through to the payment gateway.
    BillingAddress       *string                                  `json:"billing_address,omitempty"`
    // (Optional) Second line of the customer’s billing address i.e. Apt. 100
    BillingAddress2      Optional[string]                         `json:"billing_address_2"`
    // (Optional, may be required by your product configuration or gateway settings) The credit card or bank account billing address city (i.e. “Boston”). This value is merely passed through to the payment gateway.
    BillingCity          *string                                  `json:"billing_city,omitempty"`
    // (Optional, may be required by your product configuration or gateway settings) The credit card or bank account billing address state (i.e. MA). This value is merely passed through to the payment gateway. This must conform to the [ISO_3166-1](https://en.wikipedia.org/wiki/ISO_3166-1#Current_codes) in order to be valid for tax locale purposes.
    BillingState         *string                                  `json:"billing_state,omitempty"`
    // (Optional, may be required by your product configuration or gateway settings) The credit card or bank account billing address country, required in [ISO_3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format (i.e. “US”). This value is merely passed through to the payment gateway. Some gateways require country codes in a specific format. Please check your gateway’s documentation. If creating an ACH subscription, only US is supported at this time.
    BillingCountry       *string                                  `json:"billing_country,omitempty"`
    // (Optional, may be required by your product configuration or gateway settings) The credit card or bank account billing address zip code (i.e. 12345). This value is merely passed through to the payment gateway.
    BillingZip           *string                                  `json:"billing_zip,omitempty"`
    // (Optional, used only for Subscription Import) The vault that stores the payment profile with the provided vault_token.
    CurrentVault         *CurrentVault                            `json:"current_vault,omitempty"`
    // (Optional, used only for Subscription Import) The “token” provided by your vault storage for an already stored payment profile
    VaultToken           *string                                  `json:"vault_token,omitempty"`
    // (Optional, used only for Subscription Import) (only for Authorize.Net CIM storage or Square) The customerProfileId for the owner of the customerPaymentProfileId provided as the vault_token
    CustomerVaultToken   *string                                  `json:"customer_vault_token,omitempty"`
    CustomerId           *int                                     `json:"customer_id,omitempty"`
    PaypalEmail          *string                                  `json:"paypal_email,omitempty"`
    // (Required for Square unless importing with vault_token and customer_vault_token) The nonce generated by the Square Javascript library (SqPaymentForm)
    PaymentMethodNonce   *string                                  `json:"payment_method_nonce,omitempty"`
    // (Optional) This attribute is only available if MultiGateway feature is enabled for your Site. This feature is in the Private Beta currently. gateway_handle is used to directly select a gateway where a payment profile will be stored in. Every connected gateway must have a unique gateway handle specified. Read [Multigateway description](https://chargify.zendesk.com/hc/en-us/articles/4407761759643#connecting-with-multiple-gateways) to learn more about new concepts that MultiGateway introduces and the default behavior when this attribute is not passed.
    GatewayHandle        *string                                  `json:"gateway_handle,omitempty"`
    // (Optional, may be required by your gateway settings) The 3- or 4-digit Card Verification Value. This value is merely passed through to the payment gateway.
    Cvv                  *string                                  `json:"cvv,omitempty"`
    // (Optional, used only for Subscription Import) If you have the last 4 digits of the credit card number, you may supply them here so that we may create a masked card number (i.e. XXXX-XXXX-XXXX-1234) for display in the UI. Last 4 digits are required for refunds in Auth.Net.
    LastFour             *string                                  `json:"last_four,omitempty"`
    AdditionalProperties map[string]any                           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentProfileAttributes.
// It customizes the JSON marshaling process for PaymentProfileAttributes objects.
func (p PaymentProfileAttributes) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentProfileAttributes object to a map representation for JSON marshaling.
func (p PaymentProfileAttributes) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.ChargifyToken != nil {
        structMap["chargify_token"] = p.ChargifyToken
    }
    if p.Id != nil {
        structMap["id"] = p.Id
    }
    if p.PaymentType != nil {
        structMap["payment_type"] = p.PaymentType
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
    if p.FullNumber != nil {
        structMap["full_number"] = p.FullNumber
    }
    if p.CardType != nil {
        structMap["card_type"] = p.CardType
    }
    if p.ExpirationMonth != nil {
        structMap["expiration_month"] = p.ExpirationMonth.toMap()
    }
    if p.ExpirationYear != nil {
        structMap["expiration_year"] = p.ExpirationYear.toMap()
    }
    if p.BillingAddress != nil {
        structMap["billing_address"] = p.BillingAddress
    }
    if p.BillingAddress2.IsValueSet() {
        if p.BillingAddress2.Value() != nil {
            structMap["billing_address_2"] = p.BillingAddress2.Value()
        } else {
            structMap["billing_address_2"] = nil
        }
    }
    if p.BillingCity != nil {
        structMap["billing_city"] = p.BillingCity
    }
    if p.BillingState != nil {
        structMap["billing_state"] = p.BillingState
    }
    if p.BillingCountry != nil {
        structMap["billing_country"] = p.BillingCountry
    }
    if p.BillingZip != nil {
        structMap["billing_zip"] = p.BillingZip
    }
    if p.CurrentVault != nil {
        structMap["current_vault"] = p.CurrentVault
    }
    if p.VaultToken != nil {
        structMap["vault_token"] = p.VaultToken
    }
    if p.CustomerVaultToken != nil {
        structMap["customer_vault_token"] = p.CustomerVaultToken
    }
    if p.CustomerId != nil {
        structMap["customer_id"] = p.CustomerId
    }
    if p.PaypalEmail != nil {
        structMap["paypal_email"] = p.PaypalEmail
    }
    if p.PaymentMethodNonce != nil {
        structMap["payment_method_nonce"] = p.PaymentMethodNonce
    }
    if p.GatewayHandle != nil {
        structMap["gateway_handle"] = p.GatewayHandle
    }
    if p.Cvv != nil {
        structMap["cvv"] = p.Cvv
    }
    if p.LastFour != nil {
        structMap["last_four"] = p.LastFour
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentProfileAttributes.
// It customizes the JSON unmarshaling process for PaymentProfileAttributes objects.
func (p *PaymentProfileAttributes) UnmarshalJSON(input []byte) error {
    var temp paymentProfileAttributes
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "chargify_token", "id", "payment_type", "first_name", "last_name", "masked_card_number", "full_number", "card_type", "expiration_month", "expiration_year", "billing_address", "billing_address_2", "billing_city", "billing_state", "billing_country", "billing_zip", "current_vault", "vault_token", "customer_vault_token", "customer_id", "paypal_email", "payment_method_nonce", "gateway_handle", "cvv", "last_four")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.ChargifyToken = temp.ChargifyToken
    p.Id = temp.Id
    p.PaymentType = temp.PaymentType
    p.FirstName = temp.FirstName
    p.LastName = temp.LastName
    p.MaskedCardNumber = temp.MaskedCardNumber
    p.FullNumber = temp.FullNumber
    p.CardType = temp.CardType
    p.ExpirationMonth = temp.ExpirationMonth
    p.ExpirationYear = temp.ExpirationYear
    p.BillingAddress = temp.BillingAddress
    p.BillingAddress2 = temp.BillingAddress2
    p.BillingCity = temp.BillingCity
    p.BillingState = temp.BillingState
    p.BillingCountry = temp.BillingCountry
    p.BillingZip = temp.BillingZip
    p.CurrentVault = temp.CurrentVault
    p.VaultToken = temp.VaultToken
    p.CustomerVaultToken = temp.CustomerVaultToken
    p.CustomerId = temp.CustomerId
    p.PaypalEmail = temp.PaypalEmail
    p.PaymentMethodNonce = temp.PaymentMethodNonce
    p.GatewayHandle = temp.GatewayHandle
    p.Cvv = temp.Cvv
    p.LastFour = temp.LastFour
    return nil
}

// TODO
type paymentProfileAttributes  struct {
    ChargifyToken      *string                                  `json:"chargify_token,omitempty"`
    Id                 *int                                     `json:"id,omitempty"`
    PaymentType        *string                                  `json:"payment_type,omitempty"`
    FirstName          *string                                  `json:"first_name,omitempty"`
    LastName           *string                                  `json:"last_name,omitempty"`
    MaskedCardNumber   *string                                  `json:"masked_card_number,omitempty"`
    FullNumber         *string                                  `json:"full_number,omitempty"`
    CardType           *CardType                                `json:"card_type,omitempty"`
    ExpirationMonth    *PaymentProfileAttributesExpirationMonth `json:"expiration_month,omitempty"`
    ExpirationYear     *PaymentProfileAttributesExpirationYear  `json:"expiration_year,omitempty"`
    BillingAddress     *string                                  `json:"billing_address,omitempty"`
    BillingAddress2    Optional[string]                         `json:"billing_address_2"`
    BillingCity        *string                                  `json:"billing_city,omitempty"`
    BillingState       *string                                  `json:"billing_state,omitempty"`
    BillingCountry     *string                                  `json:"billing_country,omitempty"`
    BillingZip         *string                                  `json:"billing_zip,omitempty"`
    CurrentVault       *CurrentVault                            `json:"current_vault,omitempty"`
    VaultToken         *string                                  `json:"vault_token,omitempty"`
    CustomerVaultToken *string                                  `json:"customer_vault_token,omitempty"`
    CustomerId         *int                                     `json:"customer_id,omitempty"`
    PaypalEmail        *string                                  `json:"paypal_email,omitempty"`
    PaymentMethodNonce *string                                  `json:"payment_method_nonce,omitempty"`
    GatewayHandle      *string                                  `json:"gateway_handle,omitempty"`
    Cvv                *string                                  `json:"cvv,omitempty"`
    LastFour           *string                                  `json:"last_four,omitempty"`
}
