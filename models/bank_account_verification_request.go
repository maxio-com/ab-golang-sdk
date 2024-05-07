package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// BankAccountVerificationRequest represents a BankAccountVerificationRequest struct.
type BankAccountVerificationRequest struct {
    BankAccountVerification BankAccountVerification `json:"bank_account_verification"`
    AdditionalProperties    map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BankAccountVerificationRequest.
// It customizes the JSON marshaling process for BankAccountVerificationRequest objects.
func (b BankAccountVerificationRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BankAccountVerificationRequest object to a map representation for JSON marshaling.
func (b BankAccountVerificationRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    structMap["bank_account_verification"] = b.BankAccountVerification.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BankAccountVerificationRequest.
// It customizes the JSON unmarshaling process for BankAccountVerificationRequest objects.
func (b *BankAccountVerificationRequest) UnmarshalJSON(input []byte) error {
    var temp bankAccountVerificationRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "bank_account_verification")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.BankAccountVerification = *temp.BankAccountVerification
    return nil
}

// bankAccountVerificationRequest is a temporary struct used for validating the fields of BankAccountVerificationRequest.
type bankAccountVerificationRequest  struct {
    BankAccountVerification *BankAccountVerification `json:"bank_account_verification"`
}

func (b *bankAccountVerificationRequest) validate() error {
    var errs []string
    if b.BankAccountVerification == nil {
        errs = append(errs, "required field `bank_account_verification` is missing for type `Bank Account Verification Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
