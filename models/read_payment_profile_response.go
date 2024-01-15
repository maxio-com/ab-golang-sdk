package models

import (
    "encoding/json"
)

// ReadPaymentProfileResponse represents a ReadPaymentProfileResponse struct.
type ReadPaymentProfileResponse struct {
    PaymentProfile interface{} `json:"payment_profile"`
}

// MarshalJSON implements the json.Marshaler interface for ReadPaymentProfileResponse.
// It customizes the JSON marshaling process for ReadPaymentProfileResponse objects.
func (r *ReadPaymentProfileResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ReadPaymentProfileResponse object to a map representation for JSON marshaling.
func (r *ReadPaymentProfileResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["payment_profile"] = r.PaymentProfile
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReadPaymentProfileResponse.
// It customizes the JSON unmarshaling process for ReadPaymentProfileResponse objects.
func (r *ReadPaymentProfileResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PaymentProfile interface{} `json:"payment_profile"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    r.PaymentProfile = temp.PaymentProfile
    return nil
}
