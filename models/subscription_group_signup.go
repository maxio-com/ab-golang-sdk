/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SubscriptionGroupSignup represents a SubscriptionGroupSignup struct.
type SubscriptionGroupSignup struct {
    PaymentProfileId        *int                          `json:"payment_profile_id,omitempty"`
    PayerId                 *int                          `json:"payer_id,omitempty"`
    PayerReference          *string                       `json:"payer_reference,omitempty"`
    // The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`.
    PaymentCollectionMethod *CollectionMethod             `json:"payment_collection_method,omitempty"`
    PayerAttributes         *PayerAttributes              `json:"payer_attributes,omitempty"`
    CreditCardAttributes    *SubscriptionGroupCreditCard  `json:"credit_card_attributes,omitempty"`
    BankAccountAttributes   *SubscriptionGroupBankAccount `json:"bank_account_attributes,omitempty"`
    Subscriptions           []SubscriptionGroupSignupItem `json:"subscriptions"`
    AdditionalProperties    map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionGroupSignup,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroupSignup) String() string {
    return fmt.Sprintf(
    	"SubscriptionGroupSignup[PaymentProfileId=%v, PayerId=%v, PayerReference=%v, PaymentCollectionMethod=%v, PayerAttributes=%v, CreditCardAttributes=%v, BankAccountAttributes=%v, Subscriptions=%v, AdditionalProperties=%v]",
    	s.PaymentProfileId, s.PayerId, s.PayerReference, s.PaymentCollectionMethod, s.PayerAttributes, s.CreditCardAttributes, s.BankAccountAttributes, s.Subscriptions, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignup.
// It customizes the JSON marshaling process for SubscriptionGroupSignup objects.
func (s SubscriptionGroupSignup) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "payment_profile_id", "payer_id", "payer_reference", "payment_collection_method", "payer_attributes", "credit_card_attributes", "bank_account_attributes", "subscriptions"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignup object to a map representation for JSON marshaling.
func (s SubscriptionGroupSignup) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
        structMap["payer_attributes"] = s.PayerAttributes.toMap()
    }
    if s.CreditCardAttributes != nil {
        structMap["credit_card_attributes"] = s.CreditCardAttributes.toMap()
    }
    if s.BankAccountAttributes != nil {
        structMap["bank_account_attributes"] = s.BankAccountAttributes.toMap()
    }
    structMap["subscriptions"] = s.Subscriptions
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignup.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignup objects.
func (s *SubscriptionGroupSignup) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionGroupSignup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "payment_profile_id", "payer_id", "payer_reference", "payment_collection_method", "payer_attributes", "credit_card_attributes", "bank_account_attributes", "subscriptions")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.PaymentProfileId = temp.PaymentProfileId
    s.PayerId = temp.PayerId
    s.PayerReference = temp.PayerReference
    s.PaymentCollectionMethod = temp.PaymentCollectionMethod
    s.PayerAttributes = temp.PayerAttributes
    s.CreditCardAttributes = temp.CreditCardAttributes
    s.BankAccountAttributes = temp.BankAccountAttributes
    s.Subscriptions = *temp.Subscriptions
    return nil
}

// tempSubscriptionGroupSignup is a temporary struct used for validating the fields of SubscriptionGroupSignup.
type tempSubscriptionGroupSignup  struct {
    PaymentProfileId        *int                           `json:"payment_profile_id,omitempty"`
    PayerId                 *int                           `json:"payer_id,omitempty"`
    PayerReference          *string                        `json:"payer_reference,omitempty"`
    PaymentCollectionMethod *CollectionMethod              `json:"payment_collection_method,omitempty"`
    PayerAttributes         *PayerAttributes               `json:"payer_attributes,omitempty"`
    CreditCardAttributes    *SubscriptionGroupCreditCard   `json:"credit_card_attributes,omitempty"`
    BankAccountAttributes   *SubscriptionGroupBankAccount  `json:"bank_account_attributes,omitempty"`
    Subscriptions           *[]SubscriptionGroupSignupItem `json:"subscriptions"`
}

func (s *tempSubscriptionGroupSignup) validate() error {
    var errs []string
    if s.Subscriptions == nil {
        errs = append(errs, "required field `subscriptions` is missing for type `Subscription Group Signup`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
