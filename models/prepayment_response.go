// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// PrepaymentResponse represents a PrepaymentResponse struct.
type PrepaymentResponse struct {
	Prepayment           Prepayment             `json:"prepayment"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PrepaymentResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PrepaymentResponse) String() string {
	return fmt.Sprintf(
		"PrepaymentResponse[Prepayment=%v, AdditionalProperties=%v]",
		p.Prepayment, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PrepaymentResponse.
// It customizes the JSON marshaling process for PrepaymentResponse objects.
func (p PrepaymentResponse) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(p.AdditionalProperties,
		"prepayment"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(p.toMap())
}

// toMap converts the PrepaymentResponse object to a map representation for JSON marshaling.
func (p PrepaymentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, p.AdditionalProperties)
	structMap["prepayment"] = p.Prepayment.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaymentResponse.
// It customizes the JSON unmarshaling process for PrepaymentResponse objects.
func (p *PrepaymentResponse) UnmarshalJSON(input []byte) error {
	var temp tempPrepaymentResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "prepayment")
	if err != nil {
		return err
	}
	p.AdditionalProperties = additionalProperties

	p.Prepayment = *temp.Prepayment
	return nil
}

// tempPrepaymentResponse is a temporary struct used for validating the fields of PrepaymentResponse.
type tempPrepaymentResponse struct {
	Prepayment *Prepayment `json:"prepayment"`
}

func (p *tempPrepaymentResponse) validate() error {
	var errs []string
	if p.Prepayment == nil {
		errs = append(errs, "required field `prepayment` is missing for type `Prepayment Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
