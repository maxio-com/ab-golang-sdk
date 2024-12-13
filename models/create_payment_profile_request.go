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

// CreatePaymentProfileRequest represents a CreatePaymentProfileRequest struct.
type CreatePaymentProfileRequest struct {
    PaymentProfile       CreatePaymentProfile   `json:"payment_profile"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePaymentProfileRequest.
// It customizes the JSON marshaling process for CreatePaymentProfileRequest objects.
func (c CreatePaymentProfileRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "payment_profile"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePaymentProfileRequest object to a map representation for JSON marshaling.
func (c CreatePaymentProfileRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["payment_profile"] = c.PaymentProfile.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePaymentProfileRequest.
// It customizes the JSON unmarshaling process for CreatePaymentProfileRequest objects.
func (c *CreatePaymentProfileRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreatePaymentProfileRequest
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
    c.AdditionalProperties = additionalProperties
    
    c.PaymentProfile = *temp.PaymentProfile
    return nil
}

// tempCreatePaymentProfileRequest is a temporary struct used for validating the fields of CreatePaymentProfileRequest.
type tempCreatePaymentProfileRequest  struct {
    PaymentProfile *CreatePaymentProfile `json:"payment_profile"`
}

func (c *tempCreatePaymentProfileRequest) validate() error {
    var errs []string
    if c.PaymentProfile == nil {
        errs = append(errs, "required field `payment_profile` is missing for type `Create Payment Profile Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
