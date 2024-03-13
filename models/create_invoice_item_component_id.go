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

// CreateInvoiceItemComponentId represents a CreateInvoiceItemComponentId struct.
// This is a container for one-of cases.
type CreateInvoiceItemComponentId struct {
	value    any
	isString bool
	isNumber bool
}

// String converts the CreateInvoiceItemComponentId object to a string representation.
func (c CreateInvoiceItemComponentId) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceItemComponentId.
// It customizes the JSON marshaling process for CreateInvoiceItemComponentId objects.
func (c *CreateInvoiceItemComponentId) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.CreateInvoiceItemComponentIdContainer.From*` functions to initialize the CreateInvoiceItemComponentId object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceItemComponentId object to a map representation for JSON marshaling.
func (c *CreateInvoiceItemComponentId) toMap() any {
	switch obj := c.value.(type) {
	case *string:
		return *obj
	case *int:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceItemComponentId.
// It customizes the JSON unmarshaling process for CreateInvoiceItemComponentId objects.
func (c *CreateInvoiceItemComponentId) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &c.isString),
		NewTypeHolder(new(int), false, &c.isNumber),
	)

	c.value = result
	return err
}

func (c *CreateInvoiceItemComponentId) AsString() (
	*string,
	bool) {
	if !c.isString {
		return nil, false
	}
	return c.value.(*string), true
}

func (c *CreateInvoiceItemComponentId) AsNumber() (
	*int,
	bool) {
	if !c.isNumber {
		return nil, false
	}
	return c.value.(*int), true
}

// internalCreateInvoiceItemComponentId represents a createInvoiceItemComponentId struct.
// This is a container for one-of cases.
type internalCreateInvoiceItemComponentId struct{}

var CreateInvoiceItemComponentIdContainer internalCreateInvoiceItemComponentId

func (c *internalCreateInvoiceItemComponentId) FromString(val string) CreateInvoiceItemComponentId {
	return CreateInvoiceItemComponentId{value: &val}
}

func (c *internalCreateInvoiceItemComponentId) FromNumber(val int) CreateInvoiceItemComponentId {
	return CreateInvoiceItemComponentId{value: &val}
}
