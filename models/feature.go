// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// Feature represents a Feature struct.
type Feature struct {
    // A unique, lowercase, underscore-separated identifier for the feature. Immutable once set.
    Key                        string                               `json:"key"`
    // The display name of the feature.
    Name                       string                               `json:"name"`
    Description                Optional[string]                     `json:"description"`
    // The behavior of a feature:
    // - `access_right`: a boolean entitlement. A subscriber either has access or does not.
    // - `usage_limit`: a quantified allowance measured over a recurring period (for example, "10,000 API calls per month").
    // - `service_right`: a free-form value (text, boolean, or number) that isn't a simple access flag or a metered limit.
    Kind                       FeatureKind                          `json:"kind"`
    // Required when `kind` is `usage_limit`.
    Unit                       Optional[string]                     `json:"unit"`
    // Required when `kind` is `service_right`. Ignored for other kinds, where it is inferred automatically.
    ValueType                  *FeatureValueType                    `json:"value_type,omitempty"`
    DefaultValue               Optional[string]                     `json:"default_value"`
    // Only valid when `kind` is `usage_limit`.
    DefaultPeriodicityInterval Optional[int]                        `json:"default_periodicity_interval"`
    // Only valid when `kind` is `usage_limit`. Must be set together with `default_periodicity_interval`.
    DefaultPeriodicityUnit     Optional[EntitlementPeriodicityUnit] `json:"default_periodicity_unit"`
    AdditionalProperties       map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for Feature,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f Feature) String() string {
    return fmt.Sprintf(
    	"Feature[Key=%v, Name=%v, Description=%v, Kind=%v, Unit=%v, ValueType=%v, DefaultValue=%v, DefaultPeriodicityInterval=%v, DefaultPeriodicityUnit=%v, AdditionalProperties=%v]",
    	f.Key, f.Name, f.Description, f.Kind, f.Unit, f.ValueType, f.DefaultValue, f.DefaultPeriodicityInterval, f.DefaultPeriodicityUnit, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Feature.
// It customizes the JSON marshaling process for Feature objects.
func (f Feature) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(f.AdditionalProperties,
        "key", "name", "description", "kind", "unit", "value_type", "default_value", "default_periodicity_interval", "default_periodicity_unit"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(f.toMap())
}

// toMap converts the Feature object to a map representation for JSON marshaling.
func (f Feature) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, f.AdditionalProperties)
    structMap["key"] = f.Key
    structMap["name"] = f.Name
    if f.Description.IsValueSet() {
        if f.Description.Value() != nil {
            structMap["description"] = f.Description.Value()
        } else {
            structMap["description"] = nil
        }
    }
    structMap["kind"] = f.Kind
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

// UnmarshalJSON implements the json.Unmarshaler interface for Feature.
// It customizes the JSON unmarshaling process for Feature objects.
func (f *Feature) UnmarshalJSON(input []byte) error {
    var temp tempFeature
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "key", "name", "description", "kind", "unit", "value_type", "default_value", "default_periodicity_interval", "default_periodicity_unit")
    if err != nil {
    	return err
    }
    f.AdditionalProperties = additionalProperties
    
    f.Key = *temp.Key
    f.Name = *temp.Name
    f.Description = temp.Description
    f.Kind = *temp.Kind
    f.Unit = temp.Unit
    f.ValueType = temp.ValueType
    f.DefaultValue = temp.DefaultValue
    f.DefaultPeriodicityInterval = temp.DefaultPeriodicityInterval
    f.DefaultPeriodicityUnit = temp.DefaultPeriodicityUnit
    return nil
}

// tempFeature is a temporary struct used for validating the fields of Feature.
type tempFeature  struct {
    Key                        *string                              `json:"key"`
    Name                       *string                              `json:"name"`
    Description                Optional[string]                     `json:"description"`
    Kind                       *FeatureKind                         `json:"kind"`
    Unit                       Optional[string]                     `json:"unit"`
    ValueType                  *FeatureValueType                    `json:"value_type,omitempty"`
    DefaultValue               Optional[string]                     `json:"default_value"`
    DefaultPeriodicityInterval Optional[int]                        `json:"default_periodicity_interval"`
    DefaultPeriodicityUnit     Optional[EntitlementPeriodicityUnit] `json:"default_periodicity_unit"`
}

func (f *tempFeature) validate() error {
    var errs []string
    if f.Key == nil {
        errs = append(errs, "required field `key` is missing for type `Feature`")
    }
    if f.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Feature`")
    }
    if f.Kind == nil {
        errs = append(errs, "required field `kind` is missing for type `Feature`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
