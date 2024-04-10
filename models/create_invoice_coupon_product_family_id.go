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

// CreateInvoiceCouponProductFamilyId represents a CreateInvoiceCouponProductFamilyId struct.
// This is a container for one-of cases.
type CreateInvoiceCouponProductFamilyId struct {
    value    any
    isString bool
    isNumber bool
}

// String converts the CreateInvoiceCouponProductFamilyId object to a string representation.
func (c CreateInvoiceCouponProductFamilyId) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceCouponProductFamilyId.
// It customizes the JSON marshaling process for CreateInvoiceCouponProductFamilyId objects.
func (c CreateInvoiceCouponProductFamilyId) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateInvoiceCouponProductFamilyIdContainer.From*` functions to initialize the CreateInvoiceCouponProductFamilyId object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceCouponProductFamilyId object to a map representation for JSON marshaling.
func (c *CreateInvoiceCouponProductFamilyId) toMap() any {
    switch obj := c.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceCouponProductFamilyId.
// It customizes the JSON unmarshaling process for CreateInvoiceCouponProductFamilyId objects.
func (c *CreateInvoiceCouponProductFamilyId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &c.isString),
        NewTypeHolder(new(int), false, &c.isNumber),
    )
    
    c.value = result
    return err
}

func (c *CreateInvoiceCouponProductFamilyId) AsString() (
    *string,
    bool) {
    if !c.isString {
        return nil, false
    }
    return c.value.(*string), true
}

func (c *CreateInvoiceCouponProductFamilyId) AsNumber() (
    *int,
    bool) {
    if !c.isNumber {
        return nil, false
    }
    return c.value.(*int), true
}

// internalCreateInvoiceCouponProductFamilyId represents a createInvoiceCouponProductFamilyId struct.
// This is a container for one-of cases.
type internalCreateInvoiceCouponProductFamilyId struct {}

var CreateInvoiceCouponProductFamilyIdContainer internalCreateInvoiceCouponProductFamilyId

// The internalCreateInvoiceCouponProductFamilyId instance, wrapping the provided string value.
func (c *internalCreateInvoiceCouponProductFamilyId) FromString(val string) CreateInvoiceCouponProductFamilyId {
    return CreateInvoiceCouponProductFamilyId{value: &val}
}

// The internalCreateInvoiceCouponProductFamilyId instance, wrapping the provided int value.
func (c *internalCreateInvoiceCouponProductFamilyId) FromNumber(val int) CreateInvoiceCouponProductFamilyId {
    return CreateInvoiceCouponProductFamilyId{value: &val}
}
