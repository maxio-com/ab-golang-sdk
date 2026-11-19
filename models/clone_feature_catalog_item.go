// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CloneFeatureCatalogItem represents a CloneFeatureCatalogItem struct.
// A feature to put on the cloned price point. Requires Entitlements to be enabled for the site; ignored otherwise.
// Unlike Create Feature Catalog Item, nothing falls back to the feature template's defaults: send `value`, and for a `usage_limit` feature its periodicity (and any grant expiration) explicitly. If any item can't be applied, for example a missing `value`, a `usage_limit` item without periodicity, or an archived or unknown feature template, the whole clone fails with `422` and nothing is created.
type CloneFeatureCatalogItem struct {
    FeatureTemplateId       int                                  `json:"feature_template_id"`
    Value                   string                               `json:"value"`
    // Required for `usage_limit` features; must be omitted for other kinds.
    PeriodicityInterval     Optional[int]                        `json:"periodicity_interval"`
    // Required for `usage_limit` features; must be omitted for other kinds.
    PeriodicityUnit         Optional[EntitlementPeriodicityUnit] `json:"periodicity_unit"`
    // `usage_limit` features only. A positive number of `grant_expiration_unit`s; required with `day` or `month`, and must be omitted with `until_next_refill`.
    GrantExpirationInterval Optional[int]                        `json:"grant_expiration_interval"`
    // `usage_limit` features only.
    GrantExpirationUnit     Optional[GrantExpirationUnit]        `json:"grant_expiration_unit"`
    AdditionalProperties    map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for CloneFeatureCatalogItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CloneFeatureCatalogItem) String() string {
    return fmt.Sprintf(
    	"CloneFeatureCatalogItem[FeatureTemplateId=%v, Value=%v, PeriodicityInterval=%v, PeriodicityUnit=%v, GrantExpirationInterval=%v, GrantExpirationUnit=%v, AdditionalProperties=%v]",
    	c.FeatureTemplateId, c.Value, c.PeriodicityInterval, c.PeriodicityUnit, c.GrantExpirationInterval, c.GrantExpirationUnit, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CloneFeatureCatalogItem.
// It customizes the JSON marshaling process for CloneFeatureCatalogItem objects.
func (c CloneFeatureCatalogItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "feature_template_id", "value", "periodicity_interval", "periodicity_unit", "grant_expiration_interval", "grant_expiration_unit"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CloneFeatureCatalogItem object to a map representation for JSON marshaling.
func (c CloneFeatureCatalogItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["feature_template_id"] = c.FeatureTemplateId
    structMap["value"] = c.Value
    if c.PeriodicityInterval.IsValueSet() {
        if c.PeriodicityInterval.Value() != nil {
            structMap["periodicity_interval"] = c.PeriodicityInterval.Value()
        } else {
            structMap["periodicity_interval"] = nil
        }
    }
    if c.PeriodicityUnit.IsValueSet() {
        if c.PeriodicityUnit.Value() != nil {
            structMap["periodicity_unit"] = c.PeriodicityUnit.Value()
        } else {
            structMap["periodicity_unit"] = nil
        }
    }
    if c.GrantExpirationInterval.IsValueSet() {
        if c.GrantExpirationInterval.Value() != nil {
            structMap["grant_expiration_interval"] = c.GrantExpirationInterval.Value()
        } else {
            structMap["grant_expiration_interval"] = nil
        }
    }
    if c.GrantExpirationUnit.IsValueSet() {
        if c.GrantExpirationUnit.Value() != nil {
            structMap["grant_expiration_unit"] = c.GrantExpirationUnit.Value()
        } else {
            structMap["grant_expiration_unit"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CloneFeatureCatalogItem.
// It customizes the JSON unmarshaling process for CloneFeatureCatalogItem objects.
func (c *CloneFeatureCatalogItem) UnmarshalJSON(input []byte) error {
    var temp tempCloneFeatureCatalogItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "feature_template_id", "value", "periodicity_interval", "periodicity_unit", "grant_expiration_interval", "grant_expiration_unit")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.FeatureTemplateId = *temp.FeatureTemplateId
    c.Value = *temp.Value
    c.PeriodicityInterval = temp.PeriodicityInterval
    c.PeriodicityUnit = temp.PeriodicityUnit
    c.GrantExpirationInterval = temp.GrantExpirationInterval
    c.GrantExpirationUnit = temp.GrantExpirationUnit
    return nil
}

// tempCloneFeatureCatalogItem is a temporary struct used for validating the fields of CloneFeatureCatalogItem.
type tempCloneFeatureCatalogItem  struct {
    FeatureTemplateId       *int                                 `json:"feature_template_id"`
    Value                   *string                              `json:"value"`
    PeriodicityInterval     Optional[int]                        `json:"periodicity_interval"`
    PeriodicityUnit         Optional[EntitlementPeriodicityUnit] `json:"periodicity_unit"`
    GrantExpirationInterval Optional[int]                        `json:"grant_expiration_interval"`
    GrantExpirationUnit     Optional[GrantExpirationUnit]        `json:"grant_expiration_unit"`
}

func (c *tempCloneFeatureCatalogItem) validate() error {
    var errs []string
    if c.FeatureTemplateId == nil {
        errs = append(errs, "required field `feature_template_id` is missing for type `Clone Feature Catalog Item`")
    }
    if c.Value == nil {
        errs = append(errs, "required field `value` is missing for type `Clone Feature Catalog Item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
