/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// BillingSchedule represents a BillingSchedule struct.
// This attribute is particularly useful when you need to align billing events for different components on distinct schedules within a subscription. Please note this only works for site with Multifrequency enabled
type BillingSchedule struct {
    // The initial_billing_at attribute in Maxio allows you to specify a custom starting date for billing cycles associated with components that have their own billing frequency set. Only ISO8601 format is supported.
    InitialBillingAt     *time.Time             `json:"initial_billing_at,omitempty"`
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
    if b.InitialBillingAt != nil {
        structMap["initial_billing_at"] = b.InitialBillingAt.Format(DEFAULT_DATE)
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
    
    if temp.InitialBillingAt != nil {
        InitialBillingAtVal, err := time.Parse(DEFAULT_DATE, *temp.InitialBillingAt)
        if err != nil {
            log.Fatalf("Cannot Parse initial_billing_at as % s format.", DEFAULT_DATE)
        }
        b.InitialBillingAt = &InitialBillingAtVal
    }
    return nil
}

// tempBillingSchedule is a temporary struct used for validating the fields of BillingSchedule.
type tempBillingSchedule  struct {
    InitialBillingAt *string `json:"initial_billing_at,omitempty"`
}
