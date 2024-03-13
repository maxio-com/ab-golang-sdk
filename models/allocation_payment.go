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

// AllocationPayment represents a AllocationPayment struct.
// This is a container for one-of cases.
type AllocationPayment struct {
	value                  any
	isPaymentForAllocation bool
}

// String converts the AllocationPayment object to a string representation.
func (a AllocationPayment) String() string {
	if bytes, err := json.Marshal(a.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for AllocationPayment.
// It customizes the JSON marshaling process for AllocationPayment objects.
func (a *AllocationPayment) MarshalJSON() (
	[]byte,
	error) {
	if a.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.AllocationPaymentContainer.From*` functions to initialize the AllocationPayment object.")
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AllocationPayment object to a map representation for JSON marshaling.
func (a *AllocationPayment) toMap() any {
	switch obj := a.value.(type) {
	case *PaymentForAllocation:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationPayment.
// It customizes the JSON unmarshaling process for AllocationPayment objects.
func (a *AllocationPayment) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&PaymentForAllocation{}, false, &a.isPaymentForAllocation),
	)

	a.value = result
	return err
}

func (a *AllocationPayment) AsPaymentForAllocation() (
	*PaymentForAllocation,
	bool) {
	if !a.isPaymentForAllocation {
		return nil, false
	}
	return a.value.(*PaymentForAllocation), true
}

// internalAllocationPayment represents a allocationPayment struct.
// This is a container for one-of cases.
type internalAllocationPayment struct{}

var AllocationPaymentContainer internalAllocationPayment

func (a *internalAllocationPayment) FromPaymentForAllocation(val PaymentForAllocation) AllocationPayment {
	return AllocationPayment{value: &val}
}
