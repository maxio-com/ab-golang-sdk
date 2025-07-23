// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// DeductServiceCredit represents a DeductServiceCredit struct.
type DeductServiceCredit struct {
	Amount               DeductServiceCreditAmount `json:"amount"`
	Memo                 *string                   `json:"memo,omitempty"`
	AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for DeductServiceCredit,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DeductServiceCredit) String() string {
	return fmt.Sprintf(
		"DeductServiceCredit[Amount=%v, Memo=%v, AdditionalProperties=%v]",
		d.Amount, d.Memo, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DeductServiceCredit.
// It customizes the JSON marshaling process for DeductServiceCredit objects.
func (d DeductServiceCredit) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(d.AdditionalProperties,
		"amount", "memo"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(d.toMap())
}

// toMap converts the DeductServiceCredit object to a map representation for JSON marshaling.
func (d DeductServiceCredit) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, d.AdditionalProperties)
	structMap["amount"] = d.Amount.toMap()
	if d.Memo != nil {
		structMap["memo"] = d.Memo
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeductServiceCredit.
// It customizes the JSON unmarshaling process for DeductServiceCredit objects.
func (d *DeductServiceCredit) UnmarshalJSON(input []byte) error {
	var temp tempDeductServiceCredit
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "amount", "memo")
	if err != nil {
		return err
	}
	d.AdditionalProperties = additionalProperties

	d.Amount = *temp.Amount
	d.Memo = temp.Memo
	return nil
}

// tempDeductServiceCredit is a temporary struct used for validating the fields of DeductServiceCredit.
type tempDeductServiceCredit struct {
	Amount *DeductServiceCreditAmount `json:"amount"`
	Memo   *string                    `json:"memo,omitempty"`
}

func (d *tempDeductServiceCredit) validate() error {
	var errs []string
	if d.Amount == nil {
		errs = append(errs, "required field `amount` is missing for type `Deduct Service Credit`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
