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

// CouponCompoundingStrategy represents a CouponCompoundingStrategy struct.
// This is a container for any-of cases.
type CouponCompoundingStrategy struct {
	value                 any
	isCompoundingStrategy bool
}

// String converts the CouponCompoundingStrategy object to a string representation.
func (c CouponCompoundingStrategy) String() string {
	if bytes, err := json.Marshal(c.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for CouponCompoundingStrategy.
// It customizes the JSON marshaling process for CouponCompoundingStrategy objects.
func (c *CouponCompoundingStrategy) MarshalJSON() (
	[]byte,
	error) {
	if c.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.CouponCompoundingStrategyContainer.From*` functions to initialize the CouponCompoundingStrategy object.")
	}
	return json.Marshal(c.toMap())
}

// toMap converts the CouponCompoundingStrategy object to a map representation for JSON marshaling.
func (c *CouponCompoundingStrategy) toMap() any {
	switch obj := c.value.(type) {
	case *CompoundingStrategy:
		return *obj
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponCompoundingStrategy.
// It customizes the JSON unmarshaling process for CouponCompoundingStrategy objects.
func (c *CouponCompoundingStrategy) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallAnyOf(input,
		NewTypeHolder(new(CompoundingStrategy), false, &c.isCompoundingStrategy),
	)

	c.value = result
	return err
}

func (c *CouponCompoundingStrategy) AsCompoundingStrategy() (
	*CompoundingStrategy,
	bool) {
	if !c.isCompoundingStrategy {
		return nil, false
	}
	return c.value.(*CompoundingStrategy), true
}

// internalCouponCompoundingStrategy represents a couponCompoundingStrategy struct.
// This is a container for any-of cases.
type internalCouponCompoundingStrategy struct{}

var CouponCompoundingStrategyContainer internalCouponCompoundingStrategy

func (c *internalCouponCompoundingStrategy) FromCompoundingStrategy(val CompoundingStrategy) CouponCompoundingStrategy {
	return CouponCompoundingStrategy{value: &val}
}
