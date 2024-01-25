package models

import (
	"encoding/json"
)

// BankAccountVerificationRequest represents a BankAccountVerificationRequest struct.
type BankAccountVerificationRequest struct {
	BankAccountVerification BankAccountVerification `json:"bank_account_verification"`
}

// MarshalJSON implements the json.Marshaler interface for BankAccountVerificationRequest.
// It customizes the JSON marshaling process for BankAccountVerificationRequest objects.
func (b *BankAccountVerificationRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BankAccountVerificationRequest object to a map representation for JSON marshaling.
func (b *BankAccountVerificationRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["bank_account_verification"] = b.BankAccountVerification
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BankAccountVerificationRequest.
// It customizes the JSON unmarshaling process for BankAccountVerificationRequest objects.
func (b *BankAccountVerificationRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		BankAccountVerification BankAccountVerification `json:"bank_account_verification"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	b.BankAccountVerification = temp.BankAccountVerification
	return nil
}
