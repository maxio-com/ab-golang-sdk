// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// PaymentProfileResponse represents a PaymentProfileResponse struct.
type PaymentProfileResponse struct {
	PaymentProfile       PaymentProfile         `json:"payment_profile"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PaymentProfileResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PaymentProfileResponse) String() string {
	return fmt.Sprintf(
		"PaymentProfileResponse[PaymentProfile=%v, AdditionalProperties=%v]",
		p.PaymentProfile, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PaymentProfileResponse.
// It customizes the JSON marshaling process for PaymentProfileResponse objects.
func (p PaymentProfileResponse) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(p.AdditionalProperties,
		"payment_profile"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(p.toMap())
}

// toMap converts the PaymentProfileResponse object to a map representation for JSON marshaling.
func (p PaymentProfileResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, p.AdditionalProperties)
	structMap["payment_profile"] = p.PaymentProfile.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentProfileResponse.
// It customizes the JSON unmarshaling process for PaymentProfileResponse objects.
func (p *PaymentProfileResponse) UnmarshalJSON(input []byte) error {
	var temp tempPaymentProfileResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "payment_profile")
	if err != nil {
		return err
	}
	p.AdditionalProperties = additionalProperties

	p.PaymentProfile = *temp.PaymentProfile
	return nil
}

// tempPaymentProfileResponse is a temporary struct used for validating the fields of PaymentProfileResponse.
type tempPaymentProfileResponse struct {
	PaymentProfile *PaymentProfile `json:"payment_profile"`
}

func (p *tempPaymentProfileResponse) validate() error {
	var errs []string
	if p.PaymentProfile == nil {
		errs = append(errs, "required field `payment_profile` is missing for type `Payment Profile Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
