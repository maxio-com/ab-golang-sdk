// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// PurchaseQuantity represents a PurchaseQuantity struct.
// This is a container for one-of cases.
type PurchaseQuantity struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for PurchaseQuantity,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PurchaseQuantity) String() string {
    return fmt.Sprintf("%v", p.value)
}

// MarshalJSON implements the json.Marshaler interface for PurchaseQuantity.
// It customizes the JSON marshaling process for PurchaseQuantity objects.
func (p PurchaseQuantity) MarshalJSON() (
    []byte,
    error) {
    if p.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.PurchaseQuantityContainer.From*` functions to initialize the PurchaseQuantity object.")
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PurchaseQuantity object to a map representation for JSON marshaling.
func (p *PurchaseQuantity) toMap() any {
    switch obj := p.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for PurchaseQuantity.
// It customizes the JSON unmarshaling process for PurchaseQuantity objects.
func (p *PurchaseQuantity) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &p.isString),
        NewTypeHolder(new(int), false, &p.isNumber),
    )
    
    p.value = result
    return err
}

func (p *PurchaseQuantity) AsString() (
    *string,
    bool) {
    if !p.isString {
        return nil, false
    }
    return p.value.(*string), true
}

func (p *PurchaseQuantity) AsNumber() (
    *int,
    bool) {
    if !p.isNumber {
        return nil, false
    }
    return p.value.(*int), true
}

// internalPurchaseQuantity represents a purchaseQuantity struct.
// This is a container for one-of cases.
type internalPurchaseQuantity struct {}

var PurchaseQuantityContainer internalPurchaseQuantity

// The internalPurchaseQuantity instance, wrapping the provided string value.
func (p *internalPurchaseQuantity) FromString(val string) PurchaseQuantity {
    return PurchaseQuantity{value: &val}
}

// The internalPurchaseQuantity instance, wrapping the provided int value.
func (p *internalPurchaseQuantity) FromNumber(val int) PurchaseQuantity {
    return PurchaseQuantity{value: &val}
}
