// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// ScheduledRenewalUpdateRequestRenewalConfigurationItem represents a ScheduledRenewalUpdateRequestRenewalConfigurationItem struct.
// This is a container for one-of cases.
type ScheduledRenewalUpdateRequestRenewalConfigurationItem struct {
    value                                      any
    isScheduledRenewalItemRequestBodyComponent bool
    isScheduledRenewalItemRequestBodyProduct   bool
}

// String implements the fmt.Stringer interface for ScheduledRenewalUpdateRequestRenewalConfigurationItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalUpdateRequestRenewalConfigurationItem) String() string {
    return fmt.Sprintf("%v", s.value)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalUpdateRequestRenewalConfigurationItem.
// It customizes the JSON marshaling process for ScheduledRenewalUpdateRequestRenewalConfigurationItem objects.
func (s ScheduledRenewalUpdateRequestRenewalConfigurationItem) MarshalJSON() (
    []byte,
    error) {
    if s.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ScheduledRenewalUpdateRequestRenewalConfigurationItemContainer.From*` functions to initialize the ScheduledRenewalUpdateRequestRenewalConfigurationItem object.")
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalUpdateRequestRenewalConfigurationItem object to a map representation for JSON marshaling.
func (s *ScheduledRenewalUpdateRequestRenewalConfigurationItem) toMap() any {
    switch obj := s.value.(type) {
    case *ScheduledRenewalItemRequestBodyComponent:
        return obj.toMap()
    case *ScheduledRenewalItemRequestBodyProduct:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalUpdateRequestRenewalConfigurationItem.
// It customizes the JSON unmarshaling process for ScheduledRenewalUpdateRequestRenewalConfigurationItem objects.
func (s *ScheduledRenewalUpdateRequestRenewalConfigurationItem) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&ScheduledRenewalItemRequestBodyComponent{}, false, &s.isScheduledRenewalItemRequestBodyComponent),
        NewTypeHolder(&ScheduledRenewalItemRequestBodyProduct{}, false, &s.isScheduledRenewalItemRequestBodyProduct),
    )
    
    s.value = result
    return err
}

func (s *ScheduledRenewalUpdateRequestRenewalConfigurationItem) AsScheduledRenewalItemRequestBodyComponent() (
    *ScheduledRenewalItemRequestBodyComponent,
    bool) {
    if !s.isScheduledRenewalItemRequestBodyComponent {
        return nil, false
    }
    return s.value.(*ScheduledRenewalItemRequestBodyComponent), true
}

func (s *ScheduledRenewalUpdateRequestRenewalConfigurationItem) AsScheduledRenewalItemRequestBodyProduct() (
    *ScheduledRenewalItemRequestBodyProduct,
    bool) {
    if !s.isScheduledRenewalItemRequestBodyProduct {
        return nil, false
    }
    return s.value.(*ScheduledRenewalItemRequestBodyProduct), true
}

// internalScheduledRenewalUpdateRequestRenewalConfigurationItem represents a scheduledRenewalUpdateRequestRenewalConfigurationItem struct.
// This is a container for one-of cases.
type internalScheduledRenewalUpdateRequestRenewalConfigurationItem struct {}

var ScheduledRenewalUpdateRequestRenewalConfigurationItemContainer internalScheduledRenewalUpdateRequestRenewalConfigurationItem

// The internalScheduledRenewalUpdateRequestRenewalConfigurationItem instance, wrapping the provided ScheduledRenewalItemRequestBodyComponent value.
func (s *internalScheduledRenewalUpdateRequestRenewalConfigurationItem) FromScheduledRenewalItemRequestBodyComponent(val ScheduledRenewalItemRequestBodyComponent) ScheduledRenewalUpdateRequestRenewalConfigurationItem {
    return ScheduledRenewalUpdateRequestRenewalConfigurationItem{value: &val}
}

// The internalScheduledRenewalUpdateRequestRenewalConfigurationItem instance, wrapping the provided ScheduledRenewalItemRequestBodyProduct value.
func (s *internalScheduledRenewalUpdateRequestRenewalConfigurationItem) FromScheduledRenewalItemRequestBodyProduct(val ScheduledRenewalItemRequestBodyProduct) ScheduledRenewalUpdateRequestRenewalConfigurationItem {
    return ScheduledRenewalUpdateRequestRenewalConfigurationItem{value: &val}
}
