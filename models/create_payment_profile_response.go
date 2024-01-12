package models

import (
    "encoding/json"
)

// CreatePaymentProfileResponse represents a CreatePaymentProfileResponse struct.
type CreatePaymentProfileResponse struct {
    PaymentProfile CreatedPaymentProfile `json:"payment_profile"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePaymentProfileResponse.
// It customizes the JSON marshaling process for CreatePaymentProfileResponse objects.
func (c *CreatePaymentProfileResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePaymentProfileResponse object to a map representation for JSON marshaling.
func (c *CreatePaymentProfileResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["payment_profile"] = c.PaymentProfile
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePaymentProfileResponse.
// It customizes the JSON unmarshaling process for CreatePaymentProfileResponse objects.
func (c *CreatePaymentProfileResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PaymentProfile CreatedPaymentProfile `json:"payment_profile"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.PaymentProfile = temp.PaymentProfile
    return nil
}
