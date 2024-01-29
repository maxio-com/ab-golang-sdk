package models

import (
	"encoding/json"
)

// PaymentProfileResponse represents a PaymentProfileResponse struct.
type PaymentProfileResponse struct {
	PaymentProfile interface{} `json:"payment_profile"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentProfileResponse.
// It customizes the JSON marshaling process for PaymentProfileResponse objects.
func (p *PaymentProfileResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PaymentProfileResponse object to a map representation for JSON marshaling.
func (p *PaymentProfileResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["payment_profile"] = p.PaymentProfile
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentProfileResponse.
// It customizes the JSON unmarshaling process for PaymentProfileResponse objects.
func (p *PaymentProfileResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PaymentProfile interface{} `json:"payment_profile"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.PaymentProfile = temp.PaymentProfile
	return nil
}
