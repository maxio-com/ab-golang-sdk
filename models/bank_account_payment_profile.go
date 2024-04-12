package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// BankAccountPaymentProfile represents a BankAccountPaymentProfile struct.
type BankAccountPaymentProfile struct {
    // The Chargify-assigned ID of the stored bank account. This value can be used as an input to payment_profile_id when creating a subscription, in order to re-use a stored payment profile for the same customer
    Id                      *int                   `json:"id,omitempty"`
    // The first name of the bank account holder
    FirstName               *string                `json:"first_name,omitempty"`
    // The last name of the bank account holder
    LastName                *string                `json:"last_name,omitempty"`
    // The Chargify-assigned id for the customer record to which the bank account belongs
    CustomerId              *int                   `json:"customer_id,omitempty"`
    // The vault that stores the payment profile with the provided vault_token.
    CurrentVault            *BankAccountVault      `json:"current_vault,omitempty"`
    // The “token” provided by your vault storage for an already stored payment profile
    VaultToken              *string                `json:"vault_token,omitempty"`
    // The current billing street address for the bank account
    BillingAddress          Optional[string]       `json:"billing_address"`
    // The current billing address city for the bank account
    BillingCity             Optional[string]       `json:"billing_city"`
    // The current billing address state for the bank account
    BillingState            Optional[string]       `json:"billing_state"`
    // The current billing address zip code for the bank account
    BillingZip              Optional[string]       `json:"billing_zip"`
    // The current billing address country for the bank account
    BillingCountry          Optional[string]       `json:"billing_country"`
    // (only for Authorize.Net CIM storage): the customerProfileId for the owner of the customerPaymentProfileId provided as the vault_token.
    CustomerVaultToken      Optional[string]       `json:"customer_vault_token"`
    // The current billing street address, second line, for the bank account
    BillingAddress2         Optional[string]       `json:"billing_address_2"`
    // The bank where the account resides
    BankName                *string                `json:"bank_name,omitempty"`
    // A string representation of the stored bank routing number with all but the last 4 digits marked with X’s (i.e. ‘XXXXXXX1111’). payment_type will be bank_account
    MaskedBankRoutingNumber string                 `json:"masked_bank_routing_number"`
    // A string representation of the stored bank account number with all but the last 4 digits marked with X’s (i.e. ‘XXXXXXX1111’)
    MaskedBankAccountNumber string                 `json:"masked_bank_account_number"`
    // Defaults to checking
    BankAccountType         *BankAccountType       `json:"bank_account_type,omitempty"`
    // Defaults to personal
    BankAccountHolderType   *BankAccountHolderType `json:"bank_account_holder_type,omitempty"`
    PaymentType             *PaymentType           `json:"payment_type,omitempty"`
    // denotes whether a bank account has been verified by providing the amounts of two small deposits made into the account
    Verified                *bool                  `json:"verified,omitempty"`
    SiteGatewaySettingId    Optional[int]          `json:"site_gateway_setting_id"`
    GatewayHandle           Optional[string]       `json:"gateway_handle"`
    AdditionalProperties    map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BankAccountPaymentProfile.
// It customizes the JSON marshaling process for BankAccountPaymentProfile objects.
func (b BankAccountPaymentProfile) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BankAccountPaymentProfile object to a map representation for JSON marshaling.
func (b BankAccountPaymentProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    if b.Id != nil {
        structMap["id"] = b.Id
    }
    if b.FirstName != nil {
        structMap["first_name"] = b.FirstName
    }
    if b.LastName != nil {
        structMap["last_name"] = b.LastName
    }
    if b.CustomerId != nil {
        structMap["customer_id"] = b.CustomerId
    }
    if b.CurrentVault != nil {
        structMap["current_vault"] = b.CurrentVault
    }
    if b.VaultToken != nil {
        structMap["vault_token"] = b.VaultToken
    }
    if b.BillingAddress.IsValueSet() {
        if b.BillingAddress.Value() != nil {
            structMap["billing_address"] = b.BillingAddress.Value()
        } else {
            structMap["billing_address"] = nil
        }
    }
    if b.BillingCity.IsValueSet() {
        if b.BillingCity.Value() != nil {
            structMap["billing_city"] = b.BillingCity.Value()
        } else {
            structMap["billing_city"] = nil
        }
    }
    if b.BillingState.IsValueSet() {
        if b.BillingState.Value() != nil {
            structMap["billing_state"] = b.BillingState.Value()
        } else {
            structMap["billing_state"] = nil
        }
    }
    if b.BillingZip.IsValueSet() {
        if b.BillingZip.Value() != nil {
            structMap["billing_zip"] = b.BillingZip.Value()
        } else {
            structMap["billing_zip"] = nil
        }
    }
    if b.BillingCountry.IsValueSet() {
        if b.BillingCountry.Value() != nil {
            structMap["billing_country"] = b.BillingCountry.Value()
        } else {
            structMap["billing_country"] = nil
        }
    }
    if b.CustomerVaultToken.IsValueSet() {
        if b.CustomerVaultToken.Value() != nil {
            structMap["customer_vault_token"] = b.CustomerVaultToken.Value()
        } else {
            structMap["customer_vault_token"] = nil
        }
    }
    if b.BillingAddress2.IsValueSet() {
        if b.BillingAddress2.Value() != nil {
            structMap["billing_address_2"] = b.BillingAddress2.Value()
        } else {
            structMap["billing_address_2"] = nil
        }
    }
    if b.BankName != nil {
        structMap["bank_name"] = b.BankName
    }
    structMap["masked_bank_routing_number"] = b.MaskedBankRoutingNumber
    structMap["masked_bank_account_number"] = b.MaskedBankAccountNumber
    if b.BankAccountType != nil {
        structMap["bank_account_type"] = b.BankAccountType
    }
    if b.BankAccountHolderType != nil {
        structMap["bank_account_holder_type"] = b.BankAccountHolderType
    }
    if b.PaymentType != nil {
        structMap["payment_type"] = b.PaymentType
    }
    if b.Verified != nil {
        structMap["verified"] = b.Verified
    }
    if b.SiteGatewaySettingId.IsValueSet() {
        if b.SiteGatewaySettingId.Value() != nil {
            structMap["site_gateway_setting_id"] = b.SiteGatewaySettingId.Value()
        } else {
            structMap["site_gateway_setting_id"] = nil
        }
    }
    if b.GatewayHandle.IsValueSet() {
        if b.GatewayHandle.Value() != nil {
            structMap["gateway_handle"] = b.GatewayHandle.Value()
        } else {
            structMap["gateway_handle"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BankAccountPaymentProfile.
// It customizes the JSON unmarshaling process for BankAccountPaymentProfile objects.
func (b *BankAccountPaymentProfile) UnmarshalJSON(input []byte) error {
    var temp bankAccountPaymentProfile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "first_name", "last_name", "customer_id", "current_vault", "vault_token", "billing_address", "billing_city", "billing_state", "billing_zip", "billing_country", "customer_vault_token", "billing_address_2", "bank_name", "masked_bank_routing_number", "masked_bank_account_number", "bank_account_type", "bank_account_holder_type", "payment_type", "verified", "site_gateway_setting_id", "gateway_handle")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.Id = temp.Id
    b.FirstName = temp.FirstName
    b.LastName = temp.LastName
    b.CustomerId = temp.CustomerId
    b.CurrentVault = temp.CurrentVault
    b.VaultToken = temp.VaultToken
    b.BillingAddress = temp.BillingAddress
    b.BillingCity = temp.BillingCity
    b.BillingState = temp.BillingState
    b.BillingZip = temp.BillingZip
    b.BillingCountry = temp.BillingCountry
    b.CustomerVaultToken = temp.CustomerVaultToken
    b.BillingAddress2 = temp.BillingAddress2
    b.BankName = temp.BankName
    b.MaskedBankRoutingNumber = *temp.MaskedBankRoutingNumber
    b.MaskedBankAccountNumber = *temp.MaskedBankAccountNumber
    b.BankAccountType = temp.BankAccountType
    b.BankAccountHolderType = temp.BankAccountHolderType
    b.PaymentType = temp.PaymentType
    b.Verified = temp.Verified
    b.SiteGatewaySettingId = temp.SiteGatewaySettingId
    b.GatewayHandle = temp.GatewayHandle
    return nil
}

// TODO
type bankAccountPaymentProfile  struct {
    Id                      *int                   `json:"id,omitempty"`
    FirstName               *string                `json:"first_name,omitempty"`
    LastName                *string                `json:"last_name,omitempty"`
    CustomerId              *int                   `json:"customer_id,omitempty"`
    CurrentVault            *BankAccountVault      `json:"current_vault,omitempty"`
    VaultToken              *string                `json:"vault_token,omitempty"`
    BillingAddress          Optional[string]       `json:"billing_address"`
    BillingCity             Optional[string]       `json:"billing_city"`
    BillingState            Optional[string]       `json:"billing_state"`
    BillingZip              Optional[string]       `json:"billing_zip"`
    BillingCountry          Optional[string]       `json:"billing_country"`
    CustomerVaultToken      Optional[string]       `json:"customer_vault_token"`
    BillingAddress2         Optional[string]       `json:"billing_address_2"`
    BankName                *string                `json:"bank_name,omitempty"`
    MaskedBankRoutingNumber *string                `json:"masked_bank_routing_number"`
    MaskedBankAccountNumber *string                `json:"masked_bank_account_number"`
    BankAccountType         *BankAccountType       `json:"bank_account_type,omitempty"`
    BankAccountHolderType   *BankAccountHolderType `json:"bank_account_holder_type,omitempty"`
    PaymentType             *PaymentType           `json:"payment_type,omitempty"`
    Verified                *bool                  `json:"verified,omitempty"`
    SiteGatewaySettingId    Optional[int]          `json:"site_gateway_setting_id"`
    GatewayHandle           Optional[string]       `json:"gateway_handle"`
}

func (b *bankAccountPaymentProfile) validate() error {
    var errs []string
    if b.MaskedBankRoutingNumber == nil {
        errs = append(errs, "required field `masked_bank_routing_number` is missing for type `Bank Account Payment Profile`")
    }
    if b.MaskedBankAccountNumber == nil {
        errs = append(errs, "required field `masked_bank_account_number` is missing for type `Bank Account Payment Profile`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
