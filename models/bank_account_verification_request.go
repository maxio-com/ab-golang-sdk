/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// BankAccountVerificationRequest represents a BankAccountVerificationRequest struct.
type BankAccountVerificationRequest struct {
    BankAccountVerification BankAccountVerification `json:"bank_account_verification"`
    AdditionalProperties    map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for BankAccountVerificationRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b BankAccountVerificationRequest) String() string {
    return fmt.Sprintf(
    	"BankAccountVerificationRequest[BankAccountVerification=%v, AdditionalProperties=%v]",
    	b.BankAccountVerification, b.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for BankAccountVerificationRequest.
// It customizes the JSON marshaling process for BankAccountVerificationRequest objects.
func (b BankAccountVerificationRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "bank_account_verification"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BankAccountVerificationRequest object to a map representation for JSON marshaling.
func (b BankAccountVerificationRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
    structMap["bank_account_verification"] = b.BankAccountVerification.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BankAccountVerificationRequest.
// It customizes the JSON unmarshaling process for BankAccountVerificationRequest objects.
func (b *BankAccountVerificationRequest) UnmarshalJSON(input []byte) error {
    var temp tempBankAccountVerificationRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bank_account_verification")
    if err != nil {
    	return err
    }
    b.AdditionalProperties = additionalProperties
    
    b.BankAccountVerification = *temp.BankAccountVerification
    return nil
}

// tempBankAccountVerificationRequest is a temporary struct used for validating the fields of BankAccountVerificationRequest.
type tempBankAccountVerificationRequest  struct {
    BankAccountVerification *BankAccountVerification `json:"bank_account_verification"`
}

func (b *tempBankAccountVerificationRequest) validate() error {
    var errs []string
    if b.BankAccountVerification == nil {
        errs = append(errs, "required field `bank_account_verification` is missing for type `Bank Account Verification Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
