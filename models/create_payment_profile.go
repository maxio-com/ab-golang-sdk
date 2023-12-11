package models

import (
	"encoding/json"
)

// CreatePaymentProfile represents a CreatePaymentProfile struct.
type CreatePaymentProfile struct {
	// Token received after sending billing informations using chargify.js.
	ChargifyToken *string          `json:"chargify_token,omitempty"`
	Id            *int             `json:"id,omitempty"`
	PaymentType   *PaymentTypeEnum `json:"payment_type,omitempty"`
	// First name on card or bank account. If omitted, the first_name from customer attributes will be used.
	FirstName *string `json:"first_name,omitempty"`
	// Last name on card or bank account. If omitted, the last_name from customer attributes will be used.
	LastName         *string `json:"last_name,omitempty"`
	MaskedCardNumber *string `json:"masked_card_number,omitempty"`
	// The full credit card number
	FullNumber *string `json:"full_number,omitempty"`
	// The type of card used.
	CardType *CardTypeEnum `json:"card_type,omitempty"`
	// (Optional when performing an Import via vault_token, required otherwise) The 1- or 2-digit credit card expiration month, as an integer or string, i.e. 5
	ExpirationMonth *interface{} `json:"expiration_month,omitempty"`
	// (Optional when performing a Import via vault_token, required otherwise) The 4-digit credit card expiration year, as an integer or string, i.e. 2012
	ExpirationYear *interface{} `json:"expiration_year,omitempty"`
	// The credit card or bank account billing street address (i.e. 123 Main St.). This value is merely passed through to the payment gateway.
	BillingAddress *string `json:"billing_address,omitempty"`
	// Second line of the customer’s billing address i.e. Apt. 100
	BillingAddress2 Optional[string] `json:"billing_address_2"`
	// The credit card or bank account billing address city (i.e. “Boston”). This value is merely passed through to the payment gateway.
	BillingCity *string `json:"billing_city,omitempty"`
	// The credit card or bank account billing address state (i.e. MA). This value is merely passed through to the payment gateway. This must conform to the [ISO_3166-1](https://en.wikipedia.org/wiki/ISO_3166-1#Current_codes) in order to be valid for tax locale purposes.
	BillingState *string `json:"billing_state,omitempty"`
	// The credit card or bank account billing address country, required in [ISO_3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format (i.e. “US”). This value is merely passed through to the payment gateway. Some gateways require country codes in a specific format. Please check your gateway’s documentation. If creating an ACH subscription, only US is supported at this time.
	BillingCountry *string `json:"billing_country,omitempty"`
	// The credit card or bank account billing address zip code (i.e. 12345). This value is merely passed through to the payment gateway.
	BillingZip *string `json:"billing_zip,omitempty"`
	// The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing.
	CurrentVault *CurrentVaultEnum `json:"current_vault,omitempty"`
	// The “token” provided by your vault storage for an already stored payment profile
	VaultToken *string `json:"vault_token,omitempty"`
	// (only for Authorize.Net CIM storage or Square) The customerProfileId for the owner of the customerPaymentProfileId provided as the vault_token
	CustomerVaultToken *string `json:"customer_vault_token,omitempty"`
	// (Required when creating a new payment profile) The Chargify customer id.
	CustomerId *int `json:"customer_id,omitempty"`
	// used by merchants that implemented BraintreeBlue javaScript libraries on their own. We recommend using Chargify.js instead.
	PaypalEmail *string `json:"paypal_email,omitempty"` // Deprecated
	// used by merchants that implemented BraintreeBlue javaScript libraries on their own. We recommend using Chargify.js instead.
	PaymentMethodNonce *string `json:"payment_method_nonce,omitempty"` // Deprecated
	// This attribute is only available if MultiGateway feature is enabled for your Site. This feature is in the Private Beta currently. gateway_handle is used to directly select a gateway where a payment profile will be stored in. Every connected gateway must have a unique gateway handle specified. Read [Multigateway description](https://chargify.zendesk.com/hc/en-us/articles/4407761759643#connecting-with-multiple-gateways) to learn more about new concepts that MultiGateway introduces and the default behavior when this attribute is not passed.
	GatewayHandle *string `json:"gateway_handle,omitempty"`
	// The 3- or 4-digit Card Verification Value. This value is merely passed through to the payment gateway.
	Cvv *string `json:"cvv,omitempty"`
	// (Required when creating with ACH or GoCardless, optional with Stripe Direct Debit). The name of the bank where the customerʼs account resides
	BankName *string `json:"bank_name,omitempty"`
	// (Optional when creating with GoCardless, required with Stripe Direct Debit). International Bank Account Number. Alternatively, local bank details can be provided
	BankIban *string `json:"bank_iban,omitempty"`
	// (Required when creating with ACH. Optional when creating a subscription with GoCardless). The routing number of the bank. It becomes bank_code while passing via GoCardless API
	BankRoutingNumber *string `json:"bank_routing_number,omitempty"`
	// (Required when creating with ACH, GoCardless, Stripe BECS Direct Debit and bank_iban is blank) The customerʼs bank account number
	BankAccountNumber *string `json:"bank_account_number,omitempty"`
	// (Optional when creating with GoCardless, required with Stripe BECS Direct Debit) Branch code. Alternatively, an IBAN can be provided
	BankBranchCode        *string `json:"bank_branch_code,omitempty"`
	BankAccountType       *string `json:"bank_account_type,omitempty"`
	BankAccountHolderType *string `json:"bank_account_holder_type,omitempty"`
	// (Optional) Used for creating subscription with payment profile imported using vault_token, for proper display in Advanced Billing UI
	LastFour *string `json:"last_four,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePaymentProfile.
// It customizes the JSON marshaling process for CreatePaymentProfile objects.
func (c *CreatePaymentProfile) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreatePaymentProfile object to a map representation for JSON marshaling.
func (c *CreatePaymentProfile) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.ChargifyToken != nil {
		structMap["chargify_token"] = c.ChargifyToken
	}
	if c.Id != nil {
		structMap["id"] = c.Id
	}
	if c.PaymentType != nil {
		structMap["payment_type"] = c.PaymentType
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
	if c.FullNumber != nil {
		structMap["full_number"] = c.FullNumber
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
	if c.BillingAddress != nil {
		structMap["billing_address"] = c.BillingAddress
	}
	if c.BillingAddress2.IsValueSet() {
		structMap["billing_address_2"] = c.BillingAddress2.Value()
	}
	if c.BillingCity != nil {
		structMap["billing_city"] = c.BillingCity
	}
	if c.BillingState != nil {
		structMap["billing_state"] = c.BillingState
	}
	if c.BillingCountry != nil {
		structMap["billing_country"] = c.BillingCountry
	}
	if c.BillingZip != nil {
		structMap["billing_zip"] = c.BillingZip
	}
	if c.CurrentVault != nil {
		structMap["current_vault"] = c.CurrentVault
	}
	if c.VaultToken != nil {
		structMap["vault_token"] = c.VaultToken
	}
	if c.CustomerVaultToken != nil {
		structMap["customer_vault_token"] = c.CustomerVaultToken
	}
	if c.CustomerId != nil {
		structMap["customer_id"] = c.CustomerId
	}
	if c.PaypalEmail != nil {
		structMap["paypal_email"] = c.PaypalEmail
	}
	if c.PaymentMethodNonce != nil {
		structMap["payment_method_nonce"] = c.PaymentMethodNonce
	}
	if c.GatewayHandle != nil {
		structMap["gateway_handle"] = c.GatewayHandle
	}
	if c.Cvv != nil {
		structMap["cvv"] = c.Cvv
	}
	if c.BankName != nil {
		structMap["bank_name"] = c.BankName
	}
	if c.BankIban != nil {
		structMap["bank_iban"] = c.BankIban
	}
	if c.BankRoutingNumber != nil {
		structMap["bank_routing_number"] = c.BankRoutingNumber
	}
	if c.BankAccountNumber != nil {
		structMap["bank_account_number"] = c.BankAccountNumber
	}
	if c.BankBranchCode != nil {
		structMap["bank_branch_code"] = c.BankBranchCode
	}
	if c.BankAccountType != nil {
		structMap["bank_account_type"] = c.BankAccountType
	}
	if c.BankAccountHolderType != nil {
		structMap["bank_account_holder_type"] = c.BankAccountHolderType
	}
	if c.LastFour != nil {
		structMap["last_four"] = c.LastFour
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePaymentProfile.
// It customizes the JSON unmarshaling process for CreatePaymentProfile objects.
func (c *CreatePaymentProfile) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ChargifyToken         *string           `json:"chargify_token,omitempty"`
		Id                    *int              `json:"id,omitempty"`
		PaymentType           *PaymentTypeEnum  `json:"payment_type,omitempty"`
		FirstName             *string           `json:"first_name,omitempty"`
		LastName              *string           `json:"last_name,omitempty"`
		MaskedCardNumber      *string           `json:"masked_card_number,omitempty"`
		FullNumber            *string           `json:"full_number,omitempty"`
		CardType              *CardTypeEnum     `json:"card_type,omitempty"`
		ExpirationMonth       *interface{}      `json:"expiration_month,omitempty"`
		ExpirationYear        *interface{}      `json:"expiration_year,omitempty"`
		BillingAddress        *string           `json:"billing_address,omitempty"`
		BillingAddress2       Optional[string]  `json:"billing_address_2"`
		BillingCity           *string           `json:"billing_city,omitempty"`
		BillingState          *string           `json:"billing_state,omitempty"`
		BillingCountry        *string           `json:"billing_country,omitempty"`
		BillingZip            *string           `json:"billing_zip,omitempty"`
		CurrentVault          *CurrentVaultEnum `json:"current_vault,omitempty"`
		VaultToken            *string           `json:"vault_token,omitempty"`
		CustomerVaultToken    *string           `json:"customer_vault_token,omitempty"`
		CustomerId            *int              `json:"customer_id,omitempty"`
		PaypalEmail           *string           `json:"paypal_email,omitempty"`
		PaymentMethodNonce    *string           `json:"payment_method_nonce,omitempty"`
		GatewayHandle         *string           `json:"gateway_handle,omitempty"`
		Cvv                   *string           `json:"cvv,omitempty"`
		BankName              *string           `json:"bank_name,omitempty"`
		BankIban              *string           `json:"bank_iban,omitempty"`
		BankRoutingNumber     *string           `json:"bank_routing_number,omitempty"`
		BankAccountNumber     *string           `json:"bank_account_number,omitempty"`
		BankBranchCode        *string           `json:"bank_branch_code,omitempty"`
		BankAccountType       *string           `json:"bank_account_type,omitempty"`
		BankAccountHolderType *string           `json:"bank_account_holder_type,omitempty"`
		LastFour              *string           `json:"last_four,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.ChargifyToken = temp.ChargifyToken
	c.Id = temp.Id
	c.PaymentType = temp.PaymentType
	c.FirstName = temp.FirstName
	c.LastName = temp.LastName
	c.MaskedCardNumber = temp.MaskedCardNumber
	c.FullNumber = temp.FullNumber
	c.CardType = temp.CardType
	c.ExpirationMonth = temp.ExpirationMonth
	c.ExpirationYear = temp.ExpirationYear
	c.BillingAddress = temp.BillingAddress
	c.BillingAddress2 = temp.BillingAddress2
	c.BillingCity = temp.BillingCity
	c.BillingState = temp.BillingState
	c.BillingCountry = temp.BillingCountry
	c.BillingZip = temp.BillingZip
	c.CurrentVault = temp.CurrentVault
	c.VaultToken = temp.VaultToken
	c.CustomerVaultToken = temp.CustomerVaultToken
	c.CustomerId = temp.CustomerId
	c.PaypalEmail = temp.PaypalEmail
	c.PaymentMethodNonce = temp.PaymentMethodNonce
	c.GatewayHandle = temp.GatewayHandle
	c.Cvv = temp.Cvv
	c.BankName = temp.BankName
	c.BankIban = temp.BankIban
	c.BankRoutingNumber = temp.BankRoutingNumber
	c.BankAccountNumber = temp.BankAccountNumber
	c.BankBranchCode = temp.BankBranchCode
	c.BankAccountType = temp.BankAccountType
	c.BankAccountHolderType = temp.BankAccountHolderType
	c.LastFour = temp.LastFour
	return nil
}
