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

// CustomerChangeShippingAddress represents a CustomerChangeShippingAddress struct.
// This is a container for one-of cases.
type CustomerChangeShippingAddress struct {
	value           any
	isAddressChange bool
}

// String converts the CustomerChangeShippingAddress object to a string representation.
func (c CustomerChangeShippingAddress) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for CustomerChangeShippingAddress.
// It customizes the JSON marshaling process for CustomerChangeShippingAddress objects.
func (c *CustomerChangeShippingAddress) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.CustomerChangeShippingAddressContainer.From*` functions to initialize the CustomerChangeShippingAddress object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CustomerChangeShippingAddress object to a map representation for JSON marshaling.
func (c *CustomerChangeShippingAddress) toMap() any {
	switch obj := c.value.(type) {
	case *AddressChange:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerChangeShippingAddress.
// It customizes the JSON unmarshaling process for CustomerChangeShippingAddress objects.
func (c *CustomerChangeShippingAddress) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&AddressChange{}, false, &c.isAddressChange),
	)

	c.value = result
	return err
}

func (c *CustomerChangeShippingAddress) AsAddressChange() (
	*AddressChange,
	bool) {
	if !c.isAddressChange {
		return nil, false
	}
	return c.value.(*AddressChange), true
}

// internalCustomerChangeShippingAddress represents a customerChangeShippingAddress struct.
// This is a container for one-of cases.
type internalCustomerChangeShippingAddress struct{}

var CustomerChangeShippingAddressContainer internalCustomerChangeShippingAddress

// The internalCustomerChangeShippingAddress instance, wrapping the provided AddressChange value.
func (c *internalCustomerChangeShippingAddress) FromAddressChange(val AddressChange) CustomerChangeShippingAddress {
	return CustomerChangeShippingAddress{value: &val}
}
