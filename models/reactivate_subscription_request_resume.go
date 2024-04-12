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

// ReactivateSubscriptionRequestResume represents a ReactivateSubscriptionRequestResume struct.
// This is a container for one-of cases.
type ReactivateSubscriptionRequestResume struct {
    value           any
    isBoolean       bool
    isResumeOptions bool
}

// String converts the ReactivateSubscriptionRequestResume object to a string representation.
func (r ReactivateSubscriptionRequestResume) String() string {
    if bytes, err := json.Marshal(r.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ReactivateSubscriptionRequestResume.
// It customizes the JSON marshaling process for ReactivateSubscriptionRequestResume objects.
func (r ReactivateSubscriptionRequestResume) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ReactivateSubscriptionRequestResumeContainer.From*` functions to initialize the ReactivateSubscriptionRequestResume object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ReactivateSubscriptionRequestResume object to a map representation for JSON marshaling.
func (r *ReactivateSubscriptionRequestResume) toMap() any {
    switch obj := r.value.(type) {
    case *bool:
        return *obj
    case *ResumeOptions:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReactivateSubscriptionRequestResume.
// It customizes the JSON unmarshaling process for ReactivateSubscriptionRequestResume objects.
func (r *ReactivateSubscriptionRequestResume) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(new(bool), false, &r.isBoolean),
        NewTypeHolder(&ResumeOptions{}, false, &r.isResumeOptions),
    )
    
    r.value = result
    return err
}

func (r *ReactivateSubscriptionRequestResume) AsBoolean() (
    *bool,
    bool) {
    if !r.isBoolean {
        return nil, false
    }
    return r.value.(*bool), true
}

func (r *ReactivateSubscriptionRequestResume) AsResumeOptions() (
    *ResumeOptions,
    bool) {
    if !r.isResumeOptions {
        return nil, false
    }
    return r.value.(*ResumeOptions), true
}

// internalReactivateSubscriptionRequestResume represents a reactivateSubscriptionRequestResume struct.
// This is a container for one-of cases.
type internalReactivateSubscriptionRequestResume struct {}

var ReactivateSubscriptionRequestResumeContainer internalReactivateSubscriptionRequestResume

// The internalReactivateSubscriptionRequestResume instance, wrapping the provided bool value.
func (r *internalReactivateSubscriptionRequestResume) FromBoolean(val bool) ReactivateSubscriptionRequestResume {
    return ReactivateSubscriptionRequestResume{value: &val}
}

// The internalReactivateSubscriptionRequestResume instance, wrapping the provided ResumeOptions value.
func (r *internalReactivateSubscriptionRequestResume) FromResumeOptions(val ResumeOptions) ReactivateSubscriptionRequestResume {
    return ReactivateSubscriptionRequestResume{value: &val}
}
