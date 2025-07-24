// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// BankAccountAttributes represents a BankAccountAttributes struct.
type BankAccountAttributes struct {
    ChargifyToken         *string                `json:"chargify_token,omitempty"`
    // (Required when creating a subscription with ACH or GoCardless) The name of the bank where the customer’s account resides
    BankName              *string                `json:"bank_name,omitempty"`
    // (Required when creating a subscription with ACH. Optional when creating a subscription with GoCardless). The routing number of the bank. It becomes bank_code while passing via GoCardless API
    BankRoutingNumber     *string                `json:"bank_routing_number,omitempty"`
    // (Required when creating a subscription with ACH. Required when creating a subscription with GoCardless and bank_iban is blank) The customerʼs bank account number
    BankAccountNumber     *string                `json:"bank_account_number,omitempty"`
    // Defaults to checking
    BankAccountType       *BankAccountType       `json:"bank_account_type,omitempty"`
    // (Optional when creating a subscription with GoCardless) Branch code. Alternatively, an IBAN can be provided
    BankBranchCode        *string                `json:"bank_branch_code,omitempty"`
    // (Optional when creating a subscription with GoCardless). International Bank Account Number. Alternatively, local bank details can be provided
    BankIban              *string                `json:"bank_iban,omitempty"`
    // Defaults to personal
    BankAccountHolderType *BankAccountHolderType `json:"bank_account_holder_type,omitempty"`
    PaymentType           *PaymentType           `json:"payment_type,omitempty"`
    // The vault that stores the payment profile with the provided vault_token. Use `bogus` for testing.
    CurrentVault          *BankAccountVault      `json:"current_vault,omitempty"`
    VaultToken            *string                `json:"vault_token,omitempty"`
    // (only for Authorize.Net CIM storage or Square) The customerProfileId for the owner of the customerPaymentProfileId provided as the vault_token
    CustomerVaultToken    *string                `json:"customer_vault_token,omitempty"`
    AdditionalProperties  map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for BankAccountAttributes,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b BankAccountAttributes) String() string {
    return fmt.Sprintf(
    	"BankAccountAttributes[ChargifyToken=%v, BankName=%v, BankRoutingNumber=%v, BankAccountNumber=%v, BankAccountType=%v, BankBranchCode=%v, BankIban=%v, BankAccountHolderType=%v, PaymentType=%v, CurrentVault=%v, VaultToken=%v, CustomerVaultToken=%v, AdditionalProperties=%v]",
    	b.ChargifyToken, b.BankName, b.BankRoutingNumber, b.BankAccountNumber, b.BankAccountType, b.BankBranchCode, b.BankIban, b.BankAccountHolderType, b.PaymentType, b.CurrentVault, b.VaultToken, b.CustomerVaultToken, b.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for BankAccountAttributes.
// It customizes the JSON marshaling process for BankAccountAttributes objects.
func (b BankAccountAttributes) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "chargify_token", "bank_name", "bank_routing_number", "bank_account_number", "bank_account_type", "bank_branch_code", "bank_iban", "bank_account_holder_type", "payment_type", "current_vault", "vault_token", "customer_vault_token"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BankAccountAttributes object to a map representation for JSON marshaling.
func (b BankAccountAttributes) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
    if b.ChargifyToken != nil {
        structMap["chargify_token"] = b.ChargifyToken
    }
    if b.BankName != nil {
        structMap["bank_name"] = b.BankName
    }
    if b.BankRoutingNumber != nil {
        structMap["bank_routing_number"] = b.BankRoutingNumber
    }
    if b.BankAccountNumber != nil {
        structMap["bank_account_number"] = b.BankAccountNumber
    }
    if b.BankAccountType != nil {
        structMap["bank_account_type"] = b.BankAccountType
    }
    if b.BankBranchCode != nil {
        structMap["bank_branch_code"] = b.BankBranchCode
    }
    if b.BankIban != nil {
        structMap["bank_iban"] = b.BankIban
    }
    if b.BankAccountHolderType != nil {
        structMap["bank_account_holder_type"] = b.BankAccountHolderType
    }
    if b.PaymentType != nil {
        structMap["payment_type"] = b.PaymentType
    }
    if b.CurrentVault != nil {
        structMap["current_vault"] = b.CurrentVault
    }
    if b.VaultToken != nil {
        structMap["vault_token"] = b.VaultToken
    }
    if b.CustomerVaultToken != nil {
        structMap["customer_vault_token"] = b.CustomerVaultToken
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BankAccountAttributes.
// It customizes the JSON unmarshaling process for BankAccountAttributes objects.
func (b *BankAccountAttributes) UnmarshalJSON(input []byte) error {
    var temp tempBankAccountAttributes
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "chargify_token", "bank_name", "bank_routing_number", "bank_account_number", "bank_account_type", "bank_branch_code", "bank_iban", "bank_account_holder_type", "payment_type", "current_vault", "vault_token", "customer_vault_token")
    if err != nil {
    	return err
    }
    b.AdditionalProperties = additionalProperties
    
    b.ChargifyToken = temp.ChargifyToken
    b.BankName = temp.BankName
    b.BankRoutingNumber = temp.BankRoutingNumber
    b.BankAccountNumber = temp.BankAccountNumber
    b.BankAccountType = temp.BankAccountType
    b.BankBranchCode = temp.BankBranchCode
    b.BankIban = temp.BankIban
    b.BankAccountHolderType = temp.BankAccountHolderType
    b.PaymentType = temp.PaymentType
    b.CurrentVault = temp.CurrentVault
    b.VaultToken = temp.VaultToken
    b.CustomerVaultToken = temp.CustomerVaultToken
    return nil
}

// tempBankAccountAttributes is a temporary struct used for validating the fields of BankAccountAttributes.
type tempBankAccountAttributes  struct {
    ChargifyToken         *string                `json:"chargify_token,omitempty"`
    BankName              *string                `json:"bank_name,omitempty"`
    BankRoutingNumber     *string                `json:"bank_routing_number,omitempty"`
    BankAccountNumber     *string                `json:"bank_account_number,omitempty"`
    BankAccountType       *BankAccountType       `json:"bank_account_type,omitempty"`
    BankBranchCode        *string                `json:"bank_branch_code,omitempty"`
    BankIban              *string                `json:"bank_iban,omitempty"`
    BankAccountHolderType *BankAccountHolderType `json:"bank_account_holder_type,omitempty"`
    PaymentType           *PaymentType           `json:"payment_type,omitempty"`
    CurrentVault          *BankAccountVault      `json:"current_vault,omitempty"`
    VaultToken            *string                `json:"vault_token,omitempty"`
    CustomerVaultToken    *string                `json:"customer_vault_token,omitempty"`
}
