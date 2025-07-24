// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// BankAccountResponse represents a BankAccountResponse struct.
type BankAccountResponse struct {
    PaymentProfile       BankAccountPaymentProfile `json:"payment_profile"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for BankAccountResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b BankAccountResponse) String() string {
    return fmt.Sprintf(
    	"BankAccountResponse[PaymentProfile=%v, AdditionalProperties=%v]",
    	b.PaymentProfile, b.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for BankAccountResponse.
// It customizes the JSON marshaling process for BankAccountResponse objects.
func (b BankAccountResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "payment_profile"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BankAccountResponse object to a map representation for JSON marshaling.
func (b BankAccountResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
    structMap["payment_profile"] = b.PaymentProfile.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BankAccountResponse.
// It customizes the JSON unmarshaling process for BankAccountResponse objects.
func (b *BankAccountResponse) UnmarshalJSON(input []byte) error {
    var temp tempBankAccountResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "payment_profile")
    if err != nil {
    	return err
    }
    b.AdditionalProperties = additionalProperties
    
    b.PaymentProfile = *temp.PaymentProfile
    return nil
}

// tempBankAccountResponse is a temporary struct used for validating the fields of BankAccountResponse.
type tempBankAccountResponse  struct {
    PaymentProfile *BankAccountPaymentProfile `json:"payment_profile"`
}

func (b *tempBankAccountResponse) validate() error {
    var errs []string
    if b.PaymentProfile == nil {
        errs = append(errs, "required field `payment_profile` is missing for type `Bank Account Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
