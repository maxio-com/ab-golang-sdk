package models

import (
	"encoding/json"
)

// SubscriptionGroupBankAccount represents a SubscriptionGroupBankAccount struct.
type SubscriptionGroupBankAccount struct {
	// (Required when creating a subscription with ACH or GoCardless) The name of the bank where the customer’s account resides
	BankName *string `json:"bank_name,omitempty"`
	// (Required when creating a subscription with ACH. Required when creating a subscription with GoCardless and bank_iban is blank) The customerʼs bank account number
	BankAccountNumber *string `json:"bank_account_number,omitempty"`
	// (Required when creating a subscription with ACH. Optional when creating a subscription with GoCardless). The routing number of the bank. It becomes bank_code while passing via GoCardless API
	BankRoutingNumber *string `json:"bank_routing_number,omitempty"`
	// (Optional when creating a subscription with GoCardless). International Bank Account Number. Alternatively, local bank details can be provided
	BankIban *string `json:"bank_iban,omitempty"`
	// (Optional when creating a subscription with GoCardless) Branch code. Alternatively, an IBAN can be provided
	BankBranchCode *string `json:"bank_branch_code,omitempty"`
	// Defaults to checking
	BankAccountType *BankAccountType `json:"bank_account_type,omitempty"`
	// Defaults to personal
	BankAccountHolderType *BankAccountHolderType `json:"bank_account_holder_type,omitempty"`
	PaymentType           *PaymentType           `json:"payment_type,omitempty"`
	BillingAddress        *string                `json:"billing_address,omitempty"`
	BillingCity           *string                `json:"billing_city,omitempty"`
	BillingState          *string                `json:"billing_state,omitempty"`
	BillingZip            *string                `json:"billing_zip,omitempty"`
	BillingCountry        *string                `json:"billing_country,omitempty"`
	ChargifyToken         *string                `json:"chargify_token,omitempty"`
	// The vault that stores the payment profile with the provided vault_token.
	CurrentVault  *BankAccountVault `json:"current_vault,omitempty"`
	GatewayHandle *string           `json:"gateway_handle,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupBankAccount.
// It customizes the JSON marshaling process for SubscriptionGroupBankAccount objects.
func (s *SubscriptionGroupBankAccount) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupBankAccount object to a map representation for JSON marshaling.
func (s *SubscriptionGroupBankAccount) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.BankName != nil {
		structMap["bank_name"] = s.BankName
	}
	if s.BankAccountNumber != nil {
		structMap["bank_account_number"] = s.BankAccountNumber
	}
	if s.BankRoutingNumber != nil {
		structMap["bank_routing_number"] = s.BankRoutingNumber
	}
	if s.BankIban != nil {
		structMap["bank_iban"] = s.BankIban
	}
	if s.BankBranchCode != nil {
		structMap["bank_branch_code"] = s.BankBranchCode
	}
	if s.BankAccountType != nil {
		structMap["bank_account_type"] = s.BankAccountType
	}
	if s.BankAccountHolderType != nil {
		structMap["bank_account_holder_type"] = s.BankAccountHolderType
	}
	if s.PaymentType != nil {
		structMap["payment_type"] = s.PaymentType
	}
	if s.BillingAddress != nil {
		structMap["billing_address"] = s.BillingAddress
	}
	if s.BillingCity != nil {
		structMap["billing_city"] = s.BillingCity
	}
	if s.BillingState != nil {
		structMap["billing_state"] = s.BillingState
	}
	if s.BillingZip != nil {
		structMap["billing_zip"] = s.BillingZip
	}
	if s.BillingCountry != nil {
		structMap["billing_country"] = s.BillingCountry
	}
	if s.ChargifyToken != nil {
		structMap["chargify_token"] = s.ChargifyToken
	}
	if s.CurrentVault != nil {
		structMap["current_vault"] = s.CurrentVault
	}
	if s.GatewayHandle != nil {
		structMap["gateway_handle"] = s.GatewayHandle
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupBankAccount.
// It customizes the JSON unmarshaling process for SubscriptionGroupBankAccount objects.
func (s *SubscriptionGroupBankAccount) UnmarshalJSON(input []byte) error {
	var temp subscriptionGroupBankAccount
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	s.BankName = temp.BankName
	s.BankAccountNumber = temp.BankAccountNumber
	s.BankRoutingNumber = temp.BankRoutingNumber
	s.BankIban = temp.BankIban
	s.BankBranchCode = temp.BankBranchCode
	s.BankAccountType = temp.BankAccountType
	s.BankAccountHolderType = temp.BankAccountHolderType
	s.PaymentType = temp.PaymentType
	s.BillingAddress = temp.BillingAddress
	s.BillingCity = temp.BillingCity
	s.BillingState = temp.BillingState
	s.BillingZip = temp.BillingZip
	s.BillingCountry = temp.BillingCountry
	s.ChargifyToken = temp.ChargifyToken
	s.CurrentVault = temp.CurrentVault
	s.GatewayHandle = temp.GatewayHandle
	return nil
}

// TODO
type subscriptionGroupBankAccount struct {
	BankName              *string                `json:"bank_name,omitempty"`
	BankAccountNumber     *string                `json:"bank_account_number,omitempty"`
	BankRoutingNumber     *string                `json:"bank_routing_number,omitempty"`
	BankIban              *string                `json:"bank_iban,omitempty"`
	BankBranchCode        *string                `json:"bank_branch_code,omitempty"`
	BankAccountType       *BankAccountType       `json:"bank_account_type,omitempty"`
	BankAccountHolderType *BankAccountHolderType `json:"bank_account_holder_type,omitempty"`
	PaymentType           *PaymentType           `json:"payment_type,omitempty"`
	BillingAddress        *string                `json:"billing_address,omitempty"`
	BillingCity           *string                `json:"billing_city,omitempty"`
	BillingState          *string                `json:"billing_state,omitempty"`
	BillingZip            *string                `json:"billing_zip,omitempty"`
	BillingCountry        *string                `json:"billing_country,omitempty"`
	ChargifyToken         *string                `json:"chargify_token,omitempty"`
	CurrentVault          *BankAccountVault      `json:"current_vault,omitempty"`
	GatewayHandle         *string                `json:"gateway_handle,omitempty"`
}
