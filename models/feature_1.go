// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// Feature1 represents a Feature1 struct.
// `key` cannot be changed once set. `kind` cannot be changed once any feature catalog item has been created from this template.
type Feature1 struct {
    Name                       *string                              `json:"name,omitempty"`
    Description                Optional[string]                     `json:"description"`
    Unit                       Optional[string]                     `json:"unit"`
    // The data type of a feature's value. For `access_right` features this is always `boolean`, and for `usage_limit` features this is always `numeric`. For `service_right` features, you choose the value type explicitly.
    ValueType                  *FeatureValueType                    `json:"value_type,omitempty"`
    DefaultValue               Optional[string]                     `json:"default_value"`
    DefaultPeriodicityInterval Optional[int]                        `json:"default_periodicity_interval"`
    DefaultPeriodicityUnit     Optional[EntitlementPeriodicityUnit] `json:"default_periodicity_unit"`
    AdditionalProperties       map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for Feature1,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f Feature1) String() string {
    return fmt.Sprintf(
    	"Feature1[Name=%v, Description=%v, Unit=%v, ValueType=%v, DefaultValue=%v, DefaultPeriodicityInterval=%v, DefaultPeriodicityUnit=%v, AdditionalProperties=%v]",
    	f.Name, f.Description, f.Unit, f.ValueType, f.DefaultValue, f.DefaultPeriodicityInterval, f.DefaultPeriodicityUnit, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Feature1.
// It customizes the JSON marshaling process for Feature1 objects.
func (f Feature1) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(f.AdditionalProperties,
        "name", "description", "unit", "value_type", "default_value", "default_periodicity_interval", "default_periodicity_unit"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(f.toMap())
}

// toMap converts the Feature1 object to a map representation for JSON marshaling.
func (f Feature1) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, f.AdditionalProperties)
    if f.Name != nil {
        structMap["name"] = f.Name
    }
    if f.Description.IsValueSet() {
        if f.Description.Value() != nil {
            structMap["description"] = f.Description.Value()
        } else {
            structMap["description"] = nil
        }
    }
    if f.Unit.IsValueSet() {
        if f.Unit.Value() != nil {
            structMap["unit"] = f.Unit.Value()
        } else {
            structMap["unit"] = nil
        }
    }
    if f.ValueType != nil {
        structMap["value_type"] = f.ValueType
    }
    if f.DefaultValue.IsValueSet() {
        if f.DefaultValue.Value() != nil {
            structMap["default_value"] = f.DefaultValue.Value()
        } else {
            structMap["default_value"] = nil
        }
    }
    if f.DefaultPeriodicityInterval.IsValueSet() {
        if f.DefaultPeriodicityInterval.Value() != nil {
            structMap["default_periodicity_interval"] = f.DefaultPeriodicityInterval.Value()
        } else {
            structMap["default_periodicity_interval"] = nil
        }
    }
    if f.DefaultPeriodicityUnit.IsValueSet() {
        if f.DefaultPeriodicityUnit.Value() != nil {
            structMap["default_periodicity_unit"] = f.DefaultPeriodicityUnit.Value()
        } else {
            structMap["default_periodicity_unit"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Feature1.
// It customizes the JSON unmarshaling process for Feature1 objects.
func (f *Feature1) UnmarshalJSON(input []byte) error {
    var temp tempFeature1
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "description", "unit", "value_type", "default_value", "default_periodicity_interval", "default_periodicity_unit")
    if err != nil {
    	return err
    }
    f.AdditionalProperties = additionalProperties
    
    f.Name = temp.Name
    f.Description = temp.Description
    f.Unit = temp.Unit
    f.ValueType = temp.ValueType
    f.DefaultValue = temp.DefaultValue
    f.DefaultPeriodicityInterval = temp.DefaultPeriodicityInterval
    f.DefaultPeriodicityUnit = temp.DefaultPeriodicityUnit
    return nil
}

// tempFeature1 is a temporary struct used for validating the fields of Feature1.
type tempFeature1  struct {
    Name                       *string                              `json:"name,omitempty"`
    Description                Optional[string]                     `json:"description"`
    Unit                       Optional[string]                     `json:"unit"`
    ValueType                  *FeatureValueType                    `json:"value_type,omitempty"`
    DefaultValue               Optional[string]                     `json:"default_value"`
    DefaultPeriodicityInterval Optional[int]                        `json:"default_periodicity_interval"`
    DefaultPeriodicityUnit     Optional[EntitlementPeriodicityUnit] `json:"default_periodicity_unit"`
}
