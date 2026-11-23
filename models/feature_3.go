// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// Feature3 represents a Feature3 struct.
type Feature3 struct {
    Value                    *string                              `json:"value,omitempty"`
    PeriodicityInterval      Optional[int]                        `json:"periodicity_interval"`
    PeriodicityUnit          Optional[EntitlementPeriodicityUnit] `json:"periodicity_unit"`
    // When `true`, existing subscriptions on the item's price point are updated with the new `value`/periodicity by a background job shortly after. Has no effect on a product/component-level item. When `false` (default) and the item's price point already has subscribers, the update is made on a clone of the price point instead, and existing subscribers are unaffected.
    PropagateToSubscriptions *bool                                `json:"propagate_to_subscriptions,omitempty"`
    AdditionalProperties     map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for Feature3,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f Feature3) String() string {
    return fmt.Sprintf(
    	"Feature3[Value=%v, PeriodicityInterval=%v, PeriodicityUnit=%v, PropagateToSubscriptions=%v, AdditionalProperties=%v]",
    	f.Value, f.PeriodicityInterval, f.PeriodicityUnit, f.PropagateToSubscriptions, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Feature3.
// It customizes the JSON marshaling process for Feature3 objects.
func (f Feature3) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(f.AdditionalProperties,
        "value", "periodicity_interval", "periodicity_unit", "propagate_to_subscriptions"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(f.toMap())
}

// toMap converts the Feature3 object to a map representation for JSON marshaling.
func (f Feature3) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, f.AdditionalProperties)
    if f.Value != nil {
        structMap["value"] = f.Value
    }
    if f.PeriodicityInterval.IsValueSet() {
        if f.PeriodicityInterval.Value() != nil {
            structMap["periodicity_interval"] = f.PeriodicityInterval.Value()
        } else {
            structMap["periodicity_interval"] = nil
        }
    }
    if f.PeriodicityUnit.IsValueSet() {
        if f.PeriodicityUnit.Value() != nil {
            structMap["periodicity_unit"] = f.PeriodicityUnit.Value()
        } else {
            structMap["periodicity_unit"] = nil
        }
    }
    if f.PropagateToSubscriptions != nil {
        structMap["propagate_to_subscriptions"] = f.PropagateToSubscriptions
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Feature3.
// It customizes the JSON unmarshaling process for Feature3 objects.
func (f *Feature3) UnmarshalJSON(input []byte) error {
    var temp tempFeature3
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "value", "periodicity_interval", "periodicity_unit", "propagate_to_subscriptions")
    if err != nil {
    	return err
    }
    f.AdditionalProperties = additionalProperties
    
    f.Value = temp.Value
    f.PeriodicityInterval = temp.PeriodicityInterval
    f.PeriodicityUnit = temp.PeriodicityUnit
    f.PropagateToSubscriptions = temp.PropagateToSubscriptions
    return nil
}

// tempFeature3 is a temporary struct used for validating the fields of Feature3.
type tempFeature3  struct {
    Value                    *string                              `json:"value,omitempty"`
    PeriodicityInterval      Optional[int]                        `json:"periodicity_interval"`
    PeriodicityUnit          Optional[EntitlementPeriodicityUnit] `json:"periodicity_unit"`
    PropagateToSubscriptions *bool                                `json:"propagate_to_subscriptions,omitempty"`
}
