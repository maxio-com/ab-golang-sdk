/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// DeductServiceCreditAmount represents a DeductServiceCreditAmount struct.
// This is a container for one-of cases.
type DeductServiceCreditAmount struct {
	value       any
	isString    bool
	isPrecision bool
}

// String converts the DeductServiceCreditAmount object to a string representation.
func (d DeductServiceCreditAmount) String() string {
	if bytes, err := json.Marshal(d.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for DeductServiceCreditAmount.
// It customizes the JSON marshaling process for DeductServiceCreditAmount objects.
func (d *DeductServiceCreditAmount) MarshalJSON() (
	[]byte,
	error) {
	if d.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.DeductServiceCreditAmountContainer.From*` functions to initialize the DeductServiceCreditAmount object.")
	}
	return json.Marshal(d.toMap())
}

// toMap converts the DeductServiceCreditAmount object to a map representation for JSON marshaling.
func (d *DeductServiceCreditAmount) toMap() any {
	switch obj := d.value.(type) {
	case *string:
		return *obj
	case *float64:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeductServiceCreditAmount.
// It customizes the JSON unmarshaling process for DeductServiceCreditAmount objects.
func (d *DeductServiceCreditAmount) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &d.isString),
		NewTypeHolder(new(float64), false, &d.isPrecision),
	)

	d.value = result
	return err
}

func (d *DeductServiceCreditAmount) AsString() (
	*string,
	bool) {
	if !d.isString {
		return nil, false
	}
	return d.value.(*string), true
}

func (d *DeductServiceCreditAmount) AsPrecision() (
	*float64,
	bool) {
	if !d.isPrecision {
		return nil, false
	}
	return d.value.(*float64), true
}

// internalDeductServiceCreditAmount represents a deductServiceCreditAmount struct.
// This is a container for one-of cases.
type internalDeductServiceCreditAmount struct{}

var DeductServiceCreditAmountContainer internalDeductServiceCreditAmount

func (d *internalDeductServiceCreditAmount) FromString(val string) DeductServiceCreditAmount {
	return DeductServiceCreditAmount{value: &val}
}

func (d *internalDeductServiceCreditAmount) FromPrecision(val float64) DeductServiceCreditAmount {
	return DeductServiceCreditAmount{value: &val}
}
