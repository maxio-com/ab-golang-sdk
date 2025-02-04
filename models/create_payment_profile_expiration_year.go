/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// CreatePaymentProfileExpirationYear represents a CreatePaymentProfileExpirationYear struct.
// This is a container for one-of cases.
type CreatePaymentProfileExpirationYear struct {
    value    any
    isNumber bool
    isString bool
}

// String implements the fmt.Stringer interface for CreatePaymentProfileExpirationYear,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreatePaymentProfileExpirationYear) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CreatePaymentProfileExpirationYear.
// It customizes the JSON marshaling process for CreatePaymentProfileExpirationYear objects.
func (c CreatePaymentProfileExpirationYear) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreatePaymentProfileExpirationYearContainer.From*` functions to initialize the CreatePaymentProfileExpirationYear object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePaymentProfileExpirationYear object to a map representation for JSON marshaling.
func (c *CreatePaymentProfileExpirationYear) toMap() any {
    switch obj := c.value.(type) {
    case *int:
        return *obj
    case *string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePaymentProfileExpirationYear.
// It customizes the JSON unmarshaling process for CreatePaymentProfileExpirationYear objects.
func (c *CreatePaymentProfileExpirationYear) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(int), false, &c.isNumber),
        NewTypeHolder(new(string), false, &c.isString),
    )
    
    c.value = result
    return err
}

func (c *CreatePaymentProfileExpirationYear) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

func (c *CreatePaymentProfileExpirationYear) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

// internalCreatePaymentProfileExpirationYear represents a createPaymentProfileExpirationYear struct.
// This is a container for one-of cases.
type internalCreatePaymentProfileExpirationYear struct {}

var CreatePaymentProfileExpirationYearContainer internalCreatePaymentProfileExpirationYear

// The internalCreatePaymentProfileExpirationYear instance, wrapping the provided int value.
func (c *internalCreatePaymentProfileExpirationYear) FromNumber(val int) CreatePaymentProfileExpirationYear {
    return CreatePaymentProfileExpirationYear{value: &val}
}

// The internalCreatePaymentProfileExpirationYear instance, wrapping the provided string value.
func (c *internalCreatePaymentProfileExpirationYear) FromString(val string) CreatePaymentProfileExpirationYear {
    return CreatePaymentProfileExpirationYear{value: &val}
}
