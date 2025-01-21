/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
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
    AdditionalProperties    map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionGroupSignupFailureData,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroupSignupFailureData) String() string {
    return fmt.Sprintf(
    	"SubscriptionGroupSignupFailureData[PayerId=%v, PayerReference=%v, PaymentProfileId=%v, PaymentCollectionMethod=%v, PayerAttributes=%v, CreditCardAttributes=%v, BankAccountAttributes=%v, Subscriptions=%v, AdditionalProperties=%v]",
    	s.PayerId, s.PayerReference, s.PaymentProfileId, s.PaymentCollectionMethod, s.PayerAttributes, s.CreditCardAttributes, s.BankAccountAttributes, s.Subscriptions, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupFailureData.
// It customizes the JSON marshaling process for SubscriptionGroupSignupFailureData objects.
func (s SubscriptionGroupSignupFailureData) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "payer_id", "payer_reference", "payment_profile_id", "payment_collection_method", "payer_attributes", "credit_card_attributes", "bank_account_attributes", "subscriptions"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupFailureData object to a map representation for JSON marshaling.
func (s SubscriptionGroupSignupFailureData) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSubscriptionGroupSignupFailureData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "payer_id", "payer_reference", "payment_profile_id", "payment_collection_method", "payer_attributes", "credit_card_attributes", "bank_account_attributes", "subscriptions")
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

// tempSubscriptionGroupSignupFailureData is a temporary struct used for validating the fields of SubscriptionGroupSignupFailureData.
type tempSubscriptionGroupSignupFailureData  struct {
    PayerId                 *int                          `json:"payer_id,omitempty"`
    PayerReference          *string                       `json:"payer_reference,omitempty"`
    PaymentProfileId        *int                          `json:"payment_profile_id,omitempty"`
    PaymentCollectionMethod *string                       `json:"payment_collection_method,omitempty"`
    PayerAttributes         *PayerAttributes              `json:"payer_attributes,omitempty"`
    CreditCardAttributes    *SubscriptionGroupCreditCard  `json:"credit_card_attributes,omitempty"`
    BankAccountAttributes   *SubscriptionGroupBankAccount `json:"bank_account_attributes,omitempty"`
    Subscriptions           []SubscriptionGroupSignupItem `json:"subscriptions,omitempty"`
}
