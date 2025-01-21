/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// PrepaidConfiguration represents a PrepaidConfiguration struct.
type PrepaidConfiguration struct {
    Id                              *int                   `json:"id,omitempty"`
    InitialFundingAmountInCents     *int64                 `json:"initial_funding_amount_in_cents,omitempty"`
    ReplenishToAmountInCents        *int64                 `json:"replenish_to_amount_in_cents,omitempty"`
    AutoReplenish                   *bool                  `json:"auto_replenish,omitempty"`
    ReplenishThresholdAmountInCents *int64                 `json:"replenish_threshold_amount_in_cents,omitempty"`
    AdditionalProperties            map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PrepaidConfiguration,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PrepaidConfiguration) String() string {
    return fmt.Sprintf(
    	"PrepaidConfiguration[Id=%v, InitialFundingAmountInCents=%v, ReplenishToAmountInCents=%v, AutoReplenish=%v, ReplenishThresholdAmountInCents=%v, AdditionalProperties=%v]",
    	p.Id, p.InitialFundingAmountInCents, p.ReplenishToAmountInCents, p.AutoReplenish, p.ReplenishThresholdAmountInCents, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PrepaidConfiguration.
// It customizes the JSON marshaling process for PrepaidConfiguration objects.
func (p PrepaidConfiguration) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "id", "initial_funding_amount_in_cents", "replenish_to_amount_in_cents", "auto_replenish", "replenish_threshold_amount_in_cents"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaidConfiguration object to a map representation for JSON marshaling.
func (p PrepaidConfiguration) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Id != nil {
        structMap["id"] = p.Id
    }
    if p.InitialFundingAmountInCents != nil {
        structMap["initial_funding_amount_in_cents"] = p.InitialFundingAmountInCents
    }
    if p.ReplenishToAmountInCents != nil {
        structMap["replenish_to_amount_in_cents"] = p.ReplenishToAmountInCents
    }
    if p.AutoReplenish != nil {
        structMap["auto_replenish"] = p.AutoReplenish
    }
    if p.ReplenishThresholdAmountInCents != nil {
        structMap["replenish_threshold_amount_in_cents"] = p.ReplenishThresholdAmountInCents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaidConfiguration.
// It customizes the JSON unmarshaling process for PrepaidConfiguration objects.
func (p *PrepaidConfiguration) UnmarshalJSON(input []byte) error {
    var temp tempPrepaidConfiguration
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "initial_funding_amount_in_cents", "replenish_to_amount_in_cents", "auto_replenish", "replenish_threshold_amount_in_cents")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Id = temp.Id
    p.InitialFundingAmountInCents = temp.InitialFundingAmountInCents
    p.ReplenishToAmountInCents = temp.ReplenishToAmountInCents
    p.AutoReplenish = temp.AutoReplenish
    p.ReplenishThresholdAmountInCents = temp.ReplenishThresholdAmountInCents
    return nil
}

// tempPrepaidConfiguration is a temporary struct used for validating the fields of PrepaidConfiguration.
type tempPrepaidConfiguration  struct {
    Id                              *int   `json:"id,omitempty"`
    InitialFundingAmountInCents     *int64 `json:"initial_funding_amount_in_cents,omitempty"`
    ReplenishToAmountInCents        *int64 `json:"replenish_to_amount_in_cents,omitempty"`
    AutoReplenish                   *bool  `json:"auto_replenish,omitempty"`
    ReplenishThresholdAmountInCents *int64 `json:"replenish_threshold_amount_in_cents,omitempty"`
}
