package models

import (
    "encoding/json"
)

// BankAccount represents a BankAccount struct.
type BankAccount struct {
    Id                      *int              `json:"id,omitempty"`
    FirstName               *string           `json:"first_name,omitempty"`
    LastName                *string           `json:"last_name,omitempty"`
    CustomerId              *int              `json:"customer_id,omitempty"`
    // The vault that stores the payment profile with the provided vault_token.
    CurrentVault            *BankAccountVault `json:"current_vault,omitempty"`
    VaultToken              *string           `json:"vault_token,omitempty"`
    BillingAddress          *string           `json:"billing_address,omitempty"`
    BillingCity             *string           `json:"billing_city,omitempty"`
    BillingState            *string           `json:"billing_state,omitempty"`
    BillingZip              *string           `json:"billing_zip,omitempty"`
    BillingCountry          *string           `json:"billing_country,omitempty"`
    CustomerVaultToken      Optional[string]  `json:"customer_vault_token"`
    BillingAddress2         *string           `json:"billing_address_2,omitempty"`
    BankName                *string           `json:"bank_name,omitempty"`
    MaskedBankRoutingNumber *string           `json:"masked_bank_routing_number,omitempty"`
    MaskedBankAccountNumber *string           `json:"masked_bank_account_number,omitempty"`
    BankAccountType         *string           `json:"bank_account_type,omitempty"`
    BankAccountHolderType   *string           `json:"bank_account_holder_type,omitempty"`
    PaymentType             *string           `json:"payment_type,omitempty"`
    // denotes whether a bank account has been verified by providing the amounts of two small deposits made into the account
    Verified                *bool             `json:"verified,omitempty"`
    SiteGatewaySettingId    *int              `json:"site_gateway_setting_id,omitempty"`
    GatewayHandle           *string           `json:"gateway_handle,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for BankAccount.
// It customizes the JSON marshaling process for BankAccount objects.
func (b *BankAccount) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BankAccount object to a map representation for JSON marshaling.
func (b *BankAccount) toMap() map[string]any {
    structMap := make(map[string]any)
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
    if b.BillingAddress != nil {
        structMap["billing_address"] = b.BillingAddress
    }
    if b.BillingCity != nil {
        structMap["billing_city"] = b.BillingCity
    }
    if b.BillingState != nil {
        structMap["billing_state"] = b.BillingState
    }
    if b.BillingZip != nil {
        structMap["billing_zip"] = b.BillingZip
    }
    if b.BillingCountry != nil {
        structMap["billing_country"] = b.BillingCountry
    }
    if b.CustomerVaultToken.IsValueSet() {
        structMap["customer_vault_token"] = b.CustomerVaultToken.Value()
    }
    if b.BillingAddress2 != nil {
        structMap["billing_address_2"] = b.BillingAddress2
    }
    if b.BankName != nil {
        structMap["bank_name"] = b.BankName
    }
    if b.MaskedBankRoutingNumber != nil {
        structMap["masked_bank_routing_number"] = b.MaskedBankRoutingNumber
    }
    if b.MaskedBankAccountNumber != nil {
        structMap["masked_bank_account_number"] = b.MaskedBankAccountNumber
    }
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
    if b.SiteGatewaySettingId != nil {
        structMap["site_gateway_setting_id"] = b.SiteGatewaySettingId
    }
    if b.GatewayHandle != nil {
        structMap["gateway_handle"] = b.GatewayHandle
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BankAccount.
// It customizes the JSON unmarshaling process for BankAccount objects.
func (b *BankAccount) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                      *int              `json:"id,omitempty"`
        FirstName               *string           `json:"first_name,omitempty"`
        LastName                *string           `json:"last_name,omitempty"`
        CustomerId              *int              `json:"customer_id,omitempty"`
        CurrentVault            *BankAccountVault `json:"current_vault,omitempty"`
        VaultToken              *string           `json:"vault_token,omitempty"`
        BillingAddress          *string           `json:"billing_address,omitempty"`
        BillingCity             *string           `json:"billing_city,omitempty"`
        BillingState            *string           `json:"billing_state,omitempty"`
        BillingZip              *string           `json:"billing_zip,omitempty"`
        BillingCountry          *string           `json:"billing_country,omitempty"`
        CustomerVaultToken      Optional[string]  `json:"customer_vault_token"`
        BillingAddress2         *string           `json:"billing_address_2,omitempty"`
        BankName                *string           `json:"bank_name,omitempty"`
        MaskedBankRoutingNumber *string           `json:"masked_bank_routing_number,omitempty"`
        MaskedBankAccountNumber *string           `json:"masked_bank_account_number,omitempty"`
        BankAccountType         *string           `json:"bank_account_type,omitempty"`
        BankAccountHolderType   *string           `json:"bank_account_holder_type,omitempty"`
        PaymentType             *string           `json:"payment_type,omitempty"`
        Verified                *bool             `json:"verified,omitempty"`
        SiteGatewaySettingId    *int              `json:"site_gateway_setting_id,omitempty"`
        GatewayHandle           *string           `json:"gateway_handle,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
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
    b.MaskedBankRoutingNumber = temp.MaskedBankRoutingNumber
    b.MaskedBankAccountNumber = temp.MaskedBankAccountNumber
    b.BankAccountType = temp.BankAccountType
    b.BankAccountHolderType = temp.BankAccountHolderType
    b.PaymentType = temp.PaymentType
    b.Verified = temp.Verified
    b.SiteGatewaySettingId = temp.SiteGatewaySettingId
    b.GatewayHandle = temp.GatewayHandle
    return nil
}
