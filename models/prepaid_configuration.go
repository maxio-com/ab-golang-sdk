package models

import (
    "encoding/json"
)

// PrepaidConfiguration represents a PrepaidConfiguration struct.
type PrepaidConfiguration struct {
    Id                              *int           `json:"id,omitempty"`
    InitialFundingAmountInCents     *int64         `json:"initial_funding_amount_in_cents,omitempty"`
    ReplenishToAmountInCents        *int64         `json:"replenish_to_amount_in_cents,omitempty"`
    AutoReplenish                   *bool          `json:"auto_replenish,omitempty"`
    ReplenishThresholdAmountInCents *int64         `json:"replenish_threshold_amount_in_cents,omitempty"`
    AdditionalProperties            map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaidConfiguration.
// It customizes the JSON marshaling process for PrepaidConfiguration objects.
func (p PrepaidConfiguration) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaidConfiguration object to a map representation for JSON marshaling.
func (p PrepaidConfiguration) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
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
    var temp prepaidConfiguration
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "initial_funding_amount_in_cents", "replenish_to_amount_in_cents", "auto_replenish", "replenish_threshold_amount_in_cents")
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

// TODO
type prepaidConfiguration  struct {
    Id                              *int   `json:"id,omitempty"`
    InitialFundingAmountInCents     *int64 `json:"initial_funding_amount_in_cents,omitempty"`
    ReplenishToAmountInCents        *int64 `json:"replenish_to_amount_in_cents,omitempty"`
    AutoReplenish                   *bool  `json:"auto_replenish,omitempty"`
    ReplenishThresholdAmountInCents *int64 `json:"replenish_threshold_amount_in_cents,omitempty"`
}
