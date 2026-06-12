// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// BillingSchedule represents a BillingSchedule struct.
// Billing schedule settings for component allocations or usages on multi-frequency subscriptions. Use this to start a component's billing period on a custom date instead of aligning with the product charge schedule.
type BillingSchedule struct {
    // Custom start date (ISO 8601 date, YYYY-MM-DD) for the component's first billing period. If omitted or null, billing aligns with the product schedule. If provided, date must be on or after the minimum allowed date for the subscription or component.
    InitialBillingAt     Optional[time.Time]    `json:"initial_billing_at"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for BillingSchedule,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b BillingSchedule) String() string {
    return fmt.Sprintf(
    	"BillingSchedule[InitialBillingAt=%v, AdditionalProperties=%v]",
    	b.InitialBillingAt, b.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for BillingSchedule.
// It customizes the JSON marshaling process for BillingSchedule objects.
func (b BillingSchedule) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "initial_billing_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BillingSchedule object to a map representation for JSON marshaling.
func (b BillingSchedule) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
    if b.InitialBillingAt.IsValueSet() {
        var InitialBillingAtVal *string = nil
        if b.InitialBillingAt.Value() != nil {
            val := b.InitialBillingAt.Value().Format(DEFAULT_DATE)
            InitialBillingAtVal = &val
        }
        if b.InitialBillingAt.Value() != nil {
            structMap["initial_billing_at"] = InitialBillingAtVal
        } else {
            structMap["initial_billing_at"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BillingSchedule.
// It customizes the JSON unmarshaling process for BillingSchedule objects.
func (b *BillingSchedule) UnmarshalJSON(input []byte) error {
    var temp tempBillingSchedule
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "initial_billing_at")
    if err != nil {
    	return err
    }
    b.AdditionalProperties = additionalProperties
    
    b.InitialBillingAt.ShouldSetValue(temp.InitialBillingAt.IsValueSet())
    if temp.InitialBillingAt.Value() != nil {
        InitialBillingAtVal, err := time.Parse(DEFAULT_DATE, (*temp.InitialBillingAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse initial_billing_at as % s format.", DEFAULT_DATE)
        }
        b.InitialBillingAt.SetValue(&InitialBillingAtVal)
    }
    return nil
}

// tempBillingSchedule is a temporary struct used for validating the fields of BillingSchedule.
type tempBillingSchedule  struct {
    InitialBillingAt Optional[string] `json:"initial_billing_at"`
}
