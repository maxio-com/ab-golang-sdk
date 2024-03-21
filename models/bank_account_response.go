package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// BankAccountResponse represents a BankAccountResponse struct.
type BankAccountResponse struct {
	PaymentProfile BankAccountPaymentProfile `json:"payment_profile"`
}

// MarshalJSON implements the json.Marshaler interface for BankAccountResponse.
// It customizes the JSON marshaling process for BankAccountResponse objects.
func (b *BankAccountResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BankAccountResponse object to a map representation for JSON marshaling.
func (b *BankAccountResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["payment_profile"] = b.PaymentProfile.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BankAccountResponse.
// It customizes the JSON unmarshaling process for BankAccountResponse objects.
func (b *BankAccountResponse) UnmarshalJSON(input []byte) error {
	var temp bankAccountResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	b.PaymentProfile = *temp.PaymentProfile
	return nil
}

// TODO
type bankAccountResponse struct {
	PaymentProfile *BankAccountPaymentProfile `json:"payment_profile"`
}

func (b *bankAccountResponse) validate() error {
	var errs []string
	if b.PaymentProfile == nil {
		errs = append(errs, "required field `payment_profile` is missing for type `Bank Account Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
