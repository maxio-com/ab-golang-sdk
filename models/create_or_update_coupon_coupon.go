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

// CreateOrUpdateCouponCoupon represents a CreateOrUpdateCouponCoupon struct.
// This is a container for one-of cases.
type CreateOrUpdateCouponCoupon struct {
    value                            any
    isCreateOrUpdatePercentageCoupon bool
    isCreateOrUpdateFlatAmountCoupon bool
}

// String converts the CreateOrUpdateCouponCoupon object to a string representation.
func (c CreateOrUpdateCouponCoupon) String() string {
    if bytes, err := json.Marshal(c.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateCouponCoupon.
// It customizes the JSON marshaling process for CreateOrUpdateCouponCoupon objects.
func (c CreateOrUpdateCouponCoupon) MarshalJSON() (
    []byte,
    error) {
    if c.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.CreateOrUpdateCouponCouponContainer.From*` functions to initialize the CreateOrUpdateCouponCoupon object.")
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateCouponCoupon object to a map representation for JSON marshaling.
func (c *CreateOrUpdateCouponCoupon) toMap() any {
    switch obj := c.value.(type) {
    case *CreateOrUpdatePercentageCoupon:
        return obj.toMap()
    case *CreateOrUpdateFlatAmountCoupon:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdateCouponCoupon.
// It customizes the JSON unmarshaling process for CreateOrUpdateCouponCoupon objects.
func (c *CreateOrUpdateCouponCoupon) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&CreateOrUpdatePercentageCoupon{}, false, &c.isCreateOrUpdatePercentageCoupon),
        NewTypeHolder(&CreateOrUpdateFlatAmountCoupon{}, false, &c.isCreateOrUpdateFlatAmountCoupon),
    )
    
    c.value = result
    return err
}

func (c *CreateOrUpdateCouponCoupon) AsCreateOrUpdatePercentageCoupon() (
    *CreateOrUpdatePercentageCoupon,
    bool) {
    if !c.isCreateOrUpdatePercentageCoupon {
        return nil, false
    }
    return c.value.(*CreateOrUpdatePercentageCoupon), true
}

func (c *CreateOrUpdateCouponCoupon) AsCreateOrUpdateFlatAmountCoupon() (
    *CreateOrUpdateFlatAmountCoupon,
    bool) {
    if !c.isCreateOrUpdateFlatAmountCoupon {
        return nil, false
    }
    return c.value.(*CreateOrUpdateFlatAmountCoupon), true
}

// internalCreateOrUpdateCouponCoupon represents a createOrUpdateCouponCoupon struct.
// This is a container for one-of cases.
type internalCreateOrUpdateCouponCoupon struct {}

var CreateOrUpdateCouponCouponContainer internalCreateOrUpdateCouponCoupon

// The internalCreateOrUpdateCouponCoupon instance, wrapping the provided CreateOrUpdatePercentageCoupon value.
func (c *internalCreateOrUpdateCouponCoupon) FromCreateOrUpdatePercentageCoupon(val CreateOrUpdatePercentageCoupon) CreateOrUpdateCouponCoupon {
    return CreateOrUpdateCouponCoupon{value: &val}
}

// The internalCreateOrUpdateCouponCoupon instance, wrapping the provided CreateOrUpdateFlatAmountCoupon value.
func (c *internalCreateOrUpdateCouponCoupon) FromCreateOrUpdateFlatAmountCoupon(val CreateOrUpdateFlatAmountCoupon) CreateOrUpdateCouponCoupon {
    return CreateOrUpdateCouponCoupon{value: &val}
}
