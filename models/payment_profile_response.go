package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// PaymentProfileResponse represents a PaymentProfileResponse struct.
type PaymentProfileResponse struct {
	PaymentProfile PaymentProfileResponsePaymentProfile `json:"payment_profile"`
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
	structMap["payment_profile"] = p.PaymentProfile.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentProfileResponse.
// It customizes the JSON unmarshaling process for PaymentProfileResponse objects.
func (p *PaymentProfileResponse) UnmarshalJSON(input []byte) error {
	var temp paymentProfileResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	p.PaymentProfile = *temp.PaymentProfile
	return nil
}

// TODO
type paymentProfileResponse struct {
	PaymentProfile *PaymentProfileResponsePaymentProfile `json:"payment_profile"`
}

func (p *paymentProfileResponse) validate() error {
	var errs []string
	if p.PaymentProfile == nil {
		errs = append(errs, "required field `payment_profile` is missing for type `Payment Profile Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
