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

// CustomerErrorResponseErrors represents a CustomerErrorResponseErrors struct.
// This is a container for one-of cases.
type CustomerErrorResponseErrors struct {
    value           any
    isCustomerError bool
    isArrayOfString bool
}

// String implements the fmt.Stringer interface for CustomerErrorResponseErrors,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CustomerErrorResponseErrors) String() string {
    return fmt.Sprintf("%v", c.value)
}

// MarshalJSON implements the json.Marshaler interface for CustomerErrorResponseErrors.
// It customizes the JSON marshaling process for CustomerErrorResponseErrors objects.
func (c CustomerErrorResponseErrors) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CustomerErrorResponseErrorsContainer.From*` functions to initialize the CustomerErrorResponseErrors object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerErrorResponseErrors object to a map representation for JSON marshaling.
func (c *CustomerErrorResponseErrors) toMap() any {
    switch obj := c.value.(type) {
    case *CustomerError:
        return obj.toMap()
    case *[]string:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerErrorResponseErrors.
// It customizes the JSON unmarshaling process for CustomerErrorResponseErrors objects.
func (c *CustomerErrorResponseErrors) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&CustomerError{}, false, &c.isCustomerError),
        NewTypeHolder(new([]string), false, &c.isArrayOfString),
    )
    
    c.value = result
    return err
}

func (c *CustomerErrorResponseErrors) AsCustomerError() (
    *CustomerError,
    bool) {
    if !c.isCustomerError {
        return nil, false
    }
    return c.value.(*CustomerError), true
}

func (c *CustomerErrorResponseErrors) AsArrayOfString() (
    *[]string,
    bool) {
    if !c.isArrayOfString {
        return nil, false
    }
    return c.value.(*[]string), true
}

// internalCustomerErrorResponseErrors represents a customerErrorResponseErrors struct.
// This is a container for one-of cases.
type internalCustomerErrorResponseErrors struct {}

var CustomerErrorResponseErrorsContainer internalCustomerErrorResponseErrors

// The internalCustomerErrorResponseErrors instance, wrapping the provided CustomerError value.
func (c *internalCustomerErrorResponseErrors) FromCustomerError(val CustomerError) CustomerErrorResponseErrors {
    return CustomerErrorResponseErrors{value: &val}
}

// The internalCustomerErrorResponseErrors instance, wrapping the provided []string value.
func (c *internalCustomerErrorResponseErrors) FromArrayOfString(val []string) CustomerErrorResponseErrors {
    return CustomerErrorResponseErrors{value: &val}
}
