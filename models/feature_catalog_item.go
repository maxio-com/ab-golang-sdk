// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// FeatureCatalogItem represents a FeatureCatalogItem struct.
// A feature template attached to a specific product or component (or one of their price points), with a concrete value. When a subscriber signs up for or is assigned this product/component, the feature catalog item is provisioned as an entitlement on their subscription.
type FeatureCatalogItem struct {
    Id                   *int                                 `json:"id,omitempty"`
    // The id of the feature template this item was created from.
    FeatureTemplateId    *int                                 `json:"feature_template_id,omitempty"`
    // The `key` of the parent feature template.
    FeatureKey           *string                              `json:"feature_key,omitempty"`
    // The `name` of the parent feature template.
    FeatureName          *string                              `json:"feature_name,omitempty"`
    // The behavior of a feature:
    // - `access_right`: a boolean entitlement. A subscriber either has access or does not.
    // - `usage_limit`: a quantified allowance measured over a recurring period (for example, "10,000 API calls per month").
    // - `service_right`: a free-form value (text, boolean, or number) that isn't a simple access flag or a metered limit.
    FeatureKind          *FeatureKind                         `json:"feature_kind,omitempty"`
    // The value granted by this feature catalog item. Interpreted according to `feature_kind`: `"true"`/`"false"` for `access_right`, a numeric string for `usage_limit`, or any string for `service_right` (shaped by the feature template's `value_type`).
    Value                *string                              `json:"value,omitempty"`
    // Set when `feature_kind` is `usage_limit`; `null` otherwise.
    PeriodicityInterval  Optional[int]                        `json:"periodicity_interval"`
    // Set when `feature_kind` is `usage_limit`; `null` otherwise.
    PeriodicityUnit      Optional[EntitlementPeriodicityUnit] `json:"periodicity_unit"`
    // `null` when this feature catalog item applies to every price point of its owning product/component. Set when the feature catalog item is an override for one specific price point.
    PricePointType       Optional[FeatureOwnerPricePointType] `json:"price_point_type"`
    // Set together with `price_point_type` for price-point-specific overrides.
    PricePointId         Optional[int]                        `json:"price_point_id"`
    ArchivedAt           Optional[time.Time]                  `json:"archived_at"`
    CreatedAt            *time.Time                           `json:"created_at,omitempty"`
    UpdatedAt            *time.Time                           `json:"updated_at,omitempty"`
    AdditionalProperties map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for FeatureCatalogItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f FeatureCatalogItem) String() string {
    return fmt.Sprintf(
    	"FeatureCatalogItem[Id=%v, FeatureTemplateId=%v, FeatureKey=%v, FeatureName=%v, FeatureKind=%v, Value=%v, PeriodicityInterval=%v, PeriodicityUnit=%v, PricePointType=%v, PricePointId=%v, ArchivedAt=%v, CreatedAt=%v, UpdatedAt=%v, AdditionalProperties=%v]",
    	f.Id, f.FeatureTemplateId, f.FeatureKey, f.FeatureName, f.FeatureKind, f.Value, f.PeriodicityInterval, f.PeriodicityUnit, f.PricePointType, f.PricePointId, f.ArchivedAt, f.CreatedAt, f.UpdatedAt, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for FeatureCatalogItem.
// It customizes the JSON marshaling process for FeatureCatalogItem objects.
func (f FeatureCatalogItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(f.AdditionalProperties,
        "id", "feature_template_id", "feature_key", "feature_name", "feature_kind", "value", "periodicity_interval", "periodicity_unit", "price_point_type", "price_point_id", "archived_at", "created_at", "updated_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(f.toMap())
}

// toMap converts the FeatureCatalogItem object to a map representation for JSON marshaling.
func (f FeatureCatalogItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, f.AdditionalProperties)
    if f.Id != nil {
        structMap["id"] = f.Id
    }
    if f.FeatureTemplateId != nil {
        structMap["feature_template_id"] = f.FeatureTemplateId
    }
    if f.FeatureKey != nil {
        structMap["feature_key"] = f.FeatureKey
    }
    if f.FeatureName != nil {
        structMap["feature_name"] = f.FeatureName
    }
    if f.FeatureKind != nil {
        structMap["feature_kind"] = f.FeatureKind
    }
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
    if f.PricePointType.IsValueSet() {
        if f.PricePointType.Value() != nil {
            structMap["price_point_type"] = f.PricePointType.Value()
        } else {
            structMap["price_point_type"] = nil
        }
    }
    if f.PricePointId.IsValueSet() {
        if f.PricePointId.Value() != nil {
            structMap["price_point_id"] = f.PricePointId.Value()
        } else {
            structMap["price_point_id"] = nil
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
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FeatureCatalogItem.
// It customizes the JSON unmarshaling process for FeatureCatalogItem objects.
func (f *FeatureCatalogItem) UnmarshalJSON(input []byte) error {
    var temp tempFeatureCatalogItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "feature_template_id", "feature_key", "feature_name", "feature_kind", "value", "periodicity_interval", "periodicity_unit", "price_point_type", "price_point_id", "archived_at", "created_at", "updated_at")
    if err != nil {
    	return err
    }
    f.AdditionalProperties = additionalProperties
    
    f.Id = temp.Id
    f.FeatureTemplateId = temp.FeatureTemplateId
    f.FeatureKey = temp.FeatureKey
    f.FeatureName = temp.FeatureName
    f.FeatureKind = temp.FeatureKind
    f.Value = temp.Value
    f.PeriodicityInterval = temp.PeriodicityInterval
    f.PeriodicityUnit = temp.PeriodicityUnit
    f.PricePointType = temp.PricePointType
    f.PricePointId = temp.PricePointId
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
    return nil
}

// tempFeatureCatalogItem is a temporary struct used for validating the fields of FeatureCatalogItem.
type tempFeatureCatalogItem  struct {
    Id                  *int                                 `json:"id,omitempty"`
    FeatureTemplateId   *int                                 `json:"feature_template_id,omitempty"`
    FeatureKey          *string                              `json:"feature_key,omitempty"`
    FeatureName         *string                              `json:"feature_name,omitempty"`
    FeatureKind         *FeatureKind                         `json:"feature_kind,omitempty"`
    Value               *string                              `json:"value,omitempty"`
    PeriodicityInterval Optional[int]                        `json:"periodicity_interval"`
    PeriodicityUnit     Optional[EntitlementPeriodicityUnit] `json:"periodicity_unit"`
    PricePointType      Optional[FeatureOwnerPricePointType] `json:"price_point_type"`
    PricePointId        Optional[int]                        `json:"price_point_id"`
    ArchivedAt          Optional[string]                     `json:"archived_at"`
    CreatedAt           *string                              `json:"created_at,omitempty"`
    UpdatedAt           *string                              `json:"updated_at,omitempty"`
}
