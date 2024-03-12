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

// CreateMultiInvoicePaymentAmount represents a CreateMultiInvoicePaymentAmount struct.
// This is a container for one-of cases.
type CreateMultiInvoicePaymentAmount struct {
	value       any
	isString    bool
	isPrecision bool
}

// String converts the CreateMultiInvoicePaymentAmount object to a string representation.
func (c CreateMultiInvoicePaymentAmount) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateMultiInvoicePaymentAmount.
// It customizes the JSON marshaling process for CreateMultiInvoicePaymentAmount objects.
func (c *CreateMultiInvoicePaymentAmount) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.CreateMultiInvoicePaymentAmountContainer.From*` functions to initialize the CreateMultiInvoicePaymentAmount object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreateMultiInvoicePaymentAmount object to a map representation for JSON marshaling.
func (c *CreateMultiInvoicePaymentAmount) toMap() any {
	switch obj := c.value.(type) {
	case *string:
		return *obj
	case *float64:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMultiInvoicePaymentAmount.
// It customizes the JSON unmarshaling process for CreateMultiInvoicePaymentAmount objects.
func (c *CreateMultiInvoicePaymentAmount) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(string), false, &c.isString),
		NewTypeHolder(new(float64), false, &c.isPrecision),
	)

	c.value = result
	return err
}

func (c *CreateMultiInvoicePaymentAmount) AsString() (
	*string,
	bool) {
	if !c.isString {
		return nil, false
	}
	return c.value.(*string), true
}

func (c *CreateMultiInvoicePaymentAmount) AsPrecision() (
	*float64,
	bool) {
	if !c.isPrecision {
		return nil, false
	}
	return c.value.(*float64), true
}

// internalCreateMultiInvoicePaymentAmount represents a createMultiInvoicePaymentAmount struct.
// This is a container for one-of cases.
type internalCreateMultiInvoicePaymentAmount struct{}

var CreateMultiInvoicePaymentAmountContainer internalCreateMultiInvoicePaymentAmount

func (c *internalCreateMultiInvoicePaymentAmount) FromString(val string) CreateMultiInvoicePaymentAmount {
	return CreateMultiInvoicePaymentAmount{value: &val}
}

func (c *internalCreateMultiInvoicePaymentAmount) FromPrecision(val float64) CreateMultiInvoicePaymentAmount {
	return CreateMultiInvoicePaymentAmount{value: &val}
}
