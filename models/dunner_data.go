/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "strings"
    "time"
)

// DunnerData represents a DunnerData struct.
type DunnerData struct {
    State                string                 `json:"state"`
    SubscriptionId       int                    `json:"subscription_id"`
    RevenueAtRiskInCents int64                  `json:"revenue_at_risk_in_cents"`
    CreatedAt            time.Time              `json:"created_at"`
    Attempts             int                    `json:"attempts"`
    LastAttemptedAt      time.Time              `json:"last_attempted_at"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DunnerData,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DunnerData) String() string {
    return fmt.Sprintf(
    	"DunnerData[State=%v, SubscriptionId=%v, RevenueAtRiskInCents=%v, CreatedAt=%v, Attempts=%v, LastAttemptedAt=%v, AdditionalProperties=%v]",
    	d.State, d.SubscriptionId, d.RevenueAtRiskInCents, d.CreatedAt, d.Attempts, d.LastAttemptedAt, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DunnerData.
// It customizes the JSON marshaling process for DunnerData objects.
func (d DunnerData) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "state", "subscription_id", "revenue_at_risk_in_cents", "created_at", "attempts", "last_attempted_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DunnerData object to a map representation for JSON marshaling.
func (d DunnerData) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["state"] = d.State
    structMap["subscription_id"] = d.SubscriptionId
    structMap["revenue_at_risk_in_cents"] = d.RevenueAtRiskInCents
    structMap["created_at"] = d.CreatedAt.Format(time.RFC3339)
    structMap["attempts"] = d.Attempts
    structMap["last_attempted_at"] = d.LastAttemptedAt.Format(time.RFC3339)
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DunnerData.
// It customizes the JSON unmarshaling process for DunnerData objects.
func (d *DunnerData) UnmarshalJSON(input []byte) error {
    var temp tempDunnerData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "state", "subscription_id", "revenue_at_risk_in_cents", "created_at", "attempts", "last_attempted_at")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.State = *temp.State
    d.SubscriptionId = *temp.SubscriptionId
    d.RevenueAtRiskInCents = *temp.RevenueAtRiskInCents
    CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
    if err != nil {
        log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
    }
    d.CreatedAt = CreatedAtVal
    d.Attempts = *temp.Attempts
    LastAttemptedAtVal, err := time.Parse(time.RFC3339, *temp.LastAttemptedAt)
    if err != nil {
        log.Fatalf("Cannot Parse last_attempted_at as % s format.", time.RFC3339)
    }
    d.LastAttemptedAt = LastAttemptedAtVal
    return nil
}

// tempDunnerData is a temporary struct used for validating the fields of DunnerData.
type tempDunnerData  struct {
    State                *string `json:"state"`
    SubscriptionId       *int    `json:"subscription_id"`
    RevenueAtRiskInCents *int64  `json:"revenue_at_risk_in_cents"`
    CreatedAt            *string `json:"created_at"`
    Attempts             *int    `json:"attempts"`
    LastAttemptedAt      *string `json:"last_attempted_at"`
}

func (d *tempDunnerData) validate() error {
    var errs []string
    if d.State == nil {
        errs = append(errs, "required field `state` is missing for type `Dunner Data`")
    }
    if d.SubscriptionId == nil {
        errs = append(errs, "required field `subscription_id` is missing for type `Dunner Data`")
    }
    if d.RevenueAtRiskInCents == nil {
        errs = append(errs, "required field `revenue_at_risk_in_cents` is missing for type `Dunner Data`")
    }
    if d.CreatedAt == nil {
        errs = append(errs, "required field `created_at` is missing for type `Dunner Data`")
    }
    if d.Attempts == nil {
        errs = append(errs, "required field `attempts` is missing for type `Dunner Data`")
    }
    if d.LastAttemptedAt == nil {
        errs = append(errs, "required field `last_attempted_at` is missing for type `Dunner Data`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
