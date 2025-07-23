// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// UpsertPrepaidConfiguration represents a UpsertPrepaidConfiguration struct.
type UpsertPrepaidConfiguration struct {
    InitialFundingAmountInCents     *int64                 `json:"initial_funding_amount_in_cents,omitempty"`
    ReplenishToAmountInCents        *int64                 `json:"replenish_to_amount_in_cents,omitempty"`
    AutoReplenish                   *bool                  `json:"auto_replenish,omitempty"`
    ReplenishThresholdAmountInCents *int64                 `json:"replenish_threshold_amount_in_cents,omitempty"`
    AdditionalProperties            map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpsertPrepaidConfiguration,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpsertPrepaidConfiguration) String() string {
    return fmt.Sprintf(
    	"UpsertPrepaidConfiguration[InitialFundingAmountInCents=%v, ReplenishToAmountInCents=%v, AutoReplenish=%v, ReplenishThresholdAmountInCents=%v, AdditionalProperties=%v]",
    	u.InitialFundingAmountInCents, u.ReplenishToAmountInCents, u.AutoReplenish, u.ReplenishThresholdAmountInCents, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpsertPrepaidConfiguration.
// It customizes the JSON marshaling process for UpsertPrepaidConfiguration objects.
func (u UpsertPrepaidConfiguration) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "initial_funding_amount_in_cents", "replenish_to_amount_in_cents", "auto_replenish", "replenish_threshold_amount_in_cents"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpsertPrepaidConfiguration object to a map representation for JSON marshaling.
func (u UpsertPrepaidConfiguration) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.InitialFundingAmountInCents != nil {
        structMap["initial_funding_amount_in_cents"] = u.InitialFundingAmountInCents
    }
    if u.ReplenishToAmountInCents != nil {
        structMap["replenish_to_amount_in_cents"] = u.ReplenishToAmountInCents
    }
    if u.AutoReplenish != nil {
        structMap["auto_replenish"] = u.AutoReplenish
    }
    if u.ReplenishThresholdAmountInCents != nil {
        structMap["replenish_threshold_amount_in_cents"] = u.ReplenishThresholdAmountInCents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpsertPrepaidConfiguration.
// It customizes the JSON unmarshaling process for UpsertPrepaidConfiguration objects.
func (u *UpsertPrepaidConfiguration) UnmarshalJSON(input []byte) error {
    var temp tempUpsertPrepaidConfiguration
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "initial_funding_amount_in_cents", "replenish_to_amount_in_cents", "auto_replenish", "replenish_threshold_amount_in_cents")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.InitialFundingAmountInCents = temp.InitialFundingAmountInCents
    u.ReplenishToAmountInCents = temp.ReplenishToAmountInCents
    u.AutoReplenish = temp.AutoReplenish
    u.ReplenishThresholdAmountInCents = temp.ReplenishThresholdAmountInCents
    return nil
}

// tempUpsertPrepaidConfiguration is a temporary struct used for validating the fields of UpsertPrepaidConfiguration.
type tempUpsertPrepaidConfiguration  struct {
    InitialFundingAmountInCents     *int64 `json:"initial_funding_amount_in_cents,omitempty"`
    ReplenishToAmountInCents        *int64 `json:"replenish_to_amount_in_cents,omitempty"`
    AutoReplenish                   *bool  `json:"auto_replenish,omitempty"`
    ReplenishThresholdAmountInCents *int64 `json:"replenish_threshold_amount_in_cents,omitempty"`
}
