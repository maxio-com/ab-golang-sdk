package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreatePaymentProfileRequest represents a CreatePaymentProfileRequest struct.
type CreatePaymentProfileRequest struct {
    PaymentProfile       CreatePaymentProfile `json:"payment_profile"`
    AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePaymentProfileRequest.
// It customizes the JSON marshaling process for CreatePaymentProfileRequest objects.
func (c CreatePaymentProfileRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePaymentProfileRequest object to a map representation for JSON marshaling.
func (c CreatePaymentProfileRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["payment_profile"] = c.PaymentProfile.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePaymentProfileRequest.
// It customizes the JSON unmarshaling process for CreatePaymentProfileRequest objects.
func (c *CreatePaymentProfileRequest) UnmarshalJSON(input []byte) error {
    var temp createPaymentProfileRequest
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
    
    c.AdditionalProperties = additionalProperties
    c.PaymentProfile = *temp.PaymentProfile
    return nil
}

// TODO
type createPaymentProfileRequest  struct {
    PaymentProfile *CreatePaymentProfile `json:"payment_profile"`
}

func (c *createPaymentProfileRequest) validate() error {
    var errs []string
    if c.PaymentProfile == nil {
        errs = append(errs, "required field `payment_profile` is missing for type `Create Payment Profile Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
