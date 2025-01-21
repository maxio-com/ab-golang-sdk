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

// RenewalPreviewComponentPricePointId represents a RenewalPreviewComponentPricePointId struct.
// This is a container for one-of cases.
type RenewalPreviewComponentPricePointId struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for RenewalPreviewComponentPricePointId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RenewalPreviewComponentPricePointId) String() string {
    return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for RenewalPreviewComponentPricePointId.
// It customizes the JSON marshaling process for RenewalPreviewComponentPricePointId objects.
func (r RenewalPreviewComponentPricePointId) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.RenewalPreviewComponentPricePointIdContainer.From*` functions to initialize the RenewalPreviewComponentPricePointId object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RenewalPreviewComponentPricePointId object to a map representation for JSON marshaling.
func (r *RenewalPreviewComponentPricePointId) toMap() any {
    switch obj := r.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RenewalPreviewComponentPricePointId.
// It customizes the JSON unmarshaling process for RenewalPreviewComponentPricePointId objects.
func (r *RenewalPreviewComponentPricePointId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &r.isString),
        NewTypeHolder(new(int), false, &r.isNumber),
    )
    
    r.value = result
    return err
}

func (r *RenewalPreviewComponentPricePointId) AsString() (
    *string,
    bool) {
    if !r.isString {
        return nil, false
    }
    return r.value.(*string), true
}

func (r *RenewalPreviewComponentPricePointId) AsNumber() (
    *int,
    bool) {
    if !r.isNumber {
        return nil, false
    }
    return r.value.(*int), true
}

// internalRenewalPreviewComponentPricePointId represents a renewalPreviewComponentPricePointId struct.
// This is a container for one-of cases.
type internalRenewalPreviewComponentPricePointId struct {}

var RenewalPreviewComponentPricePointIdContainer internalRenewalPreviewComponentPricePointId

// The internalRenewalPreviewComponentPricePointId instance, wrapping the provided string value.
func (r *internalRenewalPreviewComponentPricePointId) FromString(val string) RenewalPreviewComponentPricePointId {
    return RenewalPreviewComponentPricePointId{value: &val}
}

// The internalRenewalPreviewComponentPricePointId instance, wrapping the provided int value.
func (r *internalRenewalPreviewComponentPricePointId) FromNumber(val int) RenewalPreviewComponentPricePointId {
    return RenewalPreviewComponentPricePointId{value: &val}
}
