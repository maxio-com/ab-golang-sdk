package models

import (
    "encoding/json"
)

// SubscriptionGroupSignupFailureData represents a SubscriptionGroupSignupFailureData struct.
type SubscriptionGroupSignupFailureData struct {
    PayerId                 *int                          `json:"payer_id,omitempty"`
    PayerReference          *string                       `json:"payer_reference,omitempty"`
    PaymentProfileId        *int                          `json:"payment_profile_id,omitempty"`
    PaymentCollectionMethod *string                       `json:"payment_collection_method,omitempty"`
    PayerAttributes         *PayerAttributes              `json:"payer_attributes,omitempty"`
    CreditCardAttributes    *SubscriptionGroupCreditCard  `json:"credit_card_attributes,omitempty"`
    BankAccountAttributes   *SubscriptionGroupBankAccount `json:"bank_account_attributes,omitempty"`
    Subscriptions           []SubscriptionGroupSignupItem `json:"subscriptions,omitempty"`
    AdditionalProperties    map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupFailureData.
// It customizes the JSON marshaling process for SubscriptionGroupSignupFailureData objects.
func (s SubscriptionGroupSignupFailureData) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupFailureData object to a map representation for JSON marshaling.
func (s SubscriptionGroupSignupFailureData) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.PayerId != nil {
        structMap["payer_id"] = s.PayerId
    }
    if s.PayerReference != nil {
        structMap["payer_reference"] = s.PayerReference
    }
    if s.PaymentProfileId != nil {
        structMap["payment_profile_id"] = s.PaymentProfileId
    }
    if s.PaymentCollectionMethod != nil {
        structMap["payment_collection_method"] = s.PaymentCollectionMethod
    }
    if s.PayerAttributes != nil {
        structMap["payer_attributes"] = s.PayerAttributes.toMap()
    }
    if s.CreditCardAttributes != nil {
        structMap["credit_card_attributes"] = s.CreditCardAttributes.toMap()
    }
    if s.BankAccountAttributes != nil {
        structMap["bank_account_attributes"] = s.BankAccountAttributes.toMap()
    }
    if s.Subscriptions != nil {
        structMap["subscriptions"] = s.Subscriptions
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupFailureData.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupFailureData objects.
func (s *SubscriptionGroupSignupFailureData) UnmarshalJSON(input []byte) error {
    var temp subscriptionGroupSignupFailureData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "payer_id", "payer_reference", "payment_profile_id", "payment_collection_method", "payer_attributes", "credit_card_attributes", "bank_account_attributes", "subscriptions")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.PayerId = temp.PayerId
    s.PayerReference = temp.PayerReference
    s.PaymentProfileId = temp.PaymentProfileId
    s.PaymentCollectionMethod = temp.PaymentCollectionMethod
    s.PayerAttributes = temp.PayerAttributes
    s.CreditCardAttributes = temp.CreditCardAttributes
    s.BankAccountAttributes = temp.BankAccountAttributes
    s.Subscriptions = temp.Subscriptions
    return nil
}

// TODO
type subscriptionGroupSignupFailureData  struct {
    PayerId                 *int                          `json:"payer_id,omitempty"`
    PayerReference          *string                       `json:"payer_reference,omitempty"`
    PaymentProfileId        *int                          `json:"payment_profile_id,omitempty"`
    PaymentCollectionMethod *string                       `json:"payment_collection_method,omitempty"`
    PayerAttributes         *PayerAttributes              `json:"payer_attributes,omitempty"`
    CreditCardAttributes    *SubscriptionGroupCreditCard  `json:"credit_card_attributes,omitempty"`
    BankAccountAttributes   *SubscriptionGroupBankAccount `json:"bank_account_attributes,omitempty"`
    Subscriptions           []SubscriptionGroupSignupItem `json:"subscriptions,omitempty"`
}
