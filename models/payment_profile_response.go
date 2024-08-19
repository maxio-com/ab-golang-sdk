/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// PaymentProfileResponse represents a PaymentProfileResponse struct.
type PaymentProfileResponse struct {
    PaymentProfile       PaymentProfile `json:"payment_profile"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentProfileResponse.
// It customizes the JSON marshaling process for PaymentProfileResponse objects.
func (p PaymentProfileResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentProfileResponse object to a map representation for JSON marshaling.
func (p PaymentProfileResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "payment_profile")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.PaymentProfile = *temp.PaymentProfile
    return nil
}

// tempPaymentProfileResponse is a temporary struct used for validating the fields of PaymentProfileResponse.
type tempPaymentProfileResponse  struct {
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
    return errors.New(strings.Join (errs, "\n"))
}
