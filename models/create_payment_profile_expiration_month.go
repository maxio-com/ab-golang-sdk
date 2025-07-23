// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

// CreatePaymentProfileExpirationMonth represents a CreatePaymentProfileExpirationMonth struct.
// This is a container for one-of cases.
type CreatePaymentProfileExpirationMonth struct {
	value    any
	isNumber bool
	isString bool
}

// String implements the fmt.Stringer interface for CreatePaymentProfileExpirationMonth,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreatePaymentProfileExpirationMonth) String() string {
	return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CreatePaymentProfileExpirationMonth.
// It customizes the JSON marshaling process for CreatePaymentProfileExpirationMonth objects.
func (c CreatePaymentProfileExpirationMonth) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.CreatePaymentProfileExpirationMonthContainer.From*` functions to initialize the CreatePaymentProfileExpirationMonth object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CreatePaymentProfileExpirationMonth object to a map representation for JSON marshaling.
func (c *CreatePaymentProfileExpirationMonth) toMap() any {
	switch obj := c.value.(type) {
	case *int:
		return *obj
	case *string:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePaymentProfileExpirationMonth.
// It customizes the JSON unmarshaling process for CreatePaymentProfileExpirationMonth objects.
func (c *CreatePaymentProfileExpirationMonth) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(new(int), false, &c.isNumber),
		NewTypeHolder(new(string), false, &c.isString),
	)

	c.value = result
	return err
}

func (c *CreatePaymentProfileExpirationMonth) AsNumber() (
	*int,
	bool) {
	if !c.isNumber {
		return nil, false
	}
	return c.value.(*int), true
}

func (c *CreatePaymentProfileExpirationMonth) AsString() (
	*string,
	bool) {
	if !c.isString {
		return nil, false
	}
	return c.value.(*string), true
}

// internalCreatePaymentProfileExpirationMonth represents a createPaymentProfileExpirationMonth struct.
// This is a container for one-of cases.
type internalCreatePaymentProfileExpirationMonth struct{}

var CreatePaymentProfileExpirationMonthContainer internalCreatePaymentProfileExpirationMonth

// The internalCreatePaymentProfileExpirationMonth instance, wrapping the provided int value.
func (c *internalCreatePaymentProfileExpirationMonth) FromNumber(val int) CreatePaymentProfileExpirationMonth {
	return CreatePaymentProfileExpirationMonth{value: &val}
}

// The internalCreatePaymentProfileExpirationMonth instance, wrapping the provided string value.
func (c *internalCreatePaymentProfileExpirationMonth) FromString(val string) CreatePaymentProfileExpirationMonth {
	return CreatePaymentProfileExpirationMonth{value: &val}
}
