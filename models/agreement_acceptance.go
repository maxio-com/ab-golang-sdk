// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// AgreementAcceptance represents a AgreementAcceptance struct.
// Required when creating a subscription with Maxio Payments.
type AgreementAcceptance struct {
	// Required when providing agreement acceptance params.
	IpAddress *string `json:"ip_address,omitempty"`
	// Required when creating a subscription with Maxio Payments. Either terms_url or provacy_policy_url required when providing agreement_acceptance params.
	TermsUrl                *string                `json:"terms_url,omitempty"`
	PrivacyPolicyUrl        *string                `json:"privacy_policy_url,omitempty"`
	ReturnRefundPolicyUrl   *string                `json:"return_refund_policy_url,omitempty"`
	DeliveryPolicyUrl       *string                `json:"delivery_policy_url,omitempty"`
	SecureCheckoutPolicyUrl *string                `json:"secure_checkout_policy_url,omitempty"`
	AdditionalProperties    map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AgreementAcceptance,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AgreementAcceptance) String() string {
	return fmt.Sprintf(
		"AgreementAcceptance[IpAddress=%v, TermsUrl=%v, PrivacyPolicyUrl=%v, ReturnRefundPolicyUrl=%v, DeliveryPolicyUrl=%v, SecureCheckoutPolicyUrl=%v, AdditionalProperties=%v]",
		a.IpAddress, a.TermsUrl, a.PrivacyPolicyUrl, a.ReturnRefundPolicyUrl, a.DeliveryPolicyUrl, a.SecureCheckoutPolicyUrl, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AgreementAcceptance.
// It customizes the JSON marshaling process for AgreementAcceptance objects.
func (a AgreementAcceptance) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"ip_address", "terms_url", "privacy_policy_url", "return_refund_policy_url", "delivery_policy_url", "secure_checkout_policy_url"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AgreementAcceptance object to a map representation for JSON marshaling.
func (a AgreementAcceptance) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.IpAddress != nil {
		structMap["ip_address"] = a.IpAddress
	}
	if a.TermsUrl != nil {
		structMap["terms_url"] = a.TermsUrl
	}
	if a.PrivacyPolicyUrl != nil {
		structMap["privacy_policy_url"] = a.PrivacyPolicyUrl
	}
	if a.ReturnRefundPolicyUrl != nil {
		structMap["return_refund_policy_url"] = a.ReturnRefundPolicyUrl
	}
	if a.DeliveryPolicyUrl != nil {
		structMap["delivery_policy_url"] = a.DeliveryPolicyUrl
	}
	if a.SecureCheckoutPolicyUrl != nil {
		structMap["secure_checkout_policy_url"] = a.SecureCheckoutPolicyUrl
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AgreementAcceptance.
// It customizes the JSON unmarshaling process for AgreementAcceptance objects.
func (a *AgreementAcceptance) UnmarshalJSON(input []byte) error {
	var temp tempAgreementAcceptance
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ip_address", "terms_url", "privacy_policy_url", "return_refund_policy_url", "delivery_policy_url", "secure_checkout_policy_url")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.IpAddress = temp.IpAddress
	a.TermsUrl = temp.TermsUrl
	a.PrivacyPolicyUrl = temp.PrivacyPolicyUrl
	a.ReturnRefundPolicyUrl = temp.ReturnRefundPolicyUrl
	a.DeliveryPolicyUrl = temp.DeliveryPolicyUrl
	a.SecureCheckoutPolicyUrl = temp.SecureCheckoutPolicyUrl
	return nil
}

// tempAgreementAcceptance is a temporary struct used for validating the fields of AgreementAcceptance.
type tempAgreementAcceptance struct {
	IpAddress               *string `json:"ip_address,omitempty"`
	TermsUrl                *string `json:"terms_url,omitempty"`
	PrivacyPolicyUrl        *string `json:"privacy_policy_url,omitempty"`
	ReturnRefundPolicyUrl   *string `json:"return_refund_policy_url,omitempty"`
	DeliveryPolicyUrl       *string `json:"delivery_policy_url,omitempty"`
	SecureCheckoutPolicyUrl *string `json:"secure_checkout_policy_url,omitempty"`
}
