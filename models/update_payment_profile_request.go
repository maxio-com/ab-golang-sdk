package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UpdatePaymentProfileRequest represents a UpdatePaymentProfileRequest struct.
type UpdatePaymentProfileRequest struct {
    PaymentProfile       UpdatePaymentProfile `json:"payment_profile"`
    AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdatePaymentProfileRequest.
// It customizes the JSON marshaling process for UpdatePaymentProfileRequest objects.
func (u UpdatePaymentProfileRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdatePaymentProfileRequest object to a map representation for JSON marshaling.
func (u UpdatePaymentProfileRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["payment_profile"] = u.PaymentProfile.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdatePaymentProfileRequest.
// It customizes the JSON unmarshaling process for UpdatePaymentProfileRequest objects.
func (u *UpdatePaymentProfileRequest) UnmarshalJSON(input []byte) error {
    var temp updatePaymentProfileRequest
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
    
    u.AdditionalProperties = additionalProperties
    u.PaymentProfile = *temp.PaymentProfile
    return nil
}

// TODO
type updatePaymentProfileRequest  struct {
    PaymentProfile *UpdatePaymentProfile `json:"payment_profile"`
}

func (u *updatePaymentProfileRequest) validate() error {
    var errs []string
    if u.PaymentProfile == nil {
        errs = append(errs, "required field `payment_profile` is missing for type `Update Payment Profile Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
