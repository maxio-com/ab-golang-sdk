package models

import (
	"encoding/json"
)

// AgreementAcceptance represents a AgreementAcceptance struct.
// Required when creating a subscription with Maxio Payments.
type AgreementAcceptance struct {
	// Required when providing agreement acceptance params.
	IpAddress *string `json:"ip_address,omitempty"`
	// Required when creating a subscription with Maxio Payments. Either terms_url or provacy_policy_url required when providing agreement_acceptance params.
	TermsUrl                *string `json:"terms_url,omitempty"`
	PrivacyPolicyUrl        *string `json:"privacy_policy_url,omitempty"`
	ReturnRefundPolicyUrl   *string `json:"return_refund_policy_url,omitempty"`
	DeliveryPolicyUrl       *string `json:"delivery_policy_url,omitempty"`
	SecureCheckoutPolicyUrl *string `json:"secure_checkout_policy_url,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AgreementAcceptance.
// It customizes the JSON marshaling process for AgreementAcceptance objects.
func (a *AgreementAcceptance) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AgreementAcceptance object to a map representation for JSON marshaling.
func (a *AgreementAcceptance) toMap() map[string]any {
	structMap := make(map[string]any)
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
	temp := &struct {
		IpAddress               *string `json:"ip_address,omitempty"`
		TermsUrl                *string `json:"terms_url,omitempty"`
		PrivacyPolicyUrl        *string `json:"privacy_policy_url,omitempty"`
		ReturnRefundPolicyUrl   *string `json:"return_refund_policy_url,omitempty"`
		DeliveryPolicyUrl       *string `json:"delivery_policy_url,omitempty"`
		SecureCheckoutPolicyUrl *string `json:"secure_checkout_policy_url,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.IpAddress = temp.IpAddress
	a.TermsUrl = temp.TermsUrl
	a.PrivacyPolicyUrl = temp.PrivacyPolicyUrl
	a.ReturnRefundPolicyUrl = temp.ReturnRefundPolicyUrl
	a.DeliveryPolicyUrl = temp.DeliveryPolicyUrl
	a.SecureCheckoutPolicyUrl = temp.SecureCheckoutPolicyUrl
	return nil
}
