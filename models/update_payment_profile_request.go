package models

import (
    "encoding/json"
)

// UpdatePaymentProfileRequest represents a UpdatePaymentProfileRequest struct.
type UpdatePaymentProfileRequest struct {
    PaymentProfile UpdatePaymentProfile `json:"payment_profile"`
}

// MarshalJSON implements the json.Marshaler interface for UpdatePaymentProfileRequest.
// It customizes the JSON marshaling process for UpdatePaymentProfileRequest objects.
func (u *UpdatePaymentProfileRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdatePaymentProfileRequest object to a map representation for JSON marshaling.
func (u *UpdatePaymentProfileRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["payment_profile"] = u.PaymentProfile
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdatePaymentProfileRequest.
// It customizes the JSON unmarshaling process for UpdatePaymentProfileRequest objects.
func (u *UpdatePaymentProfileRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PaymentProfile UpdatePaymentProfile `json:"payment_profile"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.PaymentProfile = temp.PaymentProfile
    return nil
}
