// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// Feature2 represents a Feature2 struct.
type Feature2 struct {
    // The id of the feature template to attach.
    FeatureTemplateId        int                                  `json:"feature_template_id"`
    // The value granted. Defaults to the feature template's `default_value`; required when the template has none.
    Value                    *string                              `json:"value,omitempty"`
    // Required for `usage_limit` features, unless the feature template has a default periodicity, which is used when omitted. Must be omitted for other kinds.
    PeriodicityInterval      Optional[int]                        `json:"periodicity_interval"`
    // Required for `usage_limit` features, unless the feature template has a default periodicity, which is used when omitted. Must be omitted for other kinds.
    PeriodicityUnit          Optional[EntitlementPeriodicityUnit] `json:"periodicity_unit"`
    // Set together with `price_point_id` to scope the feature catalog item to a single price point. Omit both to create a product/component-level item, which is a template: it is copied onto price points created afterwards (except bulk-created ones; Clone Component Price Point applies only the items sent with it), but not added to price points that already exist or to their subscribers.
    PricePointType           Optional[FeatureOwnerPricePointType] `json:"price_point_type"`
    PricePointId             Optional[int]                        `json:"price_point_id"`
    // Applies to price-point-scoped items only. When `true`, existing subscriptions on that price point are granted an entitlement for this feature by a background job shortly after, instead of at their next subscription change. When `false` (default) and the price point already has subscribers, the feature is created on a clone of the price point instead, and existing subscribers are unaffected. On a product/component-level item it has no effect, except when the call restores that item's archived price-point copies (creating an item again for a feature template whose item was archived): then `true` grants the feature to those price points' subscribers, and without it the call returns `422` if any of them lack the feature.
    PropagateToSubscriptions *bool                                `json:"propagate_to_subscriptions,omitempty"`
    AdditionalProperties     map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for Feature2,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f Feature2) String() string {
    return fmt.Sprintf(
    	"Feature2[FeatureTemplateId=%v, Value=%v, PeriodicityInterval=%v, PeriodicityUnit=%v, PricePointType=%v, PricePointId=%v, PropagateToSubscriptions=%v, AdditionalProperties=%v]",
    	f.FeatureTemplateId, f.Value, f.PeriodicityInterval, f.PeriodicityUnit, f.PricePointType, f.PricePointId, f.PropagateToSubscriptions, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Feature2.
// It customizes the JSON marshaling process for Feature2 objects.
func (f Feature2) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(f.AdditionalProperties,
        "feature_template_id", "value", "periodicity_interval", "periodicity_unit", "price_point_type", "price_point_id", "propagate_to_subscriptions"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(f.toMap())
}

// toMap converts the Feature2 object to a map representation for JSON marshaling.
func (f Feature2) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, f.AdditionalProperties)
    structMap["feature_template_id"] = f.FeatureTemplateId
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
    if f.PropagateToSubscriptions != nil {
        structMap["propagate_to_subscriptions"] = f.PropagateToSubscriptions
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Feature2.
// It customizes the JSON unmarshaling process for Feature2 objects.
func (f *Feature2) UnmarshalJSON(input []byte) error {
    var temp tempFeature2
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "feature_template_id", "value", "periodicity_interval", "periodicity_unit", "price_point_type", "price_point_id", "propagate_to_subscriptions")
    if err != nil {
    	return err
    }
    f.AdditionalProperties = additionalProperties
    
    f.FeatureTemplateId = *temp.FeatureTemplateId
    f.Value = temp.Value
    f.PeriodicityInterval = temp.PeriodicityInterval
    f.PeriodicityUnit = temp.PeriodicityUnit
    f.PricePointType = temp.PricePointType
    f.PricePointId = temp.PricePointId
    f.PropagateToSubscriptions = temp.PropagateToSubscriptions
    return nil
}

// tempFeature2 is a temporary struct used for validating the fields of Feature2.
type tempFeature2  struct {
    FeatureTemplateId        *int                                 `json:"feature_template_id"`
    Value                    *string                              `json:"value,omitempty"`
    PeriodicityInterval      Optional[int]                        `json:"periodicity_interval"`
    PeriodicityUnit          Optional[EntitlementPeriodicityUnit] `json:"periodicity_unit"`
    PricePointType           Optional[FeatureOwnerPricePointType] `json:"price_point_type"`
    PricePointId             Optional[int]                        `json:"price_point_id"`
    PropagateToSubscriptions *bool                                `json:"propagate_to_subscriptions,omitempty"`
}

func (f *tempFeature2) validate() error {
    var errs []string
    if f.FeatureTemplateId == nil {
        errs = append(errs, "required field `feature_template_id` is missing for type `Feature2`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
