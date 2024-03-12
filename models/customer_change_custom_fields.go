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

// CustomerChangeCustomFields represents a CustomerChangeCustomFields struct.
// This is a container for one-of cases.
type CustomerChangeCustomFields struct {
	value                        any
	isCustomerCustomFieldsChange bool
}

// String converts the CustomerChangeCustomFields object to a string representation.
func (c CustomerChangeCustomFields) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for CustomerChangeCustomFields.
// It customizes the JSON marshaling process for CustomerChangeCustomFields objects.
func (c *CustomerChangeCustomFields) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.CustomerChangeCustomFieldsContainer.From*` functions to initialize the CustomerChangeCustomFields object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CustomerChangeCustomFields object to a map representation for JSON marshaling.
func (c *CustomerChangeCustomFields) toMap() any {
	switch obj := c.value.(type) {
	case *CustomerCustomFieldsChange:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerChangeCustomFields.
// It customizes the JSON unmarshaling process for CustomerChangeCustomFields objects.
func (c *CustomerChangeCustomFields) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&CustomerCustomFieldsChange{}, false, &c.isCustomerCustomFieldsChange),
	)

	c.value = result
	return err
}

func (c *CustomerChangeCustomFields) AsCustomerCustomFieldsChange() (
	*CustomerCustomFieldsChange,
	bool) {
	if !c.isCustomerCustomFieldsChange {
		return nil, false
	}
	return c.value.(*CustomerCustomFieldsChange), true
}

// internalCustomerChangeCustomFields represents a customerChangeCustomFields struct.
// This is a container for one-of cases.
type internalCustomerChangeCustomFields struct{}

var CustomerChangeCustomFieldsContainer internalCustomerChangeCustomFields

func (c *internalCustomerChangeCustomFields) FromCustomerCustomFieldsChange(val CustomerCustomFieldsChange) CustomerChangeCustomFields {
	return CustomerChangeCustomFields{value: &val}
}
