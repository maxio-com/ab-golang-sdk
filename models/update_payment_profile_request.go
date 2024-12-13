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

// UpdatePaymentProfileRequest represents a UpdatePaymentProfileRequest struct.
type UpdatePaymentProfileRequest struct {
    PaymentProfile       UpdatePaymentProfile   `json:"payment_profile"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdatePaymentProfileRequest.
// It customizes the JSON marshaling process for UpdatePaymentProfileRequest objects.
func (u UpdatePaymentProfileRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "payment_profile"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdatePaymentProfileRequest object to a map representation for JSON marshaling.
func (u UpdatePaymentProfileRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["payment_profile"] = u.PaymentProfile.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdatePaymentProfileRequest.
// It customizes the JSON unmarshaling process for UpdatePaymentProfileRequest objects.
func (u *UpdatePaymentProfileRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdatePaymentProfileRequest
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
    u.AdditionalProperties = additionalProperties
    
    u.PaymentProfile = *temp.PaymentProfile
    return nil
}

// tempUpdatePaymentProfileRequest is a temporary struct used for validating the fields of UpdatePaymentProfileRequest.
type tempUpdatePaymentProfileRequest  struct {
    PaymentProfile *UpdatePaymentProfile `json:"payment_profile"`
}

func (u *tempUpdatePaymentProfileRequest) validate() error {
    var errs []string
    if u.PaymentProfile == nil {
        errs = append(errs, "required field `payment_profile` is missing for type `Update Payment Profile Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
