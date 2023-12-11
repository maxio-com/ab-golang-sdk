package models

import (
	"encoding/json"
)

// BankAccountResponse represents a BankAccountResponse struct.
type BankAccountResponse struct {
	PaymentProfile BankAccount `json:"payment_profile"`
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
	structMap["payment_profile"] = b.PaymentProfile
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BankAccountResponse.
// It customizes the JSON unmarshaling process for BankAccountResponse objects.
func (b *BankAccountResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PaymentProfile BankAccount `json:"payment_profile"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.PaymentProfile = temp.PaymentProfile
	return nil
}
