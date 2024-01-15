package models

import (
    "encoding/json"
)

// CreatePaymentProfileRequest represents a CreatePaymentProfileRequest struct.
type CreatePaymentProfileRequest struct {
    PaymentProfile CreatePaymentProfile `json:"payment_profile"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePaymentProfileRequest.
// It customizes the JSON marshaling process for CreatePaymentProfileRequest objects.
func (c *CreatePaymentProfileRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePaymentProfileRequest object to a map representation for JSON marshaling.
func (c *CreatePaymentProfileRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["payment_profile"] = c.PaymentProfile
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePaymentProfileRequest.
// It customizes the JSON unmarshaling process for CreatePaymentProfileRequest objects.
func (c *CreatePaymentProfileRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PaymentProfile CreatePaymentProfile `json:"payment_profile"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.PaymentProfile = temp.PaymentProfile
    return nil
}
