// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem represents a ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem struct.
// This is a container for one-of cases.
type ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem struct {
    value                                      any
    isScheduledRenewalItemRequestBodyComponent bool
    isScheduledRenewalItemRequestBodyProduct   bool
}

// String implements the fmt.Stringer interface for ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem.
// It customizes the JSON marshaling process for ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem objects.
func (s ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ScheduledRenewalConfigurationItemRequestRenewalConfigurationItemContainer.From*` functions to initialize the ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem object to a map representation for JSON marshaling.
func (s *ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem) toMap() any {
    switch obj := s.value.(type) {
    case *ScheduledRenewalItemRequestBodyComponent:
        return obj.toMap()
    case *ScheduledRenewalItemRequestBodyProduct:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem.
// It customizes the JSON unmarshaling process for ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem objects.
func (s *ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&ScheduledRenewalItemRequestBodyComponent{}, false, &s.isScheduledRenewalItemRequestBodyComponent),
        NewTypeHolder(&ScheduledRenewalItemRequestBodyProduct{}, false, &s.isScheduledRenewalItemRequestBodyProduct),
    )
    
    s.value = result
    return err
}

func (s *ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem) AsScheduledRenewalItemRequestBodyComponent() (
    *ScheduledRenewalItemRequestBodyComponent,
    bool) {
    if !s.isScheduledRenewalItemRequestBodyComponent {
        return nil, false
    }
    return s.value.(*ScheduledRenewalItemRequestBodyComponent), true
}

func (s *ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem) AsScheduledRenewalItemRequestBodyProduct() (
    *ScheduledRenewalItemRequestBodyProduct,
    bool) {
    if !s.isScheduledRenewalItemRequestBodyProduct {
        return nil, false
    }
    return s.value.(*ScheduledRenewalItemRequestBodyProduct), true
}

// internalScheduledRenewalConfigurationItemRequestRenewalConfigurationItem represents a scheduledRenewalConfigurationItemRequestRenewalConfigurationItem struct.
// This is a container for one-of cases.
type internalScheduledRenewalConfigurationItemRequestRenewalConfigurationItem struct {}

var ScheduledRenewalConfigurationItemRequestRenewalConfigurationItemContainer internalScheduledRenewalConfigurationItemRequestRenewalConfigurationItem

// The internalScheduledRenewalConfigurationItemRequestRenewalConfigurationItem instance, wrapping the provided ScheduledRenewalItemRequestBodyComponent value.
func (s *internalScheduledRenewalConfigurationItemRequestRenewalConfigurationItem) FromScheduledRenewalItemRequestBodyComponent(val ScheduledRenewalItemRequestBodyComponent) ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem {
    return ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem{value: &val}
}

// The internalScheduledRenewalConfigurationItemRequestRenewalConfigurationItem instance, wrapping the provided ScheduledRenewalItemRequestBodyProduct value.
func (s *internalScheduledRenewalConfigurationItemRequestRenewalConfigurationItem) FromScheduledRenewalItemRequestBodyProduct(val ScheduledRenewalItemRequestBodyProduct) ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem {
    return ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem{value: &val}
}
