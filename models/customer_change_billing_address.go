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

// CustomerChangeBillingAddress represents a CustomerChangeBillingAddress struct.
// This is a container for one-of cases.
type CustomerChangeBillingAddress struct {
	value           any
	isAddressChange bool
}

// String converts the CustomerChangeBillingAddress object to a string representation.
func (c CustomerChangeBillingAddress) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for CustomerChangeBillingAddress.
// It customizes the JSON marshaling process for CustomerChangeBillingAddress objects.
func (c *CustomerChangeBillingAddress) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.CustomerChangeBillingAddressContainer.From*` functions to initialize the CustomerChangeBillingAddress object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CustomerChangeBillingAddress object to a map representation for JSON marshaling.
func (c *CustomerChangeBillingAddress) toMap() any {
	switch obj := c.value.(type) {
	case *AddressChange:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerChangeBillingAddress.
// It customizes the JSON unmarshaling process for CustomerChangeBillingAddress objects.
func (c *CustomerChangeBillingAddress) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&AddressChange{}, false, &c.isAddressChange),
	)

	c.value = result
	return err
}

func (c *CustomerChangeBillingAddress) AsAddressChange() (
	*AddressChange,
	bool) {
	if !c.isAddressChange {
		return nil, false
	}
	return c.value.(*AddressChange), true
}

// internalCustomerChangeBillingAddress represents a customerChangeBillingAddress struct.
// This is a container for one-of cases.
type internalCustomerChangeBillingAddress struct{}

var CustomerChangeBillingAddressContainer internalCustomerChangeBillingAddress

func (c *internalCustomerChangeBillingAddress) FromAddressChange(val AddressChange) CustomerChangeBillingAddress {
	return CustomerChangeBillingAddress{value: &val}
}
