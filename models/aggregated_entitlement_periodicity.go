// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// AggregatedEntitlementPeriodicity represents a AggregatedEntitlementPeriodicity struct.
type AggregatedEntitlementPeriodicity struct {
    Interval             *int                        `json:"interval,omitempty"`
    // The recurring window over which a `usage_limit` feature's allowance resets.
    Unit                 *EntitlementPeriodicityUnit `json:"unit,omitempty"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for AggregatedEntitlementPeriodicity,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AggregatedEntitlementPeriodicity) String() string {
    return fmt.Sprintf(
    	"AggregatedEntitlementPeriodicity[Interval=%v, Unit=%v, AdditionalProperties=%v]",
    	a.Interval, a.Unit, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AggregatedEntitlementPeriodicity.
// It customizes the JSON marshaling process for AggregatedEntitlementPeriodicity objects.
func (a AggregatedEntitlementPeriodicity) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "interval", "unit"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AggregatedEntitlementPeriodicity object to a map representation for JSON marshaling.
func (a AggregatedEntitlementPeriodicity) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Interval != nil {
        structMap["interval"] = a.Interval
    }
    if a.Unit != nil {
        structMap["unit"] = a.Unit
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AggregatedEntitlementPeriodicity.
// It customizes the JSON unmarshaling process for AggregatedEntitlementPeriodicity objects.
func (a *AggregatedEntitlementPeriodicity) UnmarshalJSON(input []byte) error {
    var temp tempAggregatedEntitlementPeriodicity
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "interval", "unit")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Interval = temp.Interval
    a.Unit = temp.Unit
    return nil
}

// tempAggregatedEntitlementPeriodicity is a temporary struct used for validating the fields of AggregatedEntitlementPeriodicity.
type tempAggregatedEntitlementPeriodicity  struct {
    Interval *int                        `json:"interval,omitempty"`
    Unit     *EntitlementPeriodicityUnit `json:"unit,omitempty"`
}
