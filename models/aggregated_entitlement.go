// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// AggregatedEntitlement represents a AggregatedEntitlement struct.
// One entitlement in a subscriber's aggregated entitlements list. Entries are aggregated per feature key and periodicity window, not per feature key alone. A `usage_limit` feature granted with two different periodicities yields two entries sharing one `feature_key`. Use `periodicity_key` to identify an entry uniquely.
type AggregatedEntitlement struct {
    // The feature's key, prefixed by kind: `feature.*` for `access_right`, `usage.*` for `usage_limit`, `service.*` for `service_right`.
    FeatureKey           *string                                    `json:"feature_key,omitempty"`
    // Uniquely identifies this aggregated entry: the prefixed feature key, suffixed with `:{interval}:{unit}` when the entitlement has a periodicity window. Equal to `feature_key` when `periodicity` is `null`.
    PeriodicityKey       *string                                    `json:"periodicity_key,omitempty"`
    Name                 *string                                    `json:"name,omitempty"`
    // The behavior of a feature:
    // - `access_right`: a boolean entitlement. A subscriber either has access or does not.
    // - `usage_limit`: a quantified allowance measured over a recurring period (for example, "10,000 API calls per month").
    // - `service_right`: a free-form value (text, boolean, or number) that isn't a simple access flag or a metered limit.
    Type                 *FeatureKind                               `json:"type,omitempty"`
    // The aggregated value, coerced according to `type`: a boolean for `access_right` (and boolean `service_right`), a number for `usage_limit` (and numeric `service_right`), or a string for text `service_right`.
    Value                *AggregatedEntitlementValue                `json:"value,omitempty"`
    // `true` only when the aggregated value is truthy for this feature's kind, and the subscription is in a live state (`active`, `trialing`, `assessing`, `past_due`, or `soft_failure`). `false` otherwise, including for `awaiting_signup`, canceled, expired, and on-hold subscriptions. Entitlements deliberately stay enabled through dunning.
    Enabled              *bool                                      `json:"enabled,omitempty"`
    Periodicity          Optional[AggregatedEntitlementPeriodicity] `json:"periodicity"`
    // The names of the products/components contributing to this entitlement. For `access_right` features, only contributors that granted `true` are listed.
    SourceProducts       []string                                   `json:"source_products,omitempty"`
    AdditionalProperties map[string]interface{}                     `json:"_"`
}

// String implements the fmt.Stringer interface for AggregatedEntitlement,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AggregatedEntitlement) String() string {
    return fmt.Sprintf(
    	"AggregatedEntitlement[FeatureKey=%v, PeriodicityKey=%v, Name=%v, Type=%v, Value=%v, Enabled=%v, Periodicity=%v, SourceProducts=%v, AdditionalProperties=%v]",
    	a.FeatureKey, a.PeriodicityKey, a.Name, a.Type, a.Value, a.Enabled, a.Periodicity, a.SourceProducts, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AggregatedEntitlement.
// It customizes the JSON marshaling process for AggregatedEntitlement objects.
func (a AggregatedEntitlement) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "feature_key", "periodicity_key", "name", "type", "value", "enabled", "periodicity", "source_products"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AggregatedEntitlement object to a map representation for JSON marshaling.
func (a AggregatedEntitlement) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.FeatureKey != nil {
        structMap["feature_key"] = a.FeatureKey
    }
    if a.PeriodicityKey != nil {
        structMap["periodicity_key"] = a.PeriodicityKey
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.Type != nil {
        structMap["type"] = a.Type
    }
    if a.Value != nil {
        structMap["value"] = a.Value.toMap()
    }
    if a.Enabled != nil {
        structMap["enabled"] = a.Enabled
    }
    if a.Periodicity.IsValueSet() {
        if a.Periodicity.Value() != nil {
            structMap["periodicity"] = a.Periodicity.Value().toMap()
        } else {
            structMap["periodicity"] = nil
        }
    }
    if a.SourceProducts != nil {
        structMap["source_products"] = a.SourceProducts
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AggregatedEntitlement.
// It customizes the JSON unmarshaling process for AggregatedEntitlement objects.
func (a *AggregatedEntitlement) UnmarshalJSON(input []byte) error {
    var temp tempAggregatedEntitlement
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "feature_key", "periodicity_key", "name", "type", "value", "enabled", "periodicity", "source_products")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.FeatureKey = temp.FeatureKey
    a.PeriodicityKey = temp.PeriodicityKey
    a.Name = temp.Name
    a.Type = temp.Type
    a.Value = temp.Value
    a.Enabled = temp.Enabled
    a.Periodicity = temp.Periodicity
    a.SourceProducts = temp.SourceProducts
    return nil
}

// tempAggregatedEntitlement is a temporary struct used for validating the fields of AggregatedEntitlement.
type tempAggregatedEntitlement  struct {
    FeatureKey     *string                                    `json:"feature_key,omitempty"`
    PeriodicityKey *string                                    `json:"periodicity_key,omitempty"`
    Name           *string                                    `json:"name,omitempty"`
    Type           *FeatureKind                               `json:"type,omitempty"`
    Value          *AggregatedEntitlementValue                `json:"value,omitempty"`
    Enabled        *bool                                      `json:"enabled,omitempty"`
    Periodicity    Optional[AggregatedEntitlementPeriodicity] `json:"periodicity"`
    SourceProducts []string                                   `json:"source_products,omitempty"`
}
