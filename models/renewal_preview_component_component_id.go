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

// RenewalPreviewComponentComponentId represents a RenewalPreviewComponentComponentId struct.
// This is a container for one-of cases.
type RenewalPreviewComponentComponentId struct {
    value    any
    isString bool
    isNumber bool
}

// String implements the fmt.Stringer interface for RenewalPreviewComponentComponentId,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RenewalPreviewComponentComponentId) String() string {
    return fmt.Sprintf("%v", r.value)
}

// MarshalJSON implements the json.Marshaler interface for RenewalPreviewComponentComponentId.
// It customizes the JSON marshaling process for RenewalPreviewComponentComponentId objects.
func (r RenewalPreviewComponentComponentId) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.RenewalPreviewComponentComponentIdContainer.From*` functions to initialize the RenewalPreviewComponentComponentId object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RenewalPreviewComponentComponentId object to a map representation for JSON marshaling.
func (r *RenewalPreviewComponentComponentId) toMap() any {
    switch obj := r.value.(type) {
    case *string:
        return *obj
    case *int:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for RenewalPreviewComponentComponentId.
// It customizes the JSON unmarshaling process for RenewalPreviewComponentComponentId objects.
func (r *RenewalPreviewComponentComponentId) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(string), false, &r.isString),
        NewTypeHolder(new(int), false, &r.isNumber),
    )
    
    r.value = result
    return err
}

func (r *RenewalPreviewComponentComponentId) AsString() (
    *string,
    bool) {
    if !r.isString {
        return nil, false
    }
    return r.value.(*string), true
}

func (r *RenewalPreviewComponentComponentId) AsNumber() (
    *int,
    bool) {
    if !r.isNumber {
        return nil, false
    }
    return r.value.(*int), true
}

// internalRenewalPreviewComponentComponentId represents a renewalPreviewComponentComponentId struct.
// This is a container for one-of cases.
type internalRenewalPreviewComponentComponentId struct {}

var RenewalPreviewComponentComponentIdContainer internalRenewalPreviewComponentComponentId

// The internalRenewalPreviewComponentComponentId instance, wrapping the provided string value.
func (r *internalRenewalPreviewComponentComponentId) FromString(val string) RenewalPreviewComponentComponentId {
    return RenewalPreviewComponentComponentId{value: &val}
}

// The internalRenewalPreviewComponentComponentId instance, wrapping the provided int value.
func (r *internalRenewalPreviewComponentComponentId) FromNumber(val int) RenewalPreviewComponentComponentId {
    return RenewalPreviewComponentComponentId{value: &val}
}
