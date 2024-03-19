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

// CustomerChangePayer represents a CustomerChangePayer struct.
// This is a container for one-of cases.
type CustomerChangePayer struct {
	value                 any
	isCustomerPayerChange bool
}

// String converts the CustomerChangePayer object to a string representation.
func (c CustomerChangePayer) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for CustomerChangePayer.
// It customizes the JSON marshaling process for CustomerChangePayer objects.
func (c *CustomerChangePayer) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.CustomerChangePayerContainer.From*` functions to initialize the CustomerChangePayer object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CustomerChangePayer object to a map representation for JSON marshaling.
func (c *CustomerChangePayer) toMap() any {
	switch obj := c.value.(type) {
	case *CustomerPayerChange:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerChangePayer.
// It customizes the JSON unmarshaling process for CustomerChangePayer objects.
func (c *CustomerChangePayer) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&CustomerPayerChange{}, false, &c.isCustomerPayerChange),
	)

	c.value = result
	return err
}

func (c *CustomerChangePayer) AsCustomerPayerChange() (
	*CustomerPayerChange,
	bool) {
	if !c.isCustomerPayerChange {
		return nil, false
	}
	return c.value.(*CustomerPayerChange), true
}

// internalCustomerChangePayer represents a customerChangePayer struct.
// This is a container for one-of cases.
type internalCustomerChangePayer struct{}

var CustomerChangePayerContainer internalCustomerChangePayer

// The internalCustomerChangePayer instance, wrapping the provided CustomerPayerChange value.
func (c *internalCustomerChangePayer) FromCustomerPayerChange(val CustomerPayerChange) CustomerChangePayer {
	return CustomerChangePayer{value: &val}
}
