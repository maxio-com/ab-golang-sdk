package models

import (
	"encoding/json"
)

// SubscriptionBankAccount represents a SubscriptionBankAccount struct.
type SubscriptionBankAccount struct {
	// Defaults to personal
	BankAccountHolderType *string `json:"bank_account_holder_type,omitempty"`
	// Defaults to checking
	BankAccountType *string `json:"bank_account_type,omitempty"`
	// The bank where the account resides
	BankName *string `json:"bank_name,omitempty"`
	// The current billing street address for the bank account
	BillingAddress *string `json:"billing_address,omitempty"`
	// The current billing street address, second line, for the bank account
	BillingAddress2 *string `json:"billing_address_2,omitempty"`
	// The current billing address city for the bank account
	BillingCity *string `json:"billing_city,omitempty"`
	// The current billing address state for the bank account
	BillingState *string `json:"billing_state,omitempty"`
	// The current billing address zip code for the bank account
	BillingZip *string `json:"billing_zip,omitempty"`
	// The current billing address country for the bank account
	BillingCountry *string `json:"billing_country,omitempty"`
	// The vault that stores the payment profile with the provided vault_token.
	CurrentVault *BankAccountVaultEnum `json:"current_vault,omitempty"`
	// The Chargify-assigned id for the customer record to which the bank account belongs
	CustomerId *int `json:"customer_id,omitempty"`
	// (only for Authorize.Net CIM storage): the customerProfileId for the owner of the customerPaymentProfileId provided as the vault_token
	CustomerVaultToken *string `json:"customer_vault_token,omitempty"`
	// The first name of the bank account holder
	FirstName *string `json:"first_name,omitempty"`
	// The last name of the bank account holder
	LastName *string `json:"last_name,omitempty"`
	// The Chargify-assigned ID of the stored bank account. This value can be used as an input to payment_profile_id when creating a subscription, in order to re-use a stored payment profile for the same customer
	Id *int `json:"id,omitempty"`
	// A string representation of the stored bank account number with all but the last 4 digits marked with X’s (i.e. ‘XXXXXXX1111’)
	MaskedBankAccountNumber *string `json:"masked_bank_account_number,omitempty"`
	// A string representation of the stored bank routing number with all but the last 4 digits marked with X’s (i.e. ‘XXXXXXX1111’). payment_type will be bank_account
	MaskedBankRoutingNumber *string `json:"masked_bank_routing_number,omitempty"`
	// The “token” provided by your vault storage for an already stored payment profile
	VaultToken *string `json:"vault_token,omitempty"`
	// Token received after sending billing informations using chargify.js. This token will only be received if passed as a sole attribute of credit_card_attributes (i.e. tok_9g6hw85pnpt6knmskpwp4ttt)
	ChargifyToken        *string `json:"chargify_token,omitempty"`
	SiteGatewaySettingId *int    `json:"site_gateway_setting_id,omitempty"`
	GatewayHandle        *string `json:"gateway_handle,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionBankAccount.
// It customizes the JSON marshaling process for SubscriptionBankAccount objects.
func (s *SubscriptionBankAccount) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionBankAccount object to a map representation for JSON marshaling.
func (s *SubscriptionBankAccount) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.BankAccountHolderType != nil {
		structMap["bank_account_holder_type"] = s.BankAccountHolderType
	}
	if s.BankAccountType != nil {
		structMap["bank_account_type"] = s.BankAccountType
	}
	if s.BankName != nil {
		structMap["bank_name"] = s.BankName
	}
	if s.BillingAddress != nil {
		structMap["billing_address"] = s.BillingAddress
	}
	if s.BillingAddress2 != nil {
		structMap["billing_address_2"] = s.BillingAddress2
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
	if s.CurrentVault != nil {
		structMap["current_vault"] = s.CurrentVault
	}
	if s.CustomerId != nil {
		structMap["customer_id"] = s.CustomerId
	}
	if s.CustomerVaultToken != nil {
		structMap["customer_vault_token"] = s.CustomerVaultToken
	}
	if s.FirstName != nil {
		structMap["first_name"] = s.FirstName
	}
	if s.LastName != nil {
		structMap["last_name"] = s.LastName
	}
	if s.Id != nil {
		structMap["id"] = s.Id
	}
	if s.MaskedBankAccountNumber != nil {
		structMap["masked_bank_account_number"] = s.MaskedBankAccountNumber
	}
	if s.MaskedBankRoutingNumber != nil {
		structMap["masked_bank_routing_number"] = s.MaskedBankRoutingNumber
	}
	if s.VaultToken != nil {
		structMap["vault_token"] = s.VaultToken
	}
	if s.ChargifyToken != nil {
		structMap["chargify_token"] = s.ChargifyToken
	}
	if s.SiteGatewaySettingId != nil {
		structMap["site_gateway_setting_id"] = s.SiteGatewaySettingId
	}
	if s.GatewayHandle != nil {
		structMap["gateway_handle"] = s.GatewayHandle
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionBankAccount.
// It customizes the JSON unmarshaling process for SubscriptionBankAccount objects.
func (s *SubscriptionBankAccount) UnmarshalJSON(input []byte) error {
	temp := &struct {
		BankAccountHolderType   *string               `json:"bank_account_holder_type,omitempty"`
		BankAccountType         *string               `json:"bank_account_type,omitempty"`
		BankName                *string               `json:"bank_name,omitempty"`
		BillingAddress          *string               `json:"billing_address,omitempty"`
		BillingAddress2         *string               `json:"billing_address_2,omitempty"`
		BillingCity             *string               `json:"billing_city,omitempty"`
		BillingState            *string               `json:"billing_state,omitempty"`
		BillingZip              *string               `json:"billing_zip,omitempty"`
		BillingCountry          *string               `json:"billing_country,omitempty"`
		CurrentVault            *BankAccountVaultEnum `json:"current_vault,omitempty"`
		CustomerId              *int                  `json:"customer_id,omitempty"`
		CustomerVaultToken      *string               `json:"customer_vault_token,omitempty"`
		FirstName               *string               `json:"first_name,omitempty"`
		LastName                *string               `json:"last_name,omitempty"`
		Id                      *int                  `json:"id,omitempty"`
		MaskedBankAccountNumber *string               `json:"masked_bank_account_number,omitempty"`
		MaskedBankRoutingNumber *string               `json:"masked_bank_routing_number,omitempty"`
		VaultToken              *string               `json:"vault_token,omitempty"`
		ChargifyToken           *string               `json:"chargify_token,omitempty"`
		SiteGatewaySettingId    *int                  `json:"site_gateway_setting_id,omitempty"`
		GatewayHandle           *string               `json:"gateway_handle,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.BankAccountHolderType = temp.BankAccountHolderType
	s.BankAccountType = temp.BankAccountType
	s.BankName = temp.BankName
	s.BillingAddress = temp.BillingAddress
	s.BillingAddress2 = temp.BillingAddress2
	s.BillingCity = temp.BillingCity
	s.BillingState = temp.BillingState
	s.BillingZip = temp.BillingZip
	s.BillingCountry = temp.BillingCountry
	s.CurrentVault = temp.CurrentVault
	s.CustomerId = temp.CustomerId
	s.CustomerVaultToken = temp.CustomerVaultToken
	s.FirstName = temp.FirstName
	s.LastName = temp.LastName
	s.Id = temp.Id
	s.MaskedBankAccountNumber = temp.MaskedBankAccountNumber
	s.MaskedBankRoutingNumber = temp.MaskedBankRoutingNumber
	s.VaultToken = temp.VaultToken
	s.ChargifyToken = temp.ChargifyToken
	s.SiteGatewaySettingId = temp.SiteGatewaySettingId
	s.GatewayHandle = temp.GatewayHandle
	return nil
}
