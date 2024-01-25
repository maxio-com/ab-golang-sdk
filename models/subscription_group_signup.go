package models

import (
	"encoding/json"
)

// SubscriptionGroupSignup represents a SubscriptionGroupSignup struct.
type SubscriptionGroupSignup struct {
	PaymentProfileId *int    `json:"payment_profile_id,omitempty"`
	PayerId          *int    `json:"payer_id,omitempty"`
	PayerReference   *string `json:"payer_reference,omitempty"`
	// The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`.
	PaymentCollectionMethod *PaymentCollectionMethod      `json:"payment_collection_method,omitempty"`
	PayerAttributes         *PayerAttributes              `json:"payer_attributes,omitempty"`
	CreditCardAttributes    *SubscriptionGroupCreditCard  `json:"credit_card_attributes,omitempty"`
	BankAccountAttributes   *SubscriptionGroupBankAccount `json:"bank_account_attributes,omitempty"`
	Subscriptions           []SubscriptionGroupSignupItem `json:"subscriptions"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignup.
// It customizes the JSON marshaling process for SubscriptionGroupSignup objects.
func (s *SubscriptionGroupSignup) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignup object to a map representation for JSON marshaling.
func (s *SubscriptionGroupSignup) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.PaymentProfileId != nil {
		structMap["payment_profile_id"] = s.PaymentProfileId
	}
	if s.PayerId != nil {
		structMap["payer_id"] = s.PayerId
	}
	if s.PayerReference != nil {
		structMap["payer_reference"] = s.PayerReference
	}
	if s.PaymentCollectionMethod != nil {
		structMap["payment_collection_method"] = s.PaymentCollectionMethod
	}
	if s.PayerAttributes != nil {
		structMap["payer_attributes"] = s.PayerAttributes
	}
	if s.CreditCardAttributes != nil {
		structMap["credit_card_attributes"] = s.CreditCardAttributes
	}
	if s.BankAccountAttributes != nil {
		structMap["bank_account_attributes"] = s.BankAccountAttributes
	}
	structMap["subscriptions"] = s.Subscriptions
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignup.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignup objects.
func (s *SubscriptionGroupSignup) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PaymentProfileId        *int                          `json:"payment_profile_id,omitempty"`
		PayerId                 *int                          `json:"payer_id,omitempty"`
		PayerReference          *string                       `json:"payer_reference,omitempty"`
		PaymentCollectionMethod *PaymentCollectionMethod      `json:"payment_collection_method,omitempty"`
		PayerAttributes         *PayerAttributes              `json:"payer_attributes,omitempty"`
		CreditCardAttributes    *SubscriptionGroupCreditCard  `json:"credit_card_attributes,omitempty"`
		BankAccountAttributes   *SubscriptionGroupBankAccount `json:"bank_account_attributes,omitempty"`
		Subscriptions           []SubscriptionGroupSignupItem `json:"subscriptions"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.PaymentProfileId = temp.PaymentProfileId
	s.PayerId = temp.PayerId
	s.PayerReference = temp.PayerReference
	s.PaymentCollectionMethod = temp.PaymentCollectionMethod
	s.PayerAttributes = temp.PayerAttributes
	s.CreditCardAttributes = temp.CreditCardAttributes
	s.BankAccountAttributes = temp.BankAccountAttributes
	s.Subscriptions = temp.Subscriptions
	return nil
}
