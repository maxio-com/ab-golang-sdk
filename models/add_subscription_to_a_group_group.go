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

// AddSubscriptionToAGroupGroup represents a AddSubscriptionToAGroupGroup struct.
// This is a container for one-of cases.
type AddSubscriptionToAGroupGroup struct {
    value           any
    isGroupSettings bool
    isBoolean       bool
}

// String converts the AddSubscriptionToAGroupGroup object to a string representation.
func (a AddSubscriptionToAGroupGroup) String() string {
    if bytes, err := json.Marshal(a.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for AddSubscriptionToAGroupGroup.
// It customizes the JSON marshaling process for AddSubscriptionToAGroupGroup objects.
func (a AddSubscriptionToAGroupGroup) MarshalJSON() (
    []byte,
    error) {
    if a.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.AddSubscriptionToAGroupGroupContainer.From*` functions to initialize the AddSubscriptionToAGroupGroup object.")
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AddSubscriptionToAGroupGroup object to a map representation for JSON marshaling.
func (a *AddSubscriptionToAGroupGroup) toMap() any {
    switch obj := a.value.(type) {
    case *GroupSettings:
        return obj.toMap()
    case *bool:
        return *obj
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddSubscriptionToAGroupGroup.
// It customizes the JSON unmarshaling process for AddSubscriptionToAGroupGroup objects.
func (a *AddSubscriptionToAGroupGroup) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&GroupSettings{}, false, &a.isGroupSettings),
        NewTypeHolder(new(bool), false, &a.isBoolean),
    )
    
    a.value = result
    return err
}

func (a *AddSubscriptionToAGroupGroup) AsGroupSettings() (
    *GroupSettings,
    bool) {
    if !a.isGroupSettings {
        return nil, false
    }
    return a.value.(*GroupSettings), true
}

func (a *AddSubscriptionToAGroupGroup) AsBoolean() (
    *bool,
    bool) {
    if !a.isBoolean {
        return nil, false
    }
    return a.value.(*bool), true
}

// internalAddSubscriptionToAGroupGroup represents a addSubscriptionToAGroupGroup struct.
// This is a container for one-of cases.
type internalAddSubscriptionToAGroupGroup struct {}

var AddSubscriptionToAGroupGroupContainer internalAddSubscriptionToAGroupGroup

// The internalAddSubscriptionToAGroupGroup instance, wrapping the provided GroupSettings value.
func (a *internalAddSubscriptionToAGroupGroup) FromGroupSettings(val GroupSettings) AddSubscriptionToAGroupGroup {
    return AddSubscriptionToAGroupGroup{value: &val}
}

// The internalAddSubscriptionToAGroupGroup instance, wrapping the provided bool value.
func (a *internalAddSubscriptionToAGroupGroup) FromBoolean(val bool) AddSubscriptionToAGroupGroup {
    return AddSubscriptionToAGroupGroup{value: &val}
}
