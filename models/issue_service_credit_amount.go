// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// IssueServiceCreditAmount represents a IssueServiceCreditAmount struct.
// This is a container for one-of cases.
type IssueServiceCreditAmount struct {
	value       any
	isPrecision bool
	isString    bool
}

// String implements the fmt.Stringer interface for IssueServiceCreditAmount,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IssueServiceCreditAmount) String() string {
	return fmt.Sprintf("%v", i.value)
}

// MarshalJSON implements the json.Marshaler interface for IssueServiceCreditAmount.
// It customizes the JSON marshaling process for IssueServiceCreditAmount objects.
func (i IssueServiceCreditAmount) MarshalJSON() (
	[]byte,
	error) {
	if i.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.IssueServiceCreditAmountContainer.From*` functions to initialize the IssueServiceCreditAmount object.")
	}
	return json.Marshal(i.toMap())
}

// toMap converts the IssueServiceCreditAmount object to a map representation for JSON marshaling.
func (i *IssueServiceCreditAmount) toMap() any {
	switch obj := i.value.(type) {
	case *float64:
		return *obj
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssueServiceCreditAmount.
// It customizes the JSON unmarshaling process for IssueServiceCreditAmount objects.
func (i *IssueServiceCreditAmount) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(float64), false, &i.isPrecision),
		NewTypeHolder(new(string), false, &i.isString),
	)

	i.value = result
	return err
}

func (i *IssueServiceCreditAmount) AsPrecision() (
	*float64,
	bool) {
	if !i.isPrecision {
		return nil, false
	}
	return i.value.(*float64), true
}

func (i *IssueServiceCreditAmount) AsString() (
	*string,
	bool) {
	if !i.isString {
		return nil, false
	}
	return i.value.(*string), true
}

// internalIssueServiceCreditAmount represents a issueServiceCreditAmount struct.
// This is a container for one-of cases.
type internalIssueServiceCreditAmount struct{}

var IssueServiceCreditAmountContainer internalIssueServiceCreditAmount

// The internalIssueServiceCreditAmount instance, wrapping the provided float64 value.
func (i *internalIssueServiceCreditAmount) FromPrecision(val float64) IssueServiceCreditAmount {
	return IssueServiceCreditAmount{value: &val}
}

// The internalIssueServiceCreditAmount instance, wrapping the provided string value.
func (i *internalIssueServiceCreditAmount) FromString(val string) IssueServiceCreditAmount {
	return IssueServiceCreditAmount{value: &val}
}
