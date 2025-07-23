// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// BankAccountVerification represents a BankAccountVerification struct.
type BankAccountVerification struct {
    Deposit1InCents      *int64                 `json:"deposit_1_in_cents,omitempty"`
    Deposit2InCents      *int64                 `json:"deposit_2_in_cents,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for BankAccountVerification,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b BankAccountVerification) String() string {
    return fmt.Sprintf(
    	"BankAccountVerification[Deposit1InCents=%v, Deposit2InCents=%v, AdditionalProperties=%v]",
    	b.Deposit1InCents, b.Deposit2InCents, b.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for BankAccountVerification.
// It customizes the JSON marshaling process for BankAccountVerification objects.
func (b BankAccountVerification) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "deposit_1_in_cents", "deposit_2_in_cents"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BankAccountVerification object to a map representation for JSON marshaling.
func (b BankAccountVerification) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
    if b.Deposit1InCents != nil {
        structMap["deposit_1_in_cents"] = b.Deposit1InCents
    }
    if b.Deposit2InCents != nil {
        structMap["deposit_2_in_cents"] = b.Deposit2InCents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BankAccountVerification.
// It customizes the JSON unmarshaling process for BankAccountVerification objects.
func (b *BankAccountVerification) UnmarshalJSON(input []byte) error {
    var temp tempBankAccountVerification
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "deposit_1_in_cents", "deposit_2_in_cents")
    if err != nil {
    	return err
    }
    b.AdditionalProperties = additionalProperties
    
    b.Deposit1InCents = temp.Deposit1InCents
    b.Deposit2InCents = temp.Deposit2InCents
    return nil
}

// tempBankAccountVerification is a temporary struct used for validating the fields of BankAccountVerification.
type tempBankAccountVerification  struct {
    Deposit1InCents *int64 `json:"deposit_1_in_cents,omitempty"`
    Deposit2InCents *int64 `json:"deposit_2_in_cents,omitempty"`
}
