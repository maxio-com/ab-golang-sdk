// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// FeatureTemplate represents a FeatureTemplate struct.
// A feature that can be granted to subscribers, defined once at the site level and then attached to products or components.
type FeatureTemplate struct {
    // The Advanced Billing id of the feature template.
    Id                         *int                                 `json:"id,omitempty"`
    // A unique, lowercase, underscore-separated identifier for the feature. Immutable once set.
    Key                        *string                              `json:"key,omitempty"`
    // The display name of the feature.
    Name                       *string                              `json:"name,omitempty"`
    Description                Optional[string]                     `json:"description"`
    // The behavior of a feature:
    // - `access_right`: a boolean entitlement. A subscriber either has access or does not.
    // - `usage_limit`: a quantified allowance measured over a recurring period (for example, "10,000 API calls per month").
    // - `service_right`: a free-form value (text, boolean, or number) that isn't a simple access flag or a metered limit.
    Kind                       *FeatureKind                         `json:"kind,omitempty"`
    // The unit the feature is measured in (for example, `requests` or `GB`). Required when `kind` is `usage_limit`.
    Unit                       Optional[string]                     `json:"unit"`
    // The data type of a feature's value. For `access_right` features this is always `boolean`, and for `usage_limit` features this is always `numeric`. For `service_right` features, you choose the value type explicitly.
    ValueType                  *FeatureValueType                    `json:"value_type,omitempty"`
    // A default value used to pre-populate new feature catalog items created from this template.
    DefaultValue               Optional[string]                     `json:"default_value"`
    // For `usage_limit` features, the default periodicity interval used to pre-populate new feature catalog items. Always `null` for other kinds.
    DefaultPeriodicityInterval Optional[int]                        `json:"default_periodicity_interval"`
    // For `usage_limit` features, the default periodicity unit used to pre-populate new feature catalog items. Always `null` for other kinds.
    DefaultPeriodicityUnit     Optional[EntitlementPeriodicityUnit] `json:"default_periodicity_unit"`
    // The date and time the feature template was archived, or `null` if it is active.
    ArchivedAt                 Optional[time.Time]                  `json:"archived_at"`
    CreatedAt                  *time.Time                           `json:"created_at,omitempty"`
    UpdatedAt                  *time.Time                           `json:"updated_at,omitempty"`
    // The number of **components** this feature template is currently attached to via an active feature catalog item. Despite the name, this counts components, not products. In the Advanced Billing UI, components are labeled "Products."
    ProductsCount              *int                                 `json:"products_count,omitempty"`
    // The number of **products** this feature template is currently attached to via an active feature catalog item. Despite the name, this counts products, not plans. In the Advanced Billing UI, products are labeled "Plans."
    PlansCount                 *int                                 `json:"plans_count,omitempty"`
    AdditionalProperties       map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for FeatureTemplate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f FeatureTemplate) String() string {
    return fmt.Sprintf(
    	"FeatureTemplate[Id=%v, Key=%v, Name=%v, Description=%v, Kind=%v, Unit=%v, ValueType=%v, DefaultValue=%v, DefaultPeriodicityInterval=%v, DefaultPeriodicityUnit=%v, ArchivedAt=%v, CreatedAt=%v, UpdatedAt=%v, ProductsCount=%v, PlansCount=%v, AdditionalProperties=%v]",
    	f.Id, f.Key, f.Name, f.Description, f.Kind, f.Unit, f.ValueType, f.DefaultValue, f.DefaultPeriodicityInterval, f.DefaultPeriodicityUnit, f.ArchivedAt, f.CreatedAt, f.UpdatedAt, f.ProductsCount, f.PlansCount, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for FeatureTemplate.
// It customizes the JSON marshaling process for FeatureTemplate objects.
func (f FeatureTemplate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(f.AdditionalProperties,
        "id", "key", "name", "description", "kind", "unit", "value_type", "default_value", "default_periodicity_interval", "default_periodicity_unit", "archived_at", "created_at", "updated_at", "products_count", "plans_count"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(f.toMap())
}

// toMap converts the FeatureTemplate object to a map representation for JSON marshaling.
func (f FeatureTemplate) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, f.AdditionalProperties)
    if f.Id != nil {
        structMap["id"] = f.Id
    }
    if f.Key != nil {
        structMap["key"] = f.Key
    }
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
    if f.Kind != nil {
        structMap["kind"] = f.Kind
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
    if f.ArchivedAt.IsValueSet() {
        var ArchivedAtVal *string = nil
        if f.ArchivedAt.Value() != nil {
            val := f.ArchivedAt.Value().Format(time.RFC3339)
            ArchivedAtVal = &val
        }
        if f.ArchivedAt.Value() != nil {
            structMap["archived_at"] = ArchivedAtVal
        } else {
            structMap["archived_at"] = nil
        }
    }
    if f.CreatedAt != nil {
        structMap["created_at"] = f.CreatedAt.Format(time.RFC3339)
    }
    if f.UpdatedAt != nil {
        structMap["updated_at"] = f.UpdatedAt.Format(time.RFC3339)
    }
    if f.ProductsCount != nil {
        structMap["products_count"] = f.ProductsCount
    }
    if f.PlansCount != nil {
        structMap["plans_count"] = f.PlansCount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FeatureTemplate.
// It customizes the JSON unmarshaling process for FeatureTemplate objects.
func (f *FeatureTemplate) UnmarshalJSON(input []byte) error {
    var temp tempFeatureTemplate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "key", "name", "description", "kind", "unit", "value_type", "default_value", "default_periodicity_interval", "default_periodicity_unit", "archived_at", "created_at", "updated_at", "products_count", "plans_count")
    if err != nil {
    	return err
    }
    f.AdditionalProperties = additionalProperties
    
    f.Id = temp.Id
    f.Key = temp.Key
    f.Name = temp.Name
    f.Description = temp.Description
    f.Kind = temp.Kind
    f.Unit = temp.Unit
    f.ValueType = temp.ValueType
    f.DefaultValue = temp.DefaultValue
    f.DefaultPeriodicityInterval = temp.DefaultPeriodicityInterval
    f.DefaultPeriodicityUnit = temp.DefaultPeriodicityUnit
    f.ArchivedAt.ShouldSetValue(temp.ArchivedAt.IsValueSet())
    if temp.ArchivedAt.Value() != nil {
        ArchivedAtVal, err := time.Parse(time.RFC3339, (*temp.ArchivedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse archived_at as % s format.", time.RFC3339)
        }
        f.ArchivedAt.SetValue(&ArchivedAtVal)
    }
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        f.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        f.UpdatedAt = &UpdatedAtVal
    }
    f.ProductsCount = temp.ProductsCount
    f.PlansCount = temp.PlansCount
    return nil
}

// tempFeatureTemplate is a temporary struct used for validating the fields of FeatureTemplate.
type tempFeatureTemplate  struct {
    Id                         *int                                 `json:"id,omitempty"`
    Key                        *string                              `json:"key,omitempty"`
    Name                       *string                              `json:"name,omitempty"`
    Description                Optional[string]                     `json:"description"`
    Kind                       *FeatureKind                         `json:"kind,omitempty"`
    Unit                       Optional[string]                     `json:"unit"`
    ValueType                  *FeatureValueType                    `json:"value_type,omitempty"`
    DefaultValue               Optional[string]                     `json:"default_value"`
    DefaultPeriodicityInterval Optional[int]                        `json:"default_periodicity_interval"`
    DefaultPeriodicityUnit     Optional[EntitlementPeriodicityUnit] `json:"default_periodicity_unit"`
    ArchivedAt                 Optional[string]                     `json:"archived_at"`
    CreatedAt                  *string                              `json:"created_at,omitempty"`
    UpdatedAt                  *string                              `json:"updated_at,omitempty"`
    ProductsCount              *int                                 `json:"products_count,omitempty"`
    PlansCount                 *int                                 `json:"plans_count,omitempty"`
}
