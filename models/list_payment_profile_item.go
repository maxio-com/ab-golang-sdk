package models

import (
	"encoding/json"
)

// ListPaymentProfileItem represents a ListPaymentProfileItem struct.
type ListPaymentProfileItem struct {
	Id         *int    `json:"id,omitempty"`
	FirstName  *string `json:"first_name,omitempty"`
	LastName   *string `json:"last_name,omitempty"`
	CustomerId *int    `json:"customer_id,omitempty"`
	// The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing.
	CurrentVault            *CurrentVaultEnum `json:"current_vault,omitempty"`
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
	SiteGatewaySettingId    *int              `json:"site_gateway_setting_id,omitempty"`
	GatewayHandle           *string           `json:"gateway_handle,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ListPaymentProfileItem.
// It customizes the JSON marshaling process for ListPaymentProfileItem objects.
func (l *ListPaymentProfileItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListPaymentProfileItem object to a map representation for JSON marshaling.
func (l *ListPaymentProfileItem) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Id != nil {
		structMap["id"] = l.Id
	}
	if l.FirstName != nil {
		structMap["first_name"] = l.FirstName
	}
	if l.LastName != nil {
		structMap["last_name"] = l.LastName
	}
	if l.CustomerId != nil {
		structMap["customer_id"] = l.CustomerId
	}
	if l.CurrentVault != nil {
		structMap["current_vault"] = l.CurrentVault
	}
	if l.VaultToken != nil {
		structMap["vault_token"] = l.VaultToken
	}
	if l.BillingAddress != nil {
		structMap["billing_address"] = l.BillingAddress
	}
	if l.BillingCity != nil {
		structMap["billing_city"] = l.BillingCity
	}
	if l.BillingState != nil {
		structMap["billing_state"] = l.BillingState
	}
	if l.BillingZip != nil {
		structMap["billing_zip"] = l.BillingZip
	}
	if l.BillingCountry != nil {
		structMap["billing_country"] = l.BillingCountry
	}
	if l.CustomerVaultToken.IsValueSet() {
		structMap["customer_vault_token"] = l.CustomerVaultToken.Value()
	}
	if l.BillingAddress2 != nil {
		structMap["billing_address_2"] = l.BillingAddress2
	}
	if l.BankName != nil {
		structMap["bank_name"] = l.BankName
	}
	if l.MaskedBankRoutingNumber != nil {
		structMap["masked_bank_routing_number"] = l.MaskedBankRoutingNumber
	}
	if l.MaskedBankAccountNumber != nil {
		structMap["masked_bank_account_number"] = l.MaskedBankAccountNumber
	}
	if l.BankAccountType != nil {
		structMap["bank_account_type"] = l.BankAccountType
	}
	if l.BankAccountHolderType != nil {
		structMap["bank_account_holder_type"] = l.BankAccountHolderType
	}
	if l.PaymentType != nil {
		structMap["payment_type"] = l.PaymentType
	}
	if l.SiteGatewaySettingId != nil {
		structMap["site_gateway_setting_id"] = l.SiteGatewaySettingId
	}
	if l.GatewayHandle != nil {
		structMap["gateway_handle"] = l.GatewayHandle
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListPaymentProfileItem.
// It customizes the JSON unmarshaling process for ListPaymentProfileItem objects.
func (l *ListPaymentProfileItem) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                      *int              `json:"id,omitempty"`
		FirstName               *string           `json:"first_name,omitempty"`
		LastName                *string           `json:"last_name,omitempty"`
		CustomerId              *int              `json:"customer_id,omitempty"`
		CurrentVault            *CurrentVaultEnum `json:"current_vault,omitempty"`
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
		SiteGatewaySettingId    *int              `json:"site_gateway_setting_id,omitempty"`
		GatewayHandle           *string           `json:"gateway_handle,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Id = temp.Id
	l.FirstName = temp.FirstName
	l.LastName = temp.LastName
	l.CustomerId = temp.CustomerId
	l.CurrentVault = temp.CurrentVault
	l.VaultToken = temp.VaultToken
	l.BillingAddress = temp.BillingAddress
	l.BillingCity = temp.BillingCity
	l.BillingState = temp.BillingState
	l.BillingZip = temp.BillingZip
	l.BillingCountry = temp.BillingCountry
	l.CustomerVaultToken = temp.CustomerVaultToken
	l.BillingAddress2 = temp.BillingAddress2
	l.BankName = temp.BankName
	l.MaskedBankRoutingNumber = temp.MaskedBankRoutingNumber
	l.MaskedBankAccountNumber = temp.MaskedBankAccountNumber
	l.BankAccountType = temp.BankAccountType
	l.BankAccountHolderType = temp.BankAccountHolderType
	l.PaymentType = temp.PaymentType
	l.SiteGatewaySettingId = temp.SiteGatewaySettingId
	l.GatewayHandle = temp.GatewayHandle
	return nil
}
