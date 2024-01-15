package models

import (
    "encoding/json"
)

// UpdatePaymentProfileResponse represents a UpdatePaymentProfileResponse struct.
type UpdatePaymentProfileResponse struct {
    PaymentProfile UpdatedPaymentProfile `json:"payment_profile"`
}

// MarshalJSON implements the json.Marshaler interface for UpdatePaymentProfileResponse.
// It customizes the JSON marshaling process for UpdatePaymentProfileResponse objects.
func (u *UpdatePaymentProfileResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdatePaymentProfileResponse object to a map representation for JSON marshaling.
func (u *UpdatePaymentProfileResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["payment_profile"] = u.PaymentProfile
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdatePaymentProfileResponse.
// It customizes the JSON unmarshaling process for UpdatePaymentProfileResponse objects.
func (u *UpdatePaymentProfileResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PaymentProfile UpdatedPaymentProfile `json:"payment_profile"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.PaymentProfile = temp.PaymentProfile
    return nil
}
